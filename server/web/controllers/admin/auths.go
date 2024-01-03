package admin

import (
	"github.com/bitwormhole/wpm/server/classes/auths"
	"github.com/bitwormhole/wpm/server/web/vo"
	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
)

// AuthsController ...
type AuthsController struct {

	//starter:component
	_as func(libgin.Controller) //starter:as(".")

	Sender  libgin.Responder //starter:inject("#")
	Service auths.Service    //starter:inject("#")
}

func (inst *AuthsController) _impl() libgin.Controller {
	return inst
}

// Registration ...
func (inst *AuthsController) Registration() *libgin.ControllerRegistration {
	return &libgin.ControllerRegistration{Route: inst.route}
}

func (inst *AuthsController) route(rp libgin.RouterProxy) error {

	rp = rp.For("auths")

	rp.POST("", inst.handle)
	rp.PUT(":id", inst.handle)
	rp.DELETE(":id", inst.handle)

	rp.GET("", inst.handle)
	rp.GET(":id", inst.handleGetOne)

	return nil
}

func (inst *AuthsController) handle(c *gin.Context) {
	req := &myAuthsRequest{
		context:       c,
		controller:    inst,
		wantRequestID: false,
	}
	req.execute(req.doNOP)
}

func (inst *AuthsController) handleGetOne(c *gin.Context) {
	req := &myAuthsRequest{
		context:       c,
		controller:    inst,
		wantRequestID: true,
	}
	req.execute(req.doGetOne)
}

////////////////////////////////////////////////////////////////////////////////

type myAuthsRequest struct {
	context    *gin.Context
	controller *AuthsController

	wantRequestID   bool
	wantRequestBody bool

	// id    dxo.AuthsID
	body1 vo.Auth
	body2 vo.Auth
}

func (inst *myAuthsRequest) open() error {

	c := inst.context

	if inst.wantRequestID {
		// idstr := c.Param("id")
		// idnum, err := strconv.ParseInt(idstr, 10, 64)
		// if err != nil {
		// 	return err
		// }
		// inst.id = dxo.AuthsID(idnum)
	}

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *myAuthsRequest) send(err error) {
	data := &inst.body2
	code := inst.body2.Status
	resp := new(libgin.Response)
	resp.Context = inst.context
	resp.Error = err
	resp.Data = data
	resp.Status = code
	inst.controller.Sender.Send(resp)
}

func (inst *myAuthsRequest) execute(fn func() error) {
	err := inst.open()
	if err == nil {
		err = fn()
	}
	inst.send(err)
}

func (inst *myAuthsRequest) doNOP() error {
	return nil
}

func (inst *myAuthsRequest) doGetOne() error {

	return nil
}
