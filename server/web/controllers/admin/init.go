package admin

import (
	"strconv"

	"github.com/bitwormhole/wpm/common/objects/dxo"
	"github.com/bitwormhole/wpm/common/objects/vo"
	"github.com/bitwormhole/wpm/server/classes/examples"
	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
)

// InitController ...
type InitController struct {

	//starter:component
	_as func(libgin.Controller) //starter:as(".")

	Sender libgin.Responder //starter:inject("#")
	Dao    examples.DAO     //starter:inject("#")
}

func (inst *InitController) _impl() libgin.Controller {
	return inst
}

// Registration ...
func (inst *InitController) Registration() *libgin.ControllerRegistration {
	return &libgin.ControllerRegistration{Route: inst.route}
}

func (inst *InitController) route(rp libgin.RouterProxy) error {

	rp = rp.For("init")

	// rp.POST("", inst.handle)
	// rp.PUT(":id", inst.handle)
	// rp.DELETE(":id", inst.handle)

	rp.GET("", inst.handle)
	// rp.GET(":id", inst.handleGetOne)

	return nil
}

func (inst *InitController) handle(c *gin.Context) {
	req := &myInitRequest{
		context:       c,
		controller:    inst,
		wantRequestID: false,
	}
	req.execute(req.doNOP)
}

func (inst *InitController) handleGetOne(c *gin.Context) {
	req := &myInitRequest{
		context:       c,
		controller:    inst,
		wantRequestID: true,
	}
	req.execute(req.doGetOne)
}

////////////////////////////////////////////////////////////////////////////////

type myInitRequest struct {
	context    *gin.Context
	controller *InitController

	wantRequestID   bool
	wantRequestBody bool

	id    dxo.ExampleID
	body1 vo.Init
	body2 vo.Init
}

func (inst *myInitRequest) open() error {

	c := inst.context

	if inst.wantRequestID {
		idstr := c.Param("id")
		idnum, err := strconv.ParseInt(idstr, 10, 64)
		if err != nil {
			return err
		}
		inst.id = dxo.ExampleID(idnum)
	}

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *myInitRequest) send(err error) {
	data := &inst.body2
	code := inst.body2.Status
	resp := new(libgin.Response)
	resp.Context = inst.context
	resp.Error = err
	resp.Data = data
	resp.Status = code
	inst.controller.Sender.Send(resp)
}

func (inst *myInitRequest) execute(fn func() error) {
	err := inst.open()
	if err == nil {
		err = fn()
	}
	inst.send(err)
}

func (inst *myInitRequest) doNOP() error {
	return nil
}

func (inst *myInitRequest) doGetOne() error {

	inst.body2.Executable.Executables = nil
	inst.body2.ContentType.ContentTypes = nil
	inst.body2.Option.Options = nil
	inst.body2.Setup.SetupItems = nil
	inst.body2.CheckUpdate.Name = ""

	return nil
}
