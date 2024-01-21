package admin

import (
	"github.com/bitwormhole/wpm/common/classes/about"
	"github.com/bitwormhole/wpm/server/web/vo"
	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
)

// AboutController ...
type AboutController struct {

	//starter:component
	_as func(libgin.Controller) //starter:as(".")

	Sender  libgin.Responder //starter:inject("#")
	Service about.Service    //starter:inject("#")
}

func (inst *AboutController) _impl() libgin.Controller {
	return inst
}

// Registration ...
func (inst *AboutController) Registration() *libgin.ControllerRegistration {
	return &libgin.ControllerRegistration{Route: inst.route}
}

func (inst *AboutController) route(rp libgin.RouterProxy) error {

	rp = rp.For("about")

	rp.POST("", inst.handle)
	rp.PUT(":id", inst.handle)
	rp.DELETE(":id", inst.handle)

	rp.GET("", inst.handleGetAboutInfo)
	rp.GET(":id", inst.handleGetAboutInfo)

	return nil
}

func (inst *AboutController) handle(c *gin.Context) {
	req := &myAboutRequest{
		context:       c,
		controller:    inst,
		wantRequestID: false,
	}
	req.execute(req.doNOP)
}

func (inst *AboutController) handleGetAboutInfo(c *gin.Context) {
	req := &myAboutRequest{
		context:       c,
		controller:    inst,
		wantRequestID: false,
	}
	req.execute(req.doGetAboutInfo)
}

////////////////////////////////////////////////////////////////////////////////

type myAboutRequest struct {
	context    *gin.Context
	controller *AboutController

	wantRequestID   bool
	wantRequestBody bool

	// id    dxo.AboutID
	body1 vo.About
	body2 vo.About
}

func (inst *myAboutRequest) open() error {

	c := inst.context

	if inst.wantRequestID {
		// idstr := c.Param("id")
		// idnum, err := strconv.ParseInt(idstr, 10, 64)
		// if err != nil {
		// 	return err
		// }
		// inst.id = dxo.AboutID(idnum)
	}

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *myAboutRequest) send(err error) {
	data := &inst.body2
	code := inst.body2.Status
	resp := new(libgin.Response)
	resp.Context = inst.context
	resp.Error = err
	resp.Data = data
	resp.Status = code
	inst.controller.Sender.Send(resp)
}

func (inst *myAboutRequest) execute(fn func() error) {
	err := inst.open()
	if err == nil {
		err = fn()
	}
	inst.send(err)
}

func (inst *myAboutRequest) doNOP() error {
	return nil
}

func (inst *myAboutRequest) doGetAboutInfo() error {
	ctx := inst.context
	dst := &inst.body2
	ser := inst.controller.Service
	return ser.GetAboutInfo(ctx, dst)
}
