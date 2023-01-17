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

// IntentTemplateController IntentTemplate 控制器
type IntentTemplateController struct {
	markup.RestController `class:"rest-controller"`

	IntentTemplateService service.IntentTemplateService `inject:"#IntentTemplateService"`
	Responder             glass.MainResponder           `inject:"#glass-main-responder"`
}

func (inst *IntentTemplateController) _Impl() glass.Controller {
	return inst
}

// Init 初始化
func (inst *IntentTemplateController) Init(ec glass.EngineConnection) error {

	ec = ec.RequestMapping("intent-templates")

	ec.Handle(http.MethodGet, "", inst.handleGetList)
	ec.Handle(http.MethodPost, "", inst.handlePost)

	ec.Handle(http.MethodGet, ":id", inst.handleGetOne)
	ec.Handle(http.MethodPut, ":id", inst.handlePut)
	ec.Handle(http.MethodDelete, ":id", inst.handleDelete)

	return nil
}

func (inst *IntentTemplateController) handleGetList(c *gin.Context) {
	req := &myIntentTemplateRequest{
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

func (inst *IntentTemplateController) handleGetOne(c *gin.Context) {
	req := &myIntentTemplateRequest{
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

func (inst *IntentTemplateController) handlePost(c *gin.Context) {
	req := &myIntentTemplateRequest{
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

func (inst *IntentTemplateController) handlePut(c *gin.Context) {
	req := &myIntentTemplateRequest{
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

func (inst *IntentTemplateController) handleDelete(c *gin.Context) {
	req := &myIntentTemplateRequest{
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

type myIntentTemplateRequest struct {
	gc         *gin.Context
	controller *IntentTemplateController

	wantRequestID   bool
	wantRequestBody bool

	id    dxo.IntentTemplateID
	body1 vo.IntentTemplate
	body2 vo.IntentTemplate
}

func (inst *myIntentTemplateRequest) open() error {

	c := inst.gc

	if inst.wantRequestID {
		id := c.Param("id")
		n, err := strconv.Atoi(id)
		if err != nil {
			return err
		}
		inst.id = dxo.IntentTemplateID(n)
	}

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *myIntentTemplateRequest) send(err error) {
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

func (inst *myIntentTemplateRequest) doGetOne() error {
	id := inst.id
	ctx := inst.gc
	ser := inst.controller.IntentTemplateService
	o, err := ser.Find(ctx, id)
	if err != nil {
		return err
	}
	inst.body2.Templates = []*dto.IntentTemplate{o}
	return nil
}

func (inst *myIntentTemplateRequest) doGetList() error {
	ctx := inst.gc
	ser := inst.controller.IntentTemplateService
	list, err := ser.ListAll(ctx)
	if err != nil {
		return err
	}
	inst.body2.Templates = list
	return nil
}

func (inst *myIntentTemplateRequest) doPost() error {
	ctx := inst.gc
	ser := inst.controller.IntentTemplateService
	o1 := inst.body1.Templates[0]
	o2, err := ser.Insert(ctx, o1)
	if err != nil {
		return err
	}
	inst.body2.Templates = []*dto.IntentTemplate{o2}
	return nil
}

func (inst *myIntentTemplateRequest) doPut() error {
	return nil
}

func (inst *myIntentTemplateRequest) doDelete() error {

	return nil
}
