package admin

import (
	"strconv"

	"github.com/bitwormhole/wpm/server/classes/media"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/web/dto"
	"github.com/bitwormhole/wpm/server/web/vo"
	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
)

// MediaController ...
type MediaController struct {

	//starter:component
	_as func(libgin.Controller) //starter:as(".")

	Sender  libgin.Responder //starter:inject("#")
	Service media.Service    //starter:inject("#")
}

func (inst *MediaController) _impl() libgin.Controller {
	return inst
}

// Registration ...
func (inst *MediaController) Registration() *libgin.ControllerRegistration {
	return &libgin.ControllerRegistration{Route: inst.route}
}

func (inst *MediaController) route(rp libgin.RouterProxy) error {

	rp = rp.For("media")

	rp.POST("", inst.handle)
	rp.PUT(":id", inst.handle)
	rp.DELETE(":id", inst.handle)

	rp.GET("", inst.handleGetList)
	rp.GET(":id", inst.handleGetOne)

	return nil
}

func (inst *MediaController) handle(c *gin.Context) {
	req := &myMediaRequest{
		context:       c,
		controller:    inst,
		wantRequestID: false,
	}
	req.execute(req.doNOP)
}

func (inst *MediaController) handleGetOne(c *gin.Context) {
	req := &myMediaRequest{
		context:       c,
		controller:    inst,
		wantRequestID: true,
	}
	req.execute(req.doGetOne)
}

func (inst *MediaController) handleGetList(c *gin.Context) {
	req := &myMediaRequest{
		context:          c,
		controller:       inst,
		wantRequestID:    false,
		wantRequestQuery: true,
	}
	req.execute(req.doGetList)
}

////////////////////////////////////////////////////////////////////////////////

type myMediaRequest struct {
	context    *gin.Context
	controller *MediaController

	wantRequestID    bool
	wantRequestBody  bool
	wantRequestQuery bool

	id    dxo.MediaID
	query media.Query

	body1 vo.Media
	body2 vo.Media
}

func (inst *myMediaRequest) open() error {

	c := inst.context

	if inst.wantRequestID {
		idstr := c.Param("id")
		idnum, err := strconv.ParseInt(idstr, 10, 64)
		if err != nil {
			return err
		}
		inst.id = dxo.MediaID(idnum)
	}

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *myMediaRequest) send(err error) {
	data := &inst.body2
	code := inst.body2.Status
	resp := new(libgin.Response)
	resp.Context = inst.context
	resp.Error = err
	resp.Data = data
	resp.Status = code
	inst.controller.Sender.Send(resp)
}

func (inst *myMediaRequest) execute(fn func() error) {
	err := inst.open()
	if err == nil {
		err = fn()
	}
	inst.send(err)
}

func (inst *myMediaRequest) doNOP() error {
	return nil
}

func (inst *myMediaRequest) doGetOne() error {
	id := inst.id
	o1, err := inst.controller.Service.Find(nil, id)
	if err != nil {
		return err
	}
	o2 := &dto.Media{
		ID: o1.ID,
	}
	inst.body2.Mediae = []*dto.Media{o2}
	return nil
}

func (inst *myMediaRequest) doGetList() error {
	ctx := inst.context
	ser := inst.controller.Service
	q := &inst.query
	list, err := ser.List(ctx, q)
	if err != nil {
		return err
	}
	inst.body2.Mediae = list
	return nil
}
