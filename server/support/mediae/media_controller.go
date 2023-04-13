package mediae

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/util"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/utils"
	"github.com/bitwormhole/wpm/server/web/dto"
	"github.com/bitwormhole/wpm/server/web/vo"
	"github.com/gin-gonic/gin"
)

// MediaController Media 控制器
type MediaController struct {
	markup.RestController `class:"rest-controller"`

	MediaService service.MediaService `inject:"#MediaService"`
	Responder    glass.MainResponder  `inject:"#glass-main-responder"`
}

func (inst *MediaController) _Impl() glass.Controller {
	return inst
}

// Init 初始化
func (inst *MediaController) Init(ec glass.EngineConnection) error {

	ec = ec.RequestMapping("media")

	ec.Handle(http.MethodGet, "", inst.handleGetList)
	ec.Handle(http.MethodGet, ":id", inst.handleGetOne)

	ec.Handle(http.MethodPost, "", inst.handlePost)
	ec.Handle(http.MethodPost, "import-presets", inst.handlePostImportPresets)

	ec.Handle(http.MethodPut, ":id", inst.handlePut)
	ec.Handle(http.MethodDelete, ":id", inst.handleDelete)

	ec.Handle(http.MethodGet, "/media/:type1/:type2/:hex/:name", inst.handleGetFile)

	return nil
}

func (inst *MediaController) handleGetList(c *gin.Context) {
	req := &myMediaRequest{
		gc:                 c,
		controller:         inst,
		wantRequestID:      false,
		wantRequestIDs:     true,
		wantRequestBody:    false,
		wantRequestOptions: true,
	}
	err := req.open()
	if err == nil {
		err = req.doGetList()
	}
	req.send(err)
}

func (inst *MediaController) handleGetOne(c *gin.Context) {
	req := &myMediaRequest{
		gc:                 c,
		controller:         inst,
		wantRequestID:      true,
		wantRequestIDs:     false,
		wantRequestBody:    false,
		wantRequestOptions: true,
	}
	err := req.open()
	if err == nil {
		err = req.doGetOne()
	}
	req.send(err)
}

func (inst *MediaController) handlePost(c *gin.Context) {
	req := &myMediaRequest{
		gc:              c,
		controller:      inst,
		wantRequestID:   false,
		wantRequestBody: true,
	}
	err := req.open()
	if err == nil {
		err = req.doPost()
	}
	req.send(err)
}

func (inst *MediaController) handlePostImportPresets(c *gin.Context) {
	req := &myMediaRequest{
		gc:              c,
		controller:      inst,
		wantRequestID:   false,
		wantRequestBody: true,
	}
	err := req.open()
	if err == nil {
		err = req.doPostImportPresets()
	}
	req.send(err)
}

func (inst *MediaController) handlePut(c *gin.Context) {
	req := &myMediaRequest{
		gc:              c,
		controller:      inst,
		wantRequestID:   true,
		wantRequestBody: true,
	}
	err := req.open()
	if err == nil {
		err = req.doPut()
	}
	req.send(err)
}

func (inst *MediaController) handleDelete(c *gin.Context) {
	req := &myMediaRequest{
		gc:              c,
		controller:      inst,
		wantRequestID:   true,
		wantRequestBody: false,
	}
	err := req.open()
	if err == nil {
		err = req.doDelete()
	}
	req.send(err)
}

func (inst *MediaController) handleGetFile(c *gin.Context) {
	req := &myMediaRequest{
		controller: inst,
		gc:         c,
	}
	err := req.doGetFile(c)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	}
}

////////////////////////////////////////////////////////////////////////////////

type myMediaRequest struct {
	gc         *gin.Context
	controller *MediaController

	wantRequestID      bool
	wantRequestIDs     bool
	wantRequestBody    bool
	wantRequestOptions bool

	options service.MediaOptions
	id      dxo.MediaID
	ids     []dxo.MediaID
	body1   vo.Media
	body2   vo.Media
}

