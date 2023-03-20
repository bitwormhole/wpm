package setup

import (
	"net/http"

	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/vo"
	"github.com/gin-gonic/gin"
)

// WpmSetupController ...
type WpmSetupController struct {
	markup.Component `class:"rest-controller"`

	Responder glass.MainResponder `inject:"#glass-main-responder"`

	SetupService service.SetupService `inject:"#SetupService"`
}

func (inst *WpmSetupController) _Impl() glass.Controller {
	return inst
}

// Init ...
func (inst *WpmSetupController) Init(ec glass.EngineConnection) error {
	ec = ec.RequestMapping("setup")
	ec.Handle(http.MethodGet, "", inst.handleGetList)
	ec.Handle(http.MethodPost, "", inst.handlePost)
	ec.Handle(http.MethodPost, "skip-all", inst.handleSkipAll)
	return nil
}

func (inst *WpmSetupController) handlePost(c *gin.Context) {
	req := &mySetupRequest{
		context:         c,
		controller:      inst,
		wantRequestBody: true,
	}
	err := req.open()
	if err == nil {
		err = req.doSetupApply()
	}
	req.send(err)
}

func (inst *WpmSetupController) handleSkipAll(c *gin.Context) {
	req := &mySetupRequest{
		context:         c,
		controller:      inst,
		wantRequestBody: true,
	}
	err := req.open()
	if err == nil {
		err = req.doSkipAll()
	}
	req.send(err)
}

func (inst *WpmSetupController) handleGetList(c *gin.Context) {
	req := &mySetupRequest{
		context:    c,
		controller: inst,
	}
	err := req.open()
	if err == nil {
		err = req.doSetupList()
	}
	req.send(err)
}

////////////////////////////////////////////////////////////////////////////////

type mySetupRequest struct {
	context    *gin.Context
	controller *WpmSetupController

	wantRequestBody bool

	body1 vo.Setup
	body2 vo.Setup
}

func (inst *mySetupRequest) open() error {

	c := inst.context

	if inst.wantRequestBody {
		err := c.Bind(&inst.body1)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *mySetupRequest) send(err error) {
	ctx := inst.context
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

func (inst *mySetupRequest) doSetupList() error {
	ctx := inst.context
	ser := inst.controller.SetupService
	list, err := ser.ListAll(ctx)
	if err != nil {
		return err
	}
	inst.body2.SetupItems = list
	return nil
}

func (inst *mySetupRequest) doSetupApply() error {
	ctx := inst.context
	ser := inst.controller.SetupService
	all := inst.body1.SetupItems
	err := ser.Apply(ctx, all)
	inst.body2.SetupItems = all
	return err
}

func (inst *mySetupRequest) doSkipAll() error {
	ctx := inst.context
	ser := inst.controller.SetupService
	return ser.SkipAll(ctx)
}
