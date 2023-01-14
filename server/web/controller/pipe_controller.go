package controller

import (
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

// PipeController Pipe 控制器
type PipeController struct {
	markup.RestController `class:"rest-controller"`

	PipeService service.PipeService `inject:"#PipeService"`
	Responder   glass.MainResponder `inject:"#glass-main-responder"`
}

func (inst *PipeController) _Impl() glass.Controller {
	return inst
}

// Init 初始化
func (inst *PipeController) Init(ec glass.EngineConnection) error {

	ec = ec.RequestMapping("pipe")

	ec.Handle(http.MethodGet, "", inst.handleGetList)
	ec.Handle(http.MethodPost, "", inst.handlePost)
	ec.Handle(http.MethodGet, ":id", inst.handleGet)
	ec.Handle(http.MethodPut, ":id", inst.handlePut)
	ec.Handle(http.MethodDelete, ":id", inst.handleDelete)

	return nil
}

func (inst *PipeController) handleGet(c *gin.Context) {
	req := &myPipeRequest{
		gc:               c,
		controller:       inst,
		wantRequestID:    true,
		wantRequestBody:  false,
		wantRequestQuery: true,
	}
	err := req.open()
	if err == nil {
		err = req.doGet()
	}
	req.send(err)
}

func (inst *PipeController) handleGetList(c *gin.Context) {
	req := &myPipeRequest{
		gc:               c,
		controller:       inst,
		wantRequestID:    false,
		wantRequestBody:  false,
		wantRequestQuery: false,
	}
	err := req.open()
	if err == nil {
		err = req.doGetList()
	}
	req.send(err)
}

func (inst *PipeController) handlePost(c *gin.Context) {
	req := &myPipeRequest{
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

func (inst *PipeController) handlePut(c *gin.Context) {
	req := &myPipeRequest{
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

func (inst *PipeController) handleDelete(c *gin.Context) {
	req := &myPipeRequest{
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

////////////////////////////////////////////////////////////////////////////////

type myPipeRequest struct {
	gc         *gin.Context
	controller *PipeController

	wantRequestID    bool
	wantRequestBody  bool
	wantRequestQuery bool

	id      dxo.PipeID
	timeout int
	body1   vo.Pipe
	body2   vo.Pipe
}

func (inst *myPipeRequest) open() error {

	c := inst.gc

	if inst.wantRequestID {
		id := c.Param("id")
		n, err := strconv.Atoi(id)
		if err != nil {
			return err
		}
		inst.id = dxo.PipeID(n)
	}

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	if inst.wantRequestQuery {
		timeoutStr := c.Query("timeout")
		timeoutInt, _ := strconv.Atoi(timeoutStr)
		inst.timeout = timeoutInt
	}

	return nil
}

func (inst *myPipeRequest) send(err error) {
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

func (inst *myPipeRequest) doGet() error {
	id := inst.id
	ctx := inst.gc
	ser := inst.controller.PipeService
	o1 := &inst.body1
	timeout := inst.timeout
	err := ser.Pull(ctx, id, o1, timeout)
	if err != nil {
		return err
	}
	inst.body2 = *o1
	return nil
}

func (inst *myPipeRequest) doGetList() error {
	ctx := inst.gc
	ser := inst.controller.PipeService
	all, err := ser.ListAll(ctx)
	if err != nil {
		return err
	}
	inst.body2.Pipes = all
	return nil
}

func (inst *myPipeRequest) doPost() error {
	ctx := inst.gc
	ser := inst.controller.PipeService
	o1 := inst.body1.Pipes[0]
	o2, err := ser.Insert(ctx, o1)
	if err != nil {
		return err
	}
	inst.body2.Pipes = []*dto.PipeInfo{o2}
	return nil
}

func (inst *myPipeRequest) doPut() error {

	ctx := inst.gc
	ser := inst.controller.PipeService
	id := inst.id
	o1 := &inst.body1

	return ser.Push(ctx, id, o1)
}

func (inst *myPipeRequest) doDelete() error {
	id := inst.id
	ctx := inst.gc
	ser := inst.controller.PipeService
	return ser.Remove(ctx, id)
}
