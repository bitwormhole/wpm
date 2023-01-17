package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/service"
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
	ec.Handle(http.MethodPost, "", inst.handlePost)
	ec.Handle(http.MethodGet, ":id", inst.handleGetOne)
	ec.Handle(http.MethodPut, ":id", inst.handlePut)
	ec.Handle(http.MethodDelete, ":id", inst.handleDelete)

	ec.Handle(http.MethodGet, "/media/*file", inst.handleGetFile)

	return nil
}

func (inst *MediaController) handleGetList(c *gin.Context) {
	req := &myMediaRequest{
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

func (inst *MediaController) handleGetOne(c *gin.Context) {
	req := &myMediaRequest{
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

	wantRequestID   bool
	wantRequestBody bool

	id    dxo.MediaID
	body1 vo.Media
	body2 vo.Media
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

	return nil
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
	o, err := ser.Find(ctx, id)
	if err != nil {
		return err
	}
	inst.body2.Mediae = []*dto.Media{o}
	return nil
}

func (inst *myMediaRequest) doGetList() error {
	ctx := inst.gc
	ser := inst.controller.MediaService
	list, err := ser.ListAll(ctx)
	if err != nil {
		return err
	}
	inst.body2.Mediae = list
	return nil
}

func (inst *myMediaRequest) doPost() error {

	return fmt.Errorf("no impl")
}

func (inst *myMediaRequest) doPut() error {

	return fmt.Errorf("no impl")
}

func (inst *myMediaRequest) doDelete() error {

	return fmt.Errorf("no impl")
}

func (inst *myMediaRequest) doGetFile(c *gin.Context) error {

	path := c.Param("file")
	ser := inst.controller.MediaService

	me, err := ser.FindByPath(c, path)
	if err != nil {
		return err
	}

	me, err = ser.PrepareForDownload(c, me)
	if err != nil {
		return err
	}

	c.File(me.LocalFilePath)
	return nil
}
