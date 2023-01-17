package controller

import (
	"net/http"
	"strconv"

	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/vo"
	"github.com/gin-gonic/gin"
)

// PlatformController Platform 控制器
type PlatformController struct {
	markup.RestController `class:"rest-controller"`

	PlatformService service.PlatformService `inject:"#PlatformService"`
	Responder       glass.MainResponder     `inject:"#glass-main-responder"`
}

func (inst *PlatformController) _Impl() glass.Controller {
	return inst
}

// Init 初始化
func (inst *PlatformController) Init(ec glass.EngineConnection) error {

	ec = ec.RequestMapping("profiles")

	ec.Handle(http.MethodGet, "", inst.handleGetOne)
	ec.Handle(http.MethodGet, ":id", inst.handleGetOne)

	// ec.Handle(http.MethodPost, "", inst.handlePost)
	// ec.Handle(http.MethodGet, ":id", inst.handleGetOne)
	// ec.Handle(http.MethodPut, ":id", inst.handlePut)
	// ec.Handle(http.MethodDelete, ":id", inst.handleDelete)

	return nil
}

func (inst *PlatformController) handleGetOne(c *gin.Context) {
	req := &myPlatformRequest{
		gc:              c,
		controller:      inst,
		wantRequestID:   false,
		wantRequestBody: false,
	}
	err := req.open()
	if err == nil {
		err = req.doGetOne()
	}
	req.send(err)
}

////////////////////////////////////////////////////////////////////////////////

type myPlatformRequest struct {
	gc         *gin.Context
	controller *PlatformController

	wantRequestID   bool
	wantRequestBody bool

	id    dxo.PlatformID
	body1 vo.Profile
	body2 vo.Profile
}

func (inst *myPlatformRequest) open() error {

	c := inst.gc

	if inst.wantRequestID {
		id := c.Param("id")
		n, err := strconv.Atoi(id)
		if err != nil {
			return err
		}
		inst.id = dxo.PlatformID(n)
	}

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *myPlatformRequest) send(err error) {
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

func (inst *myPlatformRequest) doGetOne() error {
	ser := inst.controller.PlatformService

	o1, err := ser.GetProfile()
	if err != nil {
		return err
	}

	o2, err := ser.GetPlatform()
	if err != nil {
		return err
	}

	inst.body2.Profile = o1
	inst.body2.Platform = o2
	return nil
}
