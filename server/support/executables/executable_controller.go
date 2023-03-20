package executables

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
	ec.Handle(http.MethodPost, "", inst.handlePost)

	ec.Handle(http.MethodGet, ":id", inst.handleGetOne)
	ec.Handle(http.MethodPut, ":id", inst.handlePut)
	ec.Handle(http.MethodDelete, ":id", inst.handleDelete)

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

func (inst *ExecutableController) handlePut(c *gin.Context) {
	req := &myExecutableRequest{
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

func (inst *ExecutableController) handleDelete(c *gin.Context) {
	req := &myExecutableRequest{
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

type myExecutableRequest struct {
	gc         *gin.Context
	controller *ExecutableController

	wantRequestID   bool
	wantRequestBody bool

	id    dxo.ExecutableID
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

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	return nil
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
	o, err := ser.Find(ctx, id)
	if err != nil {
		return err
	}
	inst.body2.Executables = []*dto.Executable{o}
	return nil
}

func (inst *myExecutableRequest) doGetList() error {
	ctx := inst.gc
	ser := inst.controller.ExecutableService
	list, err := ser.ListAll(ctx)
	if err != nil {
		return err
	}
	inst.body2.Executables = list
	return nil
}

func (inst *myExecutableRequest) doPost() error {
	ctx := inst.gc
	ser := inst.controller.ExecutableService
	o1 := inst.body1.Executables[0]
	o2, err := ser.Insert(ctx, o1)
	if err != nil {
		return nil
	}
	inst.body2.Executables = append(inst.body2.Executables, o2)
	return nil
}

func (inst *myExecutableRequest) doPut() error {
	ctx := inst.gc
	ser := inst.controller.ExecutableService
	o1 := inst.body1.Executables[0]
	id := inst.id
	o2, err := ser.Update(ctx, id, o1)
	if err != nil {
		return nil
	}
	inst.body2.Executables = append(inst.body2.Executables, o2)
	return nil
}

func (inst *myExecutableRequest) doDelete() error {
	ctx := inst.gc
	ser := inst.controller.ExecutableService
	return ser.Remove(ctx, inst.id)
}
