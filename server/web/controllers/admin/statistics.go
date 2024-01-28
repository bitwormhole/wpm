package admin

import (
	"github.com/bitwormhole/wpm/common/objects/vo"
	"github.com/bitwormhole/wpm/server/classes/statistics"
	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
)

// pi/v1/statistics"

// StatisticsController ...
type StatisticsController struct {

	//starter:component
	_as func(libgin.Controller) //starter:as(".")

	Sender  libgin.Responder   //starter:inject("#")
	Service statistics.Service //starter:inject("#")
}

func (inst *StatisticsController) _impl() libgin.Controller {
	return inst
}

// Registration ...
func (inst *StatisticsController) Registration() *libgin.ControllerRegistration {
	return &libgin.ControllerRegistration{Route: inst.route}
}

func (inst *StatisticsController) route(rp libgin.RouterProxy) error {

	rp = rp.For("statistics")

	// rp.POST("", inst.handle)
	// rp.PUT(":id", inst.handle)
	// rp.DELETE(":id", inst.handle)
	// rp.GET(":id", inst.handleGetOne)

	rp.GET("", inst.handleGet)

	return nil
}

func (inst *StatisticsController) handle(c *gin.Context) {
	req := &myStatisticsRequest{
		context:       c,
		controller:    inst,
		wantRequestID: false,
	}
	req.execute(req.doNOP)
}

func (inst *StatisticsController) handleGet(c *gin.Context) {
	req := &myStatisticsRequest{
		context:       c,
		controller:    inst,
		wantRequestID: true,
	}
	req.execute(req.doGetStatisticResult)
}

////////////////////////////////////////////////////////////////////////////////

type myStatisticsRequest struct {
	context    *gin.Context
	controller *StatisticsController

	wantRequestID   bool
	wantRequestBody bool

	// id    dxo.StatisticsID
	body1 vo.Statistic
	body2 vo.Statistic
}

func (inst *myStatisticsRequest) open() error {

	c := inst.context

	if inst.wantRequestID {
		// idstr := c.Param("id")
		// idnum, err := strconv.ParseInt(idstr, 10, 64)
		// if err != nil {
		// 	return err
		// }
		// inst.id = dxo.StatisticsID(idnum)
	}

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *myStatisticsRequest) send(err error) {
	data := &inst.body2
	code := inst.body2.Status
	resp := new(libgin.Response)
	resp.Context = inst.context
	resp.Error = err
	resp.Data = data
	resp.Status = code
	inst.controller.Sender.Send(resp)
}

func (inst *myStatisticsRequest) execute(fn func() error) {
	err := inst.open()
	if err == nil {
		err = fn()
	}
	inst.send(err)
}

func (inst *myStatisticsRequest) doNOP() error {
	return nil
}

func (inst *myStatisticsRequest) doGetStatisticResult() error {

	ctx := inst.context
	ser := inst.controller.Service
	body := &inst.body2

	return ser.GetResult(ctx, body)
}
