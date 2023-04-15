package jsonbuffer

import (
	"net/http"

	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/vo"
	"github.com/gin-gonic/gin"
)

// JSONBufferController 控制器
type JSONBufferController struct {
	markup.RestController `class:"rest-controller"`

	JSONBufferService service.JSONBufferService `inject:"#JSONBufferService"`
	Responder         glass.MainResponder       `inject:"#glass-main-responder"`
}

func (inst *JSONBufferController) _Impl() glass.Controller {
	return inst
}

// Init 初始化
func (inst *JSONBufferController) Init(ec glass.EngineConnection) error {

	ec = ec.RequestMapping("jsonbuffer")

	ec.Handle(http.MethodGet, "", inst.handleGetAll)
	ec.Handle(http.MethodPut, "", inst.handlePutItems)
	ec.Handle(http.MethodPost, "reset", inst.handlePostReset)

	return nil
}

func (inst *JSONBufferController) handlePostReset(c *gin.Context) {
	req := &myJSONBufferRequest{
		gc:              c,
		controller:      inst,
		wantRequestBody: true,
	}
	err := req.open()
	if err == nil {
		err = req.doReset()
	}
	req.send(err)
}

func (inst *JSONBufferController) handlePutItems(c *gin.Context) {
	req := &myJSONBufferRequest{
		gc:              c,
		controller:      inst,
		wantRequestBody: true,
	}
	err := req.open()
	if err == nil {
		err = req.doPutItems()
	}
	req.send(err)
}

func (inst *JSONBufferController) handleGetAll(c *gin.Context) {
	req := &myJSONBufferRequest{
		gc:              c,
		controller:      inst,
		wantRequestBody: false,
	}
	err := req.open()
	if err == nil {
		err = req.doGetAll()
	}
	req.send(err)
}

////////////////////////////////////////////////////////////////////////////////

type myJSONBufferRequest struct {
	gc         *gin.Context
	controller *JSONBufferController

	// wantRequestID   bool
	wantRequestBody bool

	// id    dxo.IntentID
	body1 vo.Online
	body2 vo.Online
}

func (inst *myJSONBufferRequest) open() error {

	c := inst.gc

	// if inst.wantRequestID {
	// 	id := c.Param("id")
	// 	n, err := strconv.Atoi(id)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	inst.id = dxo.IntentID(n)
	// }

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *myJSONBufferRequest) send(err error) {
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

func (inst *myJSONBufferRequest) doGetAll() error {
	ctx := inst.gc
	ser := inst.controller.JSONBufferService
	o, err := ser.Get(ctx)
	if err != nil {
		return err
	}
	inst.body2 = *o
	return nil
}

func (inst *myJSONBufferRequest) doPutItems() error {
	ctx := inst.gc
	ser := inst.controller.JSONBufferService
	o := &inst.body1
	return ser.Put(ctx, o)
}

func (inst *myJSONBufferRequest) doReset() error {
	ctx := inst.gc
	ser := inst.controller.JSONBufferService
	return ser.Reset(ctx)
}
