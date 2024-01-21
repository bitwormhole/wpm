package admin

import (
	"strconv"

	"github.com/bitwormhole/wpm/server/classes/softwaresets"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/web/vo"
	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
)

// SoftwareSetController ...
type SoftwareSetController struct {

	//starter:component
	_as func(libgin.Controller) //starter:as(".")

	Sender  libgin.Responder     //starter:inject("#")
	Service softwaresets.Service //starter:inject("#")
}

func (inst *SoftwareSetController) _impl() libgin.Controller {
	return inst
}

// Registration ...
func (inst *SoftwareSetController) Registration() *libgin.ControllerRegistration {
	return &libgin.ControllerRegistration{Route: inst.route}
}

func (inst *SoftwareSetController) route(rp libgin.RouterProxy) error {

	rp = rp.For("software-sets")

	rp.POST("", inst.handle)
	rp.PUT(":id", inst.handle)
	rp.DELETE(":id", inst.handle)

	rp.GET(":id", inst.handleGetOne)
	rp.GET("", inst.handleGetList)

	return nil
}

func (inst *SoftwareSetController) handle(c *gin.Context) {
	req := &mySoftwareSetRequest{
		context:       c,
		controller:    inst,
		wantRequestID: false,
	}
	req.execute(req.doNOP)
}

func (inst *SoftwareSetController) handleGetOne(c *gin.Context) {
	req := &mySoftwareSetRequest{
		context:       c,
		controller:    inst,
		wantRequestID: true,
	}
	req.execute(req.doGetOne)
}

func (inst *SoftwareSetController) handleGetList(c *gin.Context) {
	req := &mySoftwareSetRequest{
		context:          c,
		controller:       inst,
		wantRequestID:    false,
		wantRequestQuery: true,
	}
	req.execute(req.doGetList)
}

////////////////////////////////////////////////////////////////////////////////

type mySoftwareSetRequest struct {
	context    *gin.Context
	controller *SoftwareSetController

	wantRequestID    bool
	wantRequestBody  bool
	wantRequestQuery bool

	id    dxo.SoftwareSetID
	query softwaresets.Query

	body1 vo.SoftwareSet
	body2 vo.SoftwareSet
}

func (inst *mySoftwareSetRequest) open() error {

	c := inst.context

	if inst.wantRequestID {
		idstr := c.Param("id")
		idnum, err := strconv.ParseInt(idstr, 10, 64)
		if err != nil {
			return err
		}
		inst.id = dxo.SoftwareSetID(idnum)
	}

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *mySoftwareSetRequest) send(err error) {
	data := &inst.body2
	code := inst.body2.Status
	resp := new(libgin.Response)
	resp.Context = inst.context
	resp.Error = err
	resp.Data = data
	resp.Status = code
	inst.controller.Sender.Send(resp)
}

func (inst *mySoftwareSetRequest) execute(fn func() error) {
	err := inst.open()
	if err == nil {
		err = fn()
	}
	inst.send(err)
}

func (inst *mySoftwareSetRequest) doNOP() error {
	return nil
}

func (inst *mySoftwareSetRequest) doGetOne() error {

	return nil
}

func (inst *mySoftwareSetRequest) doGetList() error {
	ctx := inst.context
	ser := inst.controller.Service
	q := &inst.query
	list, err := ser.List(ctx, q)
	if err != nil {
		return err
	}
	inst.body2.Sets = list
	return nil
}
