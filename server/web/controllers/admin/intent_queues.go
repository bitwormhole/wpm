package admin

import (
	"strconv"

	"github.com/bitwormhole/wpm/common/objects/dxo"
	"github.com/bitwormhole/wpm/common/objects/vo"
	"github.com/bitwormhole/wpm/server/classes/intents"
	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
)

// IntentQueueController ...
type IntentQueueController struct {

	//starter:component
	_as func(libgin.Controller) //starter:as(".")

	Sender libgin.Responder //starter:inject("#")
	Queues intents.Queues   //starter:inject("#")
}

func (inst *IntentQueueController) _impl() libgin.Controller {
	return inst
}

// Registration ...
func (inst *IntentQueueController) Registration() *libgin.ControllerRegistration {
	return &libgin.ControllerRegistration{Route: inst.route}
}

func (inst *IntentQueueController) route(rp libgin.RouterProxy) error {

	rp = rp.For("intent-queues")

	rp.GET("", inst.handleGetExample)
	rp.DELETE(":id", inst.handleCloseChannel) // close channel

	rp.POST("", inst.handleOpenChannel)              // open channel
	rp.POST("push", inst.handlePushIntent)           // write intent
	rp.POST(":id/fetch", inst.handleListenAtChannel) // read intent

	return nil
}

func (inst *IntentQueueController) handle(c *gin.Context) {
	req := &myIntentQueueRequest{
		context:       c,
		controller:    inst,
		wantRequestID: false,
	}
	req.execute(req.doNOP)
}

func (inst *IntentQueueController) handleGetExample(c *gin.Context) {
	req := &myIntentQueueRequest{
		context:         c,
		controller:      inst,
		wantRequestID:   false,
		wantRequestBody: false,
	}
	req.execute(req.doGetExample)
}

func (inst *IntentQueueController) handleOpenChannel(c *gin.Context) {
	req := &myIntentQueueRequest{
		context:         c,
		controller:      inst,
		wantRequestID:   false,
		wantRequestBody: true,
	}
	req.execute(req.doOpenChannel)
}

func (inst *IntentQueueController) handleCloseChannel(c *gin.Context) {
	req := &myIntentQueueRequest{
		context:         c,
		controller:      inst,
		wantRequestID:   true,
		wantRequestBody: false,
	}
	req.execute(req.doCloseChannel)
}

func (inst *IntentQueueController) handleListenAtChannel(c *gin.Context) {
	req := &myIntentQueueRequest{
		context:         c,
		controller:      inst,
		wantRequestID:   true,
		wantRequestBody: true,
	}
	req.execute(req.doListenAtChannel)
}

func (inst *IntentQueueController) handlePushIntent(c *gin.Context) {
	req := &myIntentQueueRequest{
		context:         c,
		controller:      inst,
		wantRequestID:   false,
		wantRequestBody: true,
	}
	req.execute(req.doPush)
}

////////////////////////////////////////////////////////////////////////////////

type myIntentQueueRequest struct {
	context    *gin.Context
	controller *IntentQueueController

	wantRequestID   bool
	wantRequestBody bool

	id    dxo.IntentQueueID
	body1 vo.IntentQueue
	body2 vo.IntentQueue
}

func (inst *myIntentQueueRequest) open() error {

	c := inst.context

	if inst.wantRequestID {
		idstr := c.Param("id")
		idnum, err := strconv.ParseInt(idstr, 10, 64)
		if err != nil {
			return err
		}
		inst.id = dxo.IntentQueueID(idnum)
	}

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *myIntentQueueRequest) send(err error) {
	data := &inst.body2
	code := inst.body2.Status
	resp := new(libgin.Response)
	resp.Context = inst.context
	resp.Error = err
	resp.Data = data
	resp.Status = code
	inst.controller.Sender.Send(resp)
}

func (inst *myIntentQueueRequest) execute(fn func() error) {
	err := inst.open()
	if err == nil {
		err = fn()
	}
	inst.send(err)
}

func (inst *myIntentQueueRequest) doNOP() error {
	return nil
}

func (inst *myIntentQueueRequest) doGetExample() error {
	return nil
}

func (inst *myIntentQueueRequest) doOpenChannel() error {
	ctx := inst.context
	ser := inst.controller.Queues
	item1 := &inst.body1.Queue
	item2, err := ser.Open(ctx, item1)
	if err != nil {
		return err
	}
	inst.body2.Queue = *item2
	return nil
}

func (inst *myIntentQueueRequest) doCloseChannel() error {
	ctx := inst.context
	ser := inst.controller.Queues
	id := inst.id
	return ser.Close(ctx, id)
}

func (inst *myIntentQueueRequest) doListenAtChannel() error {
	ctx := inst.context
	ser := inst.controller.Queues
	id := inst.id
	iq, err := ser.Listen(ctx, id)
	if err != nil {
		return err
	}
	inst.body2.Queue = *iq
	return nil
}

func (inst *myIntentQueueRequest) doPush() error {
	ctx := inst.context
	ser := inst.controller.Queues
	q := &inst.body1.Queue
	return ser.Push(ctx, q)
}
