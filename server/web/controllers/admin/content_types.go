package admin

import (
	"strconv"

	"github.com/bitwormhole/wpm/server/classes/contenttypes"
	"github.com/bitwormhole/wpm/server/web"

	"github.com/bitwormhole/wpm/common/objects/dto"
	"github.com/bitwormhole/wpm/common/objects/dxo"
	"github.com/bitwormhole/wpm/common/objects/vo"
	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
)

// ContentTypesController ...
type ContentTypesController struct {

	//starter:component
	_as func(libgin.Controller) //starter:as(".")

	Sender  libgin.Responder     //starter:inject("#")
	Service contenttypes.Service //starter:inject("#")
}

func (inst *ContentTypesController) _impl() libgin.Controller {
	return inst
}

// Registration ...
func (inst *ContentTypesController) Registration() *libgin.ControllerRegistration {
	return &libgin.ControllerRegistration{Route: inst.route}
}

func (inst *ContentTypesController) route(rp libgin.RouterProxy) error {

	rp = rp.For("content-types")

	rp.POST("", inst.handlePost)
	rp.PUT(":id", inst.handlePut)
	rp.DELETE(":id", inst.handleDelete)

	rp.GET("", inst.handleGetList)
	rp.GET(":id", inst.handleGetOne)

	return nil
}

func (inst *ContentTypesController) handle(c *gin.Context) {
	req := &myContentTypesRequest{
		context:       c,
		controller:    inst,
		wantRequestID: false,
	}
	req.execute(req.doNOP)
}

func (inst *ContentTypesController) handleGetOne(c *gin.Context) {
	req := &myContentTypesRequest{
		context:       c,
		controller:    inst,
		wantRequestID: true,
	}
	req.execute(req.doGetOne)
}

func (inst *ContentTypesController) handleGetList(c *gin.Context) {
	req := &myContentTypesRequest{
		context:          c,
		controller:       inst,
		wantRequestID:    false,
		wantRequestQuery: true,
	}
	req.execute(req.doGetList)
}

func (inst *ContentTypesController) handlePost(c *gin.Context) {
	req := &myContentTypesRequest{
		context:         c,
		controller:      inst,
		wantRequestID:   false,
		wantRequestBody: true,
	}
	req.execute(req.doInsert)
}

func (inst *ContentTypesController) handlePut(c *gin.Context) {
	req := &myContentTypesRequest{
		context:          c,
		controller:       inst,
		wantRequestID:    true,
		wantRequestQuery: false,
		wantRequestBody:  true,
	}
	req.execute(req.doUpdate)
}

func (inst *ContentTypesController) handleDelete(c *gin.Context) {
	req := &myContentTypesRequest{
		context:          c,
		controller:       inst,
		wantRequestID:    true,
		wantRequestQuery: false,
	}
	req.execute(req.doRemove)
}

////////////////////////////////////////////////////////////////////////////////

type myContentTypesRequest struct {
	context    *gin.Context
	controller *ContentTypesController

	wantRequestID    bool
	wantRequestBody  bool
	wantRequestQuery bool

	id    dxo.ContentTypeID
	query contenttypes.Query

	body1 vo.ContentType
	body2 vo.ContentType
}

func (inst *myContentTypesRequest) open() error {

	c := inst.context

	if inst.wantRequestID {
		idstr := c.Param("id")
		idnum, err := strconv.ParseInt(idstr, 10, 64)
		if err != nil {
			return err
		}
		inst.id = dxo.ContentTypeID(idnum)
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
		// inst.query.WithFileState = gex.GetWithFileState()
		// inst.query.WithProjects = gex.GetWithProjects()
		inst.query.Pagination = getter.GetPagination()
		err := getter.Error()
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *myContentTypesRequest) send(err error) {
	data := &inst.body2
	code := inst.body2.Status
	resp := new(libgin.Response)
	resp.Context = inst.context
	resp.Error = err
	resp.Data = data
	resp.Status = code
	inst.controller.Sender.Send(resp)
}

func (inst *myContentTypesRequest) execute(fn func() error) {
	err := inst.open()
	if err == nil {
		err = fn()
	}
	inst.send(err)
}

func (inst *myContentTypesRequest) doNOP() error {
	return nil
}

func (inst *myContentTypesRequest) doGetOne() error {
	id := inst.id
	o1, err := inst.controller.Service.Find(nil, id)
	if err != nil {
		return err
	}
	o2 := &dto.ContentType{
		ID: o1.ID,
		// Foo: o1.Foo,
		// Bar: o1.Bar,
	}
	inst.body2.ContentTypes = []*dto.ContentType{o2}
	return nil
}

func (inst *myContentTypesRequest) doGetList() error {

	inst.query.All = true // todo ... workaround

	ctx := inst.context
	ser := inst.controller.Service
	q := &inst.query
	list, err := ser.List(ctx, q)
	if err != nil {
		return err
	}
	inst.body2.ContentTypes = list
	return nil
}

func (inst *myContentTypesRequest) doInsert() error {
	ctx := inst.context
	ser := inst.controller.Service
	src := inst.body1.ContentTypes
	dst := inst.body2.ContentTypes
	for _, item1 := range src {
		item2, err := ser.Insert(ctx, item1)
		if err != nil {
			return err
		}
		dst = append(dst, item2)
	}
	inst.body2.ContentTypes = dst
	return nil
}

func (inst *myContentTypesRequest) doUpdate() error {
	ctx := inst.context
	ser := inst.controller.Service
	src := inst.body1.ContentTypes
	dst := inst.body2.ContentTypes
	id := inst.id
	for _, item1 := range src {
		if item1.ID == id {
			item2, err := ser.Update(ctx, id, item1)
			if err != nil {
				return err
			}
			dst = append(dst, item2)
		}
	}
	inst.body2.ContentTypes = dst
	return nil
}

func (inst *myContentTypesRequest) doRemove() error {
	ctx := inst.context
	ser := inst.controller.Service
	id := inst.id
	return ser.Remove(ctx, id)
}
