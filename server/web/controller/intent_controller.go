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

// IntentController Intent 控制器
type IntentController struct {
	markup.RestController `class:"rest-controller"`

	IntentService service.IntentService `inject:"#IntentService"`
	Responder     glass.MainResponder   `inject:"#glass-main-responder"`
}

func (inst *IntentController) _Impl() glass.Controller {
	return inst
}

// Init 初始化
func (inst *IntentController) Init(ec glass.EngineConnection) error {

	ec = ec.RequestMapping("intents")

	ec.Handle(http.MethodGet, "", inst.handleGetList)
	ec.Handle(http.MethodPost, "", inst.handlePost)

	ec.Handle(http.MethodGet, ":id", inst.handleGetOne)
	ec.Handle(http.MethodPut, ":id", inst.handlePut)
	ec.Handle(http.MethodDelete, ":id", inst.handleDelete)

	return nil
}

func (inst *IntentController) handleGetList(c *gin.Context) {
	req := &myIntentRequest{
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

func (inst *IntentController) handleGetOne(c *gin.Context) {
	req := &myIntentRequest{
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

func (inst *IntentController) handlePost(c *gin.Context) {
	req := &myIntentRequest{
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

func (inst *IntentController) handlePut(c *gin.Context) {
	req := &myIntentRequest{
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

func (inst *IntentController) handleDelete(c *gin.Context) {
	req := &myIntentRequest{
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

type myIntentRequest struct {
	gc         *gin.Context
	controller *IntentController

	wantRequestID   bool
	wantRequestBody bool

	id    dxo.IntentID
	body1 vo.Intent
	body2 vo.Intent
}

func (inst *myIntentRequest) open() error {

	c := inst.gc

	if inst.wantRequestID {
		id := c.Param("id")
		n, err := strconv.Atoi(id)
		if err != nil {
			return err
		}
		inst.id = dxo.IntentID(n)
	}

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *myIntentRequest) send(err error) {
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

func (inst *myIntentRequest) doGetOne() error {
	id := inst.id
	ctx := inst.gc
	ser := inst.controller.IntentService
	o, err := ser.Find(ctx, id)
	if err != nil {
		return err
	}
	inst.body2.Intents = []*dto.Intent{o}
	return nil
}

func (inst *myIntentRequest) doGetList() error {
	ctx := inst.gc
	ser := inst.controller.IntentService
	list, err := ser.ListAll(ctx)
	if err != nil {
		return err
	}
	inst.body2.Intents = list
	return nil
}

func (inst *myIntentRequest) doPost() error {

	return nil
}

func (inst *myIntentRequest) doPut() error {
	return nil
}

func (inst *myIntentRequest) doDelete() error {

	return nil
}
