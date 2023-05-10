package locations

import (
	"net/http"
	"strconv"

	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/dto"
	"github.com/bitwormhole/wpm/server/web/vo"
	"github.com/gin-gonic/gin"
)

// LocationController 控制器
type LocationController struct {
	markup.RestController `class:"rest-controller"`

	LocationService service.LocationService `inject:"#LocationService"`
	Responder       glass.MainResponder     `inject:"#glass-main-responder"`
}

func (inst *LocationController) _Impl() glass.Controller {
	return inst
}

// Init 初始化
func (inst *LocationController) Init(ec glass.EngineConnection) error {
	ec = ec.RequestMapping("locations")
	ec.Handle(http.MethodGet, "", inst.handleGetAll)
	return nil
}

func (inst *LocationController) handleGetAll(c *gin.Context) {
	req := &myLocationRequest{
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

func (inst *LocationController) handleGetOne(c *gin.Context) {
	req := &myLocationRequest{
		gc:              c,
		controller:      inst,
		wantRequestBody: false,
		wantRequestID:   true,
	}
	err := req.open()
	if err == nil {
		err = req.doGetOne()
	}
	req.send(err)
}

////////////////////////////////////////////////////////////////////////////////

type myLocationRequest struct {
	gc         *gin.Context
	controller *LocationController

	wantRequestID   bool
	wantRequestBody bool

	id    dxo.LocationID
	body1 vo.Location
	body2 vo.Location
}

func (inst *myLocationRequest) open() error {

	c := inst.gc

	if inst.wantRequestID {
		id := c.Param("id")
		n, err := strconv.Atoi(id)
		if err != nil {
			return err
		}
		inst.id = dxo.LocationID(n)
	}

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *myLocationRequest) send(err error) {
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

func (inst *myLocationRequest) doGetAll() error {
	ctx := inst.gc
	ser := inst.controller.LocationService
	opt := &service.LocationOptions{}
	list, err := ser.ListAll(ctx, opt)
	if err != nil {
		return err
	}
	inst.body2.Locations = list
	return nil
}

func (inst *myLocationRequest) doGetOne() error {
	ctx := inst.gc
	ser := inst.controller.LocationService
	opt := &service.LocationOptions{}
	id := inst.id
	o, err := ser.Find(ctx, id, opt)
	if err != nil {
		return err
	}
	inst.body2.Locations = []*dto.Location{o}
	return nil
}
