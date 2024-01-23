package admin

import (
	"strconv"

	"github.com/bitwormhole/wpm/common/objects/dto"
	"github.com/bitwormhole/wpm/common/objects/dxo"
	"github.com/bitwormhole/wpm/common/objects/vo"
	"github.com/bitwormhole/wpm/server/classes/repositories"
	"github.com/bitwormhole/wpm/server/web"
	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
)

// LocalRepositoryController ...
type LocalRepositoryController struct {

	//starter:component
	_as func(libgin.Controller) //starter:as(".")

	Sender  libgin.Responder                    //starter:inject("#")
	Service repositories.LocalRepositoryService //starter:inject("#")
}

func (inst *LocalRepositoryController) _impl() libgin.Controller {
	return inst
}

// Registration ...
func (inst *LocalRepositoryController) Registration() *libgin.ControllerRegistration {
	return &libgin.ControllerRegistration{Route: inst.route}
}

func (inst *LocalRepositoryController) route(rp libgin.RouterProxy) error {

	rp = rp.For("local-repositories")

	rp.POST("", inst.handle)
	rp.PUT(":id", inst.handle)
	rp.DELETE(":id", inst.handle)

	rp.GET("", inst.handleGetList)
	rp.GET(":id", inst.handleGetOne)

	return nil
}

func (inst *LocalRepositoryController) handle(c *gin.Context) {
	req := &myLocalRepositoryRequest{
		context:       c,
		controller:    inst,
		wantRequestID: false,
	}
	req.execute(req.doNOP)
}

func (inst *LocalRepositoryController) handleGetOne(c *gin.Context) {
	req := &myLocalRepositoryRequest{
		context:       c,
		controller:    inst,
		wantRequestID: true,
	}
	req.execute(req.doGetOne)
}

func (inst *LocalRepositoryController) handleGetList(c *gin.Context) {
	req := &myLocalRepositoryRequest{
		context:          c,
		controller:       inst,
		wantRequestID:    false,
		wantRequestQuery: true,
	}
	req.execute(req.doGetList)
}

////////////////////////////////////////////////////////////////////////////////

type myLocalRepositoryRequest struct {
	context    *gin.Context
	controller *LocalRepositoryController

	wantRequestID    bool
	wantRequestBody  bool
	wantRequestQuery bool

	id    dxo.LocalRepositoryID
	query repositories.Query

	body1 vo.LocalRepository
	body2 vo.LocalRepository
}

func (inst *myLocalRepositoryRequest) open() error {

	c := inst.context

	if inst.wantRequestID {
		idstr := c.Param("id")
		idnum, err := strconv.ParseInt(idstr, 10, 64)
		if err != nil {
			return err
		}
		inst.id = dxo.LocalRepositoryID(idnum)
	}

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	if inst.wantRequestQuery {
		getter := web.NewQueryGetter(c).Optional()
		gex := web.NewGetterEx(getter)
		inst.query.All = gex.GetAll()
		inst.query.WithFileState = gex.GetWithFileState()
		inst.query.WithProjects = gex.GetWithProjects()
		inst.query.Pagination = getter.GetPagination()
		err := getter.Error()
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *myLocalRepositoryRequest) send(err error) {
	data := &inst.body2
	code := inst.body2.Status
	resp := new(libgin.Response)
	resp.Context = inst.context
	resp.Error = err
	resp.Data = data
	resp.Status = code
	inst.controller.Sender.Send(resp)
}

func (inst *myLocalRepositoryRequest) execute(fn func() error) {
	err := inst.open()
	if err == nil {
		err = fn()
	}
	inst.send(err)
}

func (inst *myLocalRepositoryRequest) doNOP() error {
	return nil
}

func (inst *myLocalRepositoryRequest) doGetOne() error {
	ctx := inst.context
	ser := inst.controller.Service
	id := inst.id
	item, err := ser.Find(ctx, id)
	if err != nil {
		return err
	}
	inst.body2.Repositories = []*dto.LocalRepository{item}
	return nil
}

func (inst *myLocalRepositoryRequest) doGetList() error {
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
