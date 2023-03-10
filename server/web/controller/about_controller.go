package controller

import (
	"net/http"

	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/vo"
	"github.com/gin-gonic/gin"
)

// AboutController About 控制器
type AboutController struct {
	markup.RestController `class:"rest-controller"`

	AboutService service.AboutService `inject:"#AboutService"`
	Responder    glass.MainResponder  `inject:"#glass-main-responder"`
	Profile      string               `inject:"${application.profiles.active}"`
}

func (inst *AboutController) _Impl() glass.Controller {
	return inst
}

// Init 初始化
func (inst *AboutController) Init(ec glass.EngineConnection) error {

	ec = ec.RequestMapping("about")
	ec.Handle(http.MethodGet, "", inst.handleGet)
	ec.Handle(http.MethodGet, "debug-mode", inst.handleGetDebug)
	return nil
}

func (inst *AboutController) handleGet(c *gin.Context) {
	req := &myAboutRequest{
		gc:              c,
		controller:      inst,
		wantRequestID:   false,
		wantRequestBody: false,
	}
	err := req.open()
	if err == nil {
		err = req.doGet()
	}
	req.send(err)
}

func (inst *AboutController) handleGetDebug(c *gin.Context) {
	req := &myAboutRequest{
		gc:              c,
		controller:      inst,
		wantRequestID:   false,
		wantRequestBody: false,
	}
	err := req.open()
	if err == nil {
		err = req.doGetDebug()
	}
	req.send(err)
}

////////////////////////////////////////////////////////////////////////////////

type myAboutRequest struct {
	gc         *gin.Context
	controller *AboutController

	wantRequestID   bool
	wantRequestBody bool

	// id    dxo.AboutID
	body1 vo.About
	body2 vo.About
}

func (inst *myAboutRequest) open() error {

	c := inst.gc

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *myAboutRequest) send(err error) {
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

func (inst *myAboutRequest) doGet() error {
	ctx := inst.gc
	ser := inst.controller.AboutService
	view, err := ser.GetInfo(ctx)
	if err != nil {
		return err
	}
	inst.body2 = *view
	return nil
}

func (inst *myAboutRequest) doGetDebug() error {
	ctl := inst.controller
	inst.body2.Profile = ctl.Profile
	return nil
}
