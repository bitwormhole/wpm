package init

import (
	"net/http"

	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/vo"
	"github.com/gin-gonic/gin"
)

// WpmInitController Init 控制器
type WpmInitController struct {
	markup.RestController `class:"rest-controller"`

	InitService service.InitService `inject:"#InitService"`
	Responder   glass.MainResponder `inject:"#glass-main-responder"`
}

func (inst *WpmInitController) _Impl() glass.Controller {
	return inst
}

// Init 初始化
func (inst *WpmInitController) Init(ec glass.EngineConnection) error {

	ec = ec.RequestMapping("init")

	ec.Handle(http.MethodGet, "", inst.handleGet)
	ec.Handle(http.MethodPost, "", inst.handlePost)

	// ec.Handle(http.MethodGet, ":id", inst.handleGetOne)
	// ec.Handle(http.MethodPut, ":id", inst.handlePut)
	// ec.Handle(http.MethodDelete, ":id", inst.handleDelete)

	return nil
}

func (inst *WpmInitController) handleGet(c *gin.Context) {
	req := &myInitRequest{
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

func (inst *WpmInitController) handlePost(c *gin.Context) {
	req := &myInitRequest{
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

////////////////////////////////////////////////////////////////////////////////

type myInitRequest struct {
	gc         *gin.Context
	controller *WpmInitController

	wantRequestID   bool
	wantRequestBody bool

	// id    dxo.InitID
	body1 vo.Init
	body2 vo.Init
}

func (inst *myInitRequest) open() error {

	c := inst.gc

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *myInitRequest) send(err error) {
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

func (inst *myInitRequest) doGet() error {
	ctx := inst.gc
	ser := inst.controller.InitService
	o2, err := ser.InitGet(ctx)
	if err != nil {
		return nil
	}
	inst.body2 = *o2
	return nil
}

func (inst *myInitRequest) doPost() error {
	ctx := inst.gc
	ser := inst.controller.InitService
	o1 := &inst.body1
	o2, err := ser.InitSet(ctx, o1)
	if err != nil {
		return nil
	}
	inst.body2 = *o2
	return nil
}