func (inst *myMediaRequest) open() error {

	c := inst.gc

	if inst.wantRequestID {
		id := c.Param("id")
		n, err := strconv.Atoi(id)
		if err != nil {
			return err
		}
		inst.id = dxo.MediaID(n)
	}

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	if inst.wantRequestOptions {
		inst.options.All = inst.hasFlag(c, "all")
		inst.options.WithFileState = inst.hasFlag(c, "with-file-state")
	}

	if inst.wantRequestIDs {
		idstr := c.Query("ids")
		ids, err := inst.parseIds(idstr)
		if err != nil {
			return err
		}
		inst.ids = ids
	}

	return nil
}

func (inst *myMediaRequest) parseIds(str string) ([]dxo.MediaID, error) {
	list := make([]dxo.MediaID, 0)
	err := (&utils.GinUtils{}).ParseIDs(str, ".", func(n int) {
		list = append(list, dxo.MediaID(n))
	})
	return list, err
}

func (inst *myMediaRequest) hasFlag(c *gin.Context, name string) bool {
	return (&utils.GinUtils{}).HasFlag(c, name)
}

func (inst *myMediaRequest) send(err error) {
	ctx := inst.gc
	data := &inst.body2
	status := data.Status
	resp := &glass.Response{
		Context: ctx,
		Data:    data,
		Error:   err,
		Status:  status,
	}
	inst.controller.Responder.Send(resp)
}

func (inst *myMediaRequest) doGetOne() error {
	id := inst.id
	ctx := inst.gc
	ser := inst.controller.MediaService
	opt := &inst.options
	o, err := ser.Find(ctx, id, opt)
	if err != nil {
		return err
	}
	inst.body2.Mediae = []*dto.Media{o}
	return nil
}

func (inst *myMediaRequest) doGetList() error {

	ctx := inst.gc
	ser := inst.controller.MediaService
	opt := &inst.options
	all := inst.options.All
	err := fmt.Errorf("")
	list := inst.body2.Mediae

	if all {
		list, err = ser.ListAll(ctx, opt)
	} else {
		ids := inst.ids
		list, err = ser.FindByIDs(ctx, ids, opt)
	}

	if err != nil {
		return err
	}
	inst.body2.Mediae = list
	return nil
}

func (inst *myMediaRequest) doPost() error {
	ctx := inst.gc
	ser := inst.controller.MediaService
	o1 := inst.body1.Mediae[0]
	o2, err := ser.Insert(ctx, o1, nil)
	if err != nil {
		return err
	}
	inst.body2.Mediae = []*dto.Media{o2}
	return nil
}

func (inst *myMediaRequest) doPostImportPresets() error {
	ctx := inst.gc
	ser := inst.controller.MediaService
	return ser.ImportPresets(ctx)
}

func (inst *myMediaRequest) doPut() error {
	ctx := inst.gc
	ser := inst.controller.MediaService
	id := inst.id
	o1 := inst.body1.Mediae[0]
	o2, err := ser.Update(ctx, id, o1, nil)
	if err != nil {
		return err
	}
	inst.body2.Mediae = []*dto.Media{o2}
	return nil
}

func (inst *myMediaRequest) doDelete() error {
	ctx := inst.gc
	ser := inst.controller.MediaService
	id := inst.id
	return ser.Remove(ctx, id)
}

func (inst *myMediaRequest) doGetFile(c *gin.Context) error {

	type1 := c.Param("type1")
	type2 := c.Param("type2")
	hex := c.Param("hex")
	name := c.Param("name")

	me := &dto.Media{
		ContentType: type1 + "/" + type2,
		Name:        name,
		SHA256SUM:   util.Hex(hex),
	}

	ser := inst.controller.MediaService
	me, err := ser.PrepareForDownload(c, me)
	if err != nil {
		return err
	}

	c.File(me.Source)
	return nil
}
