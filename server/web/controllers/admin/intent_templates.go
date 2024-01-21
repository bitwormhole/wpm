package admin

import (
	"strconv"

	"github.com/bitwormhole/wpm/server/classes/intenttemplates"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/web/vo"
	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
)

// IntentTemplateController ...
type IntentTemplateController struct {

	//starter:component
	_as func(libgin.Controller) //starter:as(".")

	Sender  libgin.Responder        //starter:inject("#")
	Service intenttemplates.Service //starter:inject("#")
}

func (inst *IntentTemplateController) _impl() libgin.Controller {
	return inst
}

// Registration ...
func (inst *IntentTemplateController) Registration() *libgin.ControllerRegistration {
	return &libgin.ControllerRegistration{Route: inst.route}
}

func (inst *IntentTemplateController) route(rp libgin.RouterProxy) error {

	rp = rp.For("intent-templates")

	rp.POST("", inst.handle)
	rp.PUT(":id", inst.handle)
	rp.DELETE(":id", inst.handle)

	rp.GET("", inst.handleGetList)
	rp.GET(":id", inst.handleGetOne)

	return nil
}

func (inst *IntentTemplateController) handle(c *gin.Context) {
	req := &myIntentTemplateRequest{
		context:       c,
		controller:    inst,
		wantRequestID: false,
	}
	req.execute(req.doNOP)
}

func (inst *IntentTemplateController) handleGetOne(c *gin.Context) {
	req := &myIntentTemplateRequest{
		context:       c,
		controller:    inst,
		wantRequestID: true,
	}
	req.execute(req.doGetOne)
}

func (inst *IntentTemplateController) handleGetList(c *gin.Context) {
	req := &myIntentTemplateRequest{
		context:          c,
		controller:       inst,
		wantRequestID:    false,
		wantRequestQuery: true,
	}
	req.execute(req.doGetList)
}

////////////////////////////////////////////////////////////////////////////////

type myIntentTemplateRequest struct {
	context    *gin.Context
	controller *IntentTemplateController

	wantRequestID    bool
	wantRequestBody  bool
	wantRequestQuery bool

	id    dxo.IntentTemplateID
	query intenttemplates.Query

	body1 vo.IntentTemplate
	body2 vo.IntentTemplate
}

func (inst *myIntentTemplateRequest) open() error {

	c := inst.context

	if inst.wantRequestID {
		idstr := c.Param("id")
		idnum, err := strconv.ParseInt(idstr, 10, 64)
		if err != nil {
			return err
		}
		inst.id = dxo.IntentTemplateID(idnum)
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
	data := &inst.body2
	code := inst.body2.Status
	resp := new(libgin.Response)
	resp.Context = inst.context
	resp.Error = err
	resp.Data = data
	resp.Status = code
	inst.controller.Sender.Send(resp)
}

func (inst *myIntentTemplateRequest) execute(fn func() error) {
	err := inst.open()
	if err == nil {
		err = fn()
	}
	inst.send(err)
}

func (inst *myIntentTemplateRequest) doNOP() error {
	return nil
}

func (inst *myIntentTemplateRequest) doGetOne() error {

	return nil
}

func (inst *myIntentTemplateRequest) doGetList() error {
	ctx := inst.context
	ser := inst.controller.Service
	q := &inst.query
	list, err := ser.List(ctx, q)
	if err != nil {
		return err
	}
	inst.body2.Templates = list
	return nil
}
