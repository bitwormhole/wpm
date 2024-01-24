package admin

import (
	"strconv"

	"github.com/bitwormhole/wpm/common/objects/dto"
	"github.com/bitwormhole/wpm/common/objects/dxo"
	"github.com/bitwormhole/wpm/common/objects/vo"
	"github.com/bitwormhole/wpm/server/classes/executables"
	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
)

// ExecutablesController ...
type ExecutablesController struct {

	//starter:component
	_as func(libgin.Controller) //starter:as(".")

	Sender  libgin.Responder    //starter:inject("#")
	Service executables.Service //starter:inject("#")
}

func (inst *ExecutablesController) _impl() libgin.Controller {
	return inst
}

// Registration ...
func (inst *ExecutablesController) Registration() *libgin.ControllerRegistration {
	return &libgin.ControllerRegistration{Route: inst.route}
}

func (inst *ExecutablesController) route(rp libgin.RouterProxy) error {

	rp = rp.For("executables")

	rp.POST("", inst.handlePost)
	rp.PUT(":id", inst.handlePut)
	rp.DELETE(":id", inst.handleDelete)

	rp.GET("", inst.handleGetList)
	rp.GET(":id", inst.handleGetOne)

	return nil
}

func (inst *ExecutablesController) handle(c *gin.Context) {
	req := &myExecutablesRequest{
		context:       c,
		controller:    inst,
		wantRequestID: false,
	}
	req.execute(req.doNOP)
}

func (inst *ExecutablesController) handleGetOne(c *gin.Context) {
	req := &myExecutablesRequest{
		context:       c,
		controller:    inst,
		wantRequestID: true,
	}
	req.execute(req.doGetOne)
}

func (inst *ExecutablesController) handleGetList(c *gin.Context) {
	req := &myExecutablesRequest{
		context:          c,
		controller:       inst,
		wantRequestID:    false,
		wantRequestQuery: true,
	}
	req.execute(req.doGetList)
}

func (inst *ExecutablesController) handlePost(c *gin.Context) {
	req := &myExecutablesRequest{
		context:          c,
		controller:       inst,
		wantRequestID:    false,
		wantRequestQuery: false,
		wantRequestBody:  true,
	}
	req.execute(req.doInsert)
}

func (inst *ExecutablesController) handlePut(c *gin.Context) {
	req := &myExecutablesRequest{
		context:          c,
		controller:       inst,
		wantRequestID:    true,
		wantRequestQuery: false,
		wantRequestBody:  true,
	}
	req.execute(req.doUpdate)
}

func (inst *ExecutablesController) handleDelete(c *gin.Context) {
	req := &myExecutablesRequest{
		context:          c,
		controller:       inst,
		wantRequestID:    true,
		wantRequestQuery: false,
	}
	req.execute(req.doRemove)
}

////////////////////////////////////////////////////////////////////////////////

type myExecutablesRequest struct {
	context    *gin.Context
	controller *ExecutablesController

	wantRequestID    bool
	wantRequestBody  bool
	wantRequestQuery bool

	id    dxo.ExecutableID
	query executables.Query

	body1 vo.Executable
	body2 vo.Executable
}

func (inst *myExecutablesRequest) open() error {

	c := inst.context

	if inst.wantRequestID {
		idstr := c.Param("id")
		idnum, err := strconv.ParseInt(idstr, 10, 64)
		if err != nil {
			return err
		}
		inst.id = dxo.ExecutableID(idnum)
	}

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *myExecutablesRequest) send(err error) {
	data := &inst.body2
	code := inst.body2.Status
	resp := new(libgin.Response)
	resp.Context = inst.context
	resp.Error = err
	resp.Data = data
	resp.Status = code
	inst.controller.Sender.Send(resp)
}

func (inst *myExecutablesRequest) execute(fn func() error) {
	err := inst.open()
	if err == nil {
		err = fn()
	}
	inst.send(err)
}

func (inst *myExecutablesRequest) doNOP() error {
	return nil
}

func (inst *myExecutablesRequest) doGetOne() error {
	id := inst.id
	o1, err := inst.controller.Service.Find(nil, id)
	if err != nil {
		return err
	}
	o2 := &dto.Executable{
		ID: o1.ID,
	}
	inst.body2.Executables = []*dto.Executable{o2}
	return nil
}

func (inst *myExecutablesRequest) doGetList() error {

	inst.query.All = true // todo ... workaround

	ctx := inst.context
	ser := inst.controller.Service
	q := &inst.query
	list, err := ser.List(ctx, q)
	if err != nil {
		return err
	}
	inst.body2.Executables = list
	return nil
}

func (inst *myExecutablesRequest) doInsert() error {
	ctx := inst.context
	ser := inst.controller.Service
	src := inst.body1.Executables
	dst := inst.body2.Executables
	for _, item1 := range src {
		item2, err := ser.Insert(ctx, item1)
		if err != nil {
			return err
		}
		dst = append(dst, item2)
	}
	inst.body2.Executables = dst
	return nil
}

func (inst *myExecutablesRequest) doUpdate() error {
	ctx := inst.context
	ser := inst.controller.Service
	src := inst.body1.Executables
	dst := inst.body2.Executables
	id := inst.id
	for _, item1 := range src {
		if item1.ID != id {
			continue
		}
		item2, err := ser.Update(ctx, id, item1)
		if err != nil {
			return err
		}
		dst = append(dst, item2)
	}
	inst.body2.Executables = dst
	return nil
}

func (inst *myExecutablesRequest) doRemove() error {
	ctx := inst.context
	ser := inst.controller.Service
	id := inst.id
	return ser.Remove(ctx, id)
}
