package admin

import (
	"strconv"

	"github.com/bitwormhole/wpm/server/classes/executables"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/web/dto"
	"github.com/bitwormhole/wpm/server/web/vo"
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

	rp.POST("", inst.handle)
	rp.PUT(":id", inst.handle)
	rp.DELETE(":id", inst.handle)

	rp.GET("", inst.handle)
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

////////////////////////////////////////////////////////////////////////////////

type myExecutablesRequest struct {
	context    *gin.Context
	controller *ExecutablesController

	wantRequestID   bool
	wantRequestBody bool

	id    dxo.ExecutableID
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
		// Foo: o1.Foo,
		// Bar: o1.Bar,
	}
	inst.body2.Executables = []*dto.Executable{o2}
	return nil
}
