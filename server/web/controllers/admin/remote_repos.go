package admin

import (
	"strconv"

	"github.com/bitwormhole/wpm/common/objects/dto"
	"github.com/bitwormhole/wpm/common/objects/dxo"
	"github.com/bitwormhole/wpm/common/objects/vo"
	"github.com/bitwormhole/wpm/server/classes/repositories"
	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
)

// RemoteRepositoryController ...
type RemoteRepositoryController struct {

	//starter:component
	_as func(libgin.Controller) //starter:as(".")

	Sender  libgin.Responder                     //starter:inject("#")
	Service repositories.RemoteRepositoryService //starter:inject("#")
}

func (inst *RemoteRepositoryController) _impl() libgin.Controller {
	return inst
}

// Registration ...
func (inst *RemoteRepositoryController) Registration() *libgin.ControllerRegistration {
	return &libgin.ControllerRegistration{Route: inst.route}
}

func (inst *RemoteRepositoryController) route(rp libgin.RouterProxy) error {

	rp = rp.For("remote-repositories")

	rp.POST("", inst.handle)
	rp.PUT(":id", inst.handle)
	rp.DELETE(":id", inst.handle)

	rp.GET("", inst.handleGetList)
	rp.GET(":id", inst.handleGetOne)

	return nil
}

func (inst *RemoteRepositoryController) handle(c *gin.Context) {
	req := &myRemoteRepositoryRequest{
		context:       c,
		controller:    inst,
		wantRequestID: false,
	}
	req.execute(req.doNOP)
}

func (inst *RemoteRepositoryController) handleGetOne(c *gin.Context) {
	req := &myRemoteRepositoryRequest{
		context:       c,
		controller:    inst,
		wantRequestID: true,
	}
	req.execute(req.doGetOne)
}

func (inst *RemoteRepositoryController) handleGetList(c *gin.Context) {
	req := &myRemoteRepositoryRequest{
		context:          c,
		controller:       inst,
		wantRequestID:    false,
		wantRequestBody:  false,
		wantRequestQuery: true,
	}
	req.execute(req.doGetList)
}

////////////////////////////////////////////////////////////////////////////////

type myRemoteRepositoryRequest struct {
	context    *gin.Context
	controller *RemoteRepositoryController

	wantRequestID    bool
	wantRequestBody  bool
	wantRequestQuery bool

	id    dxo.RemoteRepositoryID
	query repositories.Query

	body1 vo.RemoteRepository
	body2 vo.RemoteRepository
}

func (inst *myRemoteRepositoryRequest) open() error {

	c := inst.context

	if inst.wantRequestID {
		idstr := c.Param("id")
		idnum, err := strconv.ParseInt(idstr, 10, 64)
		if err != nil {
			return err
		}
		inst.id = dxo.RemoteRepositoryID(idnum)
	}

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *myRemoteRepositoryRequest) send(err error) {
	data := &inst.body2
	code := inst.body2.Status
	resp := new(libgin.Response)
	resp.Context = inst.context
	resp.Error = err
	resp.Data = data
	resp.Status = code
	inst.controller.Sender.Send(resp)
}

func (inst *myRemoteRepositoryRequest) execute(fn func() error) {
	err := inst.open()
	if err == nil {
		err = fn()
	}
	inst.send(err)
}

func (inst *myRemoteRepositoryRequest) doNOP() error {
	return nil
}

func (inst *myRemoteRepositoryRequest) doGetOne() error {
	id := inst.id
	o1, err := inst.controller.Service.Find(nil, id)
	if err != nil {
		return err
	}
	o2 := &dto.RemoteRepository{
		ID: o1.ID,
	}
	inst.body2.Repositories = []*dto.RemoteRepository{o2}
	return nil
}

func (inst *myRemoteRepositoryRequest) doGetList() error {
	ctx := inst.context
	ser := inst.controller.Service
	q := &inst.query
	list, err := ser.List(ctx, q)
	if err != nil {
		return err
	}
	inst.body2.Repositories = list
	return nil
}