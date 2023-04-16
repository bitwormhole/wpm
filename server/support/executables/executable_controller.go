package executables

import (
	"net/http"
	"strconv"

	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/utils"
	"github.com/bitwormhole/wpm/server/web/dto"
	"github.com/bitwormhole/wpm/server/web/vo"
	"github.com/gin-gonic/gin"
)

// ExecutableController 仓库控制器
type ExecutableController struct {
	markup.RestController `class:"rest-controller"`

	ExecutableService service.ExecutableService `inject:"#ExecutableService"`
	Responder         glass.MainResponder       `inject:"#glass-main-responder"`
}

func (inst *ExecutableController) _Impl() glass.Controller {
	return inst
}

// Init 初始化
func (inst *ExecutableController) Init(ec glass.EngineConnection) error {

	ec = ec.RequestMapping("executables")

	ec.Handle(http.MethodGet, "", inst.handleGetList)
	ec.Handle(http.MethodGet, ":id", inst.handleGetOne)

	ec.Handle(http.MethodPost, "", inst.handlePost)

	ec.Handle(http.MethodPut, ":id", inst.handlePut)

	ec.Handle(http.MethodDelete, ":id", inst.handleDeleteOne)
	ec.Handle(http.MethodDelete, "", inst.handleDeleteMulti)

	return nil
}

func (inst *ExecutableController) handleGetList(c *gin.Context) {
	req := &myExecutableRequest{
		gc:              c,
		controller:      inst,
		wantRequestID:   false,
		wantRequestBody: false,
	}
	err := req.open()
	if err == nil {
		err = req.doGetList()
	}
	req.send(err)
}

func (inst *ExecutableController) handleGetOne(c *gin.Context) {
	req := &myExecutableRequest{
		gc:              c,
		controller:      inst,
		wantRequestID:   true,
		wantRequestBody: false,
	}
	err := req.open()
	if err == nil {
		err = req.doGetOne()
	}
	req.send(err)
}

func (inst *ExecutableController) handlePost(c *gin.Context) {
	req := &myExecutableRequest{
		gc:                 c,
		controller:         inst,
		wantRequestID:      false,
		wantRequestBody:    true,
		wantRequestOptions: true,
	}
	err := req.open()
	if err == nil {
		err = req.doPost()
	}
	req.send(err)
}

func (inst *ExecutableController) handlePut(c *gin.Context) {
	req := &myExecutableRequest{
		gc:                 c,
		controller:         inst,
		wantRequestID:      true,
		wantRequestBody:    true,
		wantRequestOptions: true,
	}
	err := req.open()
	if err == nil {
		err = req.doPut()
	}
	req.send(err)
}

func (inst *ExecutableController) handleDeleteOne(c *gin.Context) {
	req := &myExecutableRequest{
		gc:              c,
		controller:      inst,
		wantRequestID:   true,
		wantRequestBody: false,
	}
	err := req.open()
	if err == nil {
		err = req.doDelete(req.id)
	}
	req.send(err)
}

func (inst *ExecutableController) handleDeleteMulti(c *gin.Context) {
	req := &myExecutableRequest{
		gc:              c,
		controller:      inst,
		wantRequestIDs:  true,
		wantRequestBody: false,
	}
	err := req.open()
	if err == nil {
		err = req.doDelete(req.ids...)
	}
	req.send(err)
}

////////////////////////////////////////////////////////////////////////////////

type myExecutableRequest struct {
	gc         *gin.Context
	controller *ExecutableController

	wantRequestID      bool
	wantRequestIDs     bool
	wantRequestBody    bool
	wantRequestOptions bool

	options *service.ExecutableOptions
	id      dxo.ExecutableID
	ids     []dxo.ExecutableID

	body1 vo.Executable
	body2 vo.Executable
}

func (inst *myExecutableRequest) open() error {

	c := inst.gc

	if inst.wantRequestID {
		id := c.Param("id")
		n, err := strconv.Atoi(id)
		if err != nil {
			return err
		}
		inst.id = dxo.ExecutableID(n)
	}

	if inst.wantRequestIDs {
		idsstr := c.Query("ids")
		ids, err := inst.parseIds(idsstr)
		if err != nil {
			return err
		}
		inst.ids = ids
	}

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	if inst.wantRequestOptions {
		opt, err := inst.readOptions()
		if err != nil {
			return err
		}
		inst.options = opt
	}

	return nil
}

func (inst *myExecutableRequest) parseIds(str string) ([]dxo.ExecutableID, error) {
	list := make([]dxo.ExecutableID, 0)
	err := (&utils.GinUtils{}).ParseIDs(str, ".", func(n int) {
		list = append(list, dxo.ExecutableID(n))
	})
	return list, err
}

func (inst *myExecutableRequest) readOptions() (*service.ExecutableOptions, error) {

	opt := &service.ExecutableOptions{}
	data := &inst.body1
	c := inst.gc
	gu := utils.GinUtils{}

	// from url.query
	skipFileChecking := gu.HasFlag(c, "skip-file-checking")
	if skipFileChecking {
		data.OptionSkipFileChecking = true
	}

	// from request body
	opt.SkipFileChecking = data.OptionSkipFileChecking

	return opt, nil
}

func (inst *myExecutableRequest) send(err error) {
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

func (inst *myExecutableRequest) doGetOne() error {
	id := inst.id
	ctx := inst.gc
	ser := inst.controller.ExecutableService
	o, err := ser.Find(ctx, id, nil)
	if err != nil {
		return err
	}
	inst.body2.Executables = []*dto.Executable{o}
	return nil
}

func (inst *myExecutableRequest) doGetList() error {
	ctx := inst.gc
	ser := inst.controller.ExecutableService
	list, err := ser.ListAll(ctx, nil)
	if err != nil {
		return err
	}
	inst.body2.Executables = list
	return nil
}

func (inst *myExecutableRequest) doPost() error {
	ctx := inst.gc
	ser := inst.controller.ExecutableService
	opt := inst.options
	o1 := inst.body1.Executables[0]
	o2, err := ser.Insert(ctx, o1, opt)
	if err != nil {
		return err
	}
	inst.body2.Executables = append(inst.body2.Executables, o2)
	return nil
}

func (inst *myExecutableRequest) doPut() error {
	ctx := inst.gc
	ser := inst.controller.ExecutableService
	opt := inst.options
	o1 := inst.body1.Executables[0]
	id := inst.id
	o2, err := ser.Update(ctx, id, o1, opt)
	if err != nil {
		return err
	}
	inst.body2.Executables = append(inst.body2.Executables, o2)
	return nil
}

func (inst *myExecutableRequest) doDelete(ids ...dxo.ExecutableID) error {
	ctx := inst.gc
	ser := inst.controller.ExecutableService
	elist := utils.ErrorList{}
	for _, id := range ids {
		err := ser.Remove(ctx, id, nil)
		elist.Append(err)
	}
	return elist.First()
}
