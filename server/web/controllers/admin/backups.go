package admin

import (
	"strconv"

	"github.com/bitwormhole/wpm/server/classes/backups"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/web/dto"
	"github.com/bitwormhole/wpm/server/web/vo"
	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
)

// BackupsController ...
type BackupsController struct {

	//starter:component
	_as func(libgin.Controller) //starter:as(".")

	Sender  libgin.Responder //starter:inject("#")
	Service backups.Service  //starter:inject("#")
}

func (inst *BackupsController) _impl() libgin.Controller {
	return inst
}

// Registration ...
func (inst *BackupsController) Registration() *libgin.ControllerRegistration {
	return &libgin.ControllerRegistration{Route: inst.route}
}

func (inst *BackupsController) route(rp libgin.RouterProxy) error {

	rp = rp.For("backups")

	rp.POST("", inst.handle)
	rp.PUT(":id", inst.handle)
	rp.DELETE(":id", inst.handle)

	rp.GET("", inst.handle)
	rp.GET(":id", inst.handleGetOne)

	return nil
}

func (inst *BackupsController) handle(c *gin.Context) {
	req := &myBackupsRequest{
		context:       c,
		controller:    inst,
		wantRequestID: false,
	}
	req.execute(req.doNOP)
}

func (inst *BackupsController) handleGetOne(c *gin.Context) {
	req := &myBackupsRequest{
		context:       c,
		controller:    inst,
		wantRequestID: true,
	}
	req.execute(req.doGetOne)
}

////////////////////////////////////////////////////////////////////////////////

type myBackupsRequest struct {
	context    *gin.Context
	controller *BackupsController

	wantRequestID   bool
	wantRequestBody bool

	id    dxo.BackupID
	body1 vo.Backup
	body2 vo.Backup
}

func (inst *myBackupsRequest) open() error {

	c := inst.context

	if inst.wantRequestID {
		idstr := c.Param("id")
		idnum, err := strconv.ParseInt(idstr, 10, 64)
		if err != nil {
			return err
		}
		inst.id = dxo.BackupID(idnum)
	}

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *myBackupsRequest) send(err error) {
	data := &inst.body2
	code := inst.body2.Status
	resp := new(libgin.Response)
	resp.Context = inst.context
	resp.Error = err
	resp.Data = data
	resp.Status = code
	inst.controller.Sender.Send(resp)
}

func (inst *myBackupsRequest) execute(fn func() error) {
	err := inst.open()
	if err == nil {
		err = fn()
	}
	inst.send(err)
}

func (inst *myBackupsRequest) doNOP() error {
	return nil
}

func (inst *myBackupsRequest) doGetOne() error {
	id := inst.id
	o1, err := inst.controller.Service.Find(nil, id)
	if err != nil {
		return err
	}
	o2 := &dto.Backup{
		ID: o1.ID,
		// Foo: o1.Foo,
		// Bar: o1.Bar,
	}
	inst.body2.Backups = []*dto.Backup{o2}
	return nil
}
