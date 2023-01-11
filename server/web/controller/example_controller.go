package controller

import (
	"fmt"
	"strconv"

	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/dto"
	"github.com/bitwormhole/wpm/server/web/vo"
	"github.com/gin-gonic/gin"
)

// ExampleController Example 控制器
type ExampleController struct {
	markup.RestController `class:"rest-controller"`

	ExampleService service.ExampleService // `inject:"#ExampleService"`
	Responder      glass.MainResponder    `inject:"#glass-main-responder"`
}

func (inst *ExampleController) _Impl() glass.Controller {
	return inst
}

// Init 初始化
func (inst *ExampleController) Init(ec glass.EngineConnection) error {

	ec = ec.RequestMapping("Examples")

	// ec.Handle(http.MethodGet, "", inst.handleGetList)
	// ec.Handle(http.MethodPost, "", inst.handlePost)
	// ec.Handle(http.MethodGet, ":id", inst.handleGetOne)
	// ec.Handle(http.MethodPut, ":id", inst.handlePut)
	// ec.Handle(http.MethodDelete, ":id", inst.handleDelete)

	return nil
}

func (inst *ExampleController) handleGetList(c *gin.Context) {
	req := &myExampleRequest{
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

func (inst *ExampleController) handleGetOne(c *gin.Context) {
	req := &myExampleRequest{
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

func (inst *ExampleController) handlePost(c *gin.Context) {
	req := &myExampleRequest{
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

func (inst *ExampleController) handlePut(c *gin.Context) {
	req := &myExampleRequest{
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

func (inst *ExampleController) handleDelete(c *gin.Context) {
	req := &myExampleRequest{
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

type myExampleRequest struct {
	gc         *gin.Context
	controller *ExampleController

	wantRequestID   bool
	wantRequestBody bool

	id    dxo.ExampleID
	body1 vo.Example
	body2 vo.Example
}

func (inst *myExampleRequest) open() error {

	c := inst.gc

	if inst.wantRequestID {
		id := c.Param("id")
		n, err := strconv.Atoi(id)
		if err != nil {
			return err
		}
		inst.id = dxo.ExampleID(n)
	}

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *myExampleRequest) send(err error) {
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

func (inst *myExampleRequest) doGetOne() error {
	id := inst.id
	ctx := inst.gc
	ser := inst.controller.ExampleService
	o, err := ser.Find(ctx, id)
	if err != nil {
		return err
	}
	inst.body2.Examples = []*dto.Example{o}
	return nil
}

func (inst *myExampleRequest) doGetList() error {
	ctx := inst.gc
	ser := inst.controller.ExampleService
	list, err := ser.ListAll(ctx)
	if err != nil {
		return err
	}
	inst.body2.Examples = list
	return nil
}

func (inst *myExampleRequest) doPost() error {

	return fmt.Errorf("no impl")
}

func (inst *myExampleRequest) doPut() error {

	return fmt.Errorf("no impl")
}

func (inst *myExampleRequest) doDelete() error {

	return fmt.Errorf("no impl")
}
