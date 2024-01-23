package admin

import (
	"strconv"

	"github.com/bitwormhole/wpm/common/objects/dxo"
	"github.com/bitwormhole/wpm/common/objects/vo"
	"github.com/bitwormhole/wpm/server/classes/packages"
	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
)

// SoftwarePackageController ...
type SoftwarePackageController struct {

	//starter:component
	_as func(libgin.Controller) //starter:as(".")

	Sender  libgin.Responder //starter:inject("#")
	Service packages.Service //starter:inject("#")
}

func (inst *SoftwarePackageController) _impl() libgin.Controller {
	return inst
}

// Registration ...
func (inst *SoftwarePackageController) Registration() *libgin.ControllerRegistration {
	return &libgin.ControllerRegistration{Route: inst.route}
}

func (inst *SoftwarePackageController) route(rp libgin.RouterProxy) error {

	rp = rp.For("packages")

	rp.POST("", inst.handle)
	rp.PUT(":id", inst.handle)
	rp.DELETE(":id", inst.handle)

	rp.GET("", inst.handleGetList)
	rp.GET(":id", inst.handleGetOne)

	return nil
}

func (inst *SoftwarePackageController) handle(c *gin.Context) {
	req := &mySoftwarePackageRequest{
		context:       c,
		controller:    inst,
		wantRequestID: false,
	}
	req.execute(req.doNOP)
}

func (inst *SoftwarePackageController) handleGetOne(c *gin.Context) {
	req := &mySoftwarePackageRequest{
		context:       c,
		controller:    inst,
		wantRequestID: true,
	}
	req.execute(req.doGetOne)
}

func (inst *SoftwarePackageController) handleGetList(c *gin.Context) {
	req := &mySoftwarePackageRequest{
		context:          c,
		controller:       inst,
		wantRequestID:    false,
		wantRequestQuery: true,
	}
	req.execute(req.doGetList)
}

////////////////////////////////////////////////////////////////////////////////

type mySoftwarePackageRequest struct {
	context    *gin.Context
	controller *SoftwarePackageController

	wantRequestID    bool
	wantRequestBody  bool
	wantRequestQuery bool

	id    dxo.SoftwarePackageID
	query packages.Query

	body1 vo.SoftwarePackage
	body2 vo.SoftwarePackage
}

func (inst *mySoftwarePackageRequest) open() error {

	c := inst.context

	if inst.wantRequestID {
		idstr := c.Param("id")
		idnum, err := strconv.ParseInt(idstr, 10, 64)
		if err != nil {
			return err
		}
		inst.id = dxo.SoftwarePackageID(idnum)
	}

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *mySoftwarePackageRequest) send(err error) {
	data := &inst.body2
	code := inst.body2.Status
	resp := new(libgin.Response)
	resp.Context = inst.context
	resp.Error = err
	resp.Data = data
	resp.Status = code
	inst.controller.Sender.Send(resp)
}

func (inst *mySoftwarePackageRequest) execute(fn func() error) {
	err := inst.open()
	if err == nil {
		err = fn()
	}
	inst.send(err)
}

func (inst *mySoftwarePackageRequest) doNOP() error {
	return nil
}

func (inst *mySoftwarePackageRequest) doGetOne() error {

	return nil
}

func (inst *mySoftwarePackageRequest) doGetList() error {
	ctx := inst.context
	ser := inst.controller.Service
	q := &inst.query
	list, err := ser.List(ctx, q)
	if err != nil {
		return err
	}
	inst.body2.Packages = list
	return nil
}
