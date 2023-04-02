package contenttypes

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

// ContentTypeController ContentType 控制器
type ContentTypeController struct {
	markup.RestController `class:"rest-controller"`

	ContentTypeService service.ContentTypeService `inject:"#ContentTypeService"`
	Responder          glass.MainResponder        `inject:"#glass-main-responder"`
}

func (inst *ContentTypeController) _Impl() glass.Controller {
	return inst
}

// Init 初始化
func (inst *ContentTypeController) Init(ec glass.EngineConnection) error {

	ec = ec.RequestMapping("content-types")

	ec.Handle(http.MethodGet, "", inst.handleGetList)
	ec.Handle(http.MethodPost, "", inst.handlePost)
	ec.Handle(http.MethodGet, ":id", inst.handleGetOne)
	ec.Handle(http.MethodPut, ":id", inst.handlePut)
	ec.Handle(http.MethodDelete, ":id", inst.handleDelete)

	return nil
}

func (inst *ContentTypeController) handleGetList(c *gin.Context) {
	req := &myContentTypeRequest{
		gc:              c,
		controller:      inst,
		wantRequestID:   false,
		wantRequestBody: false,
	}
	err := req.open()
	if err == nil {
		err = req.doGetList()
	}
	req.send(err)
}

func (inst *ContentTypeController) handleGetOne(c *gin.Context) {
	req := &myContentTypeRequest{
		gc:              c,
		controller:      inst,
		wantRequestID:   true,
		wantRequestBody: false,
	}
	err := req.open()
	if err == nil {
		err = req.doGetOne()
	}
	req.send(err)
}

func (inst *ContentTypeController) handlePost(c *gin.Context) {
	req := &myContentTypeRequest{
		gc:              c,
		controller:      inst,
		wantRequestID:   false,
		wantRequestBody: true,
	}
	err := req.open()
	if err == nil {
		err = req.doPost()
	}
	req.send(err)
}

func (inst *ContentTypeController) handlePut(c *gin.Context) {
	req := &myContentTypeRequest{
		gc:              c,
		controller:      inst,
		wantRequestID:   true,
		wantRequestBody: true,
	}
	err := req.open()
	if err == nil {
		err = req.doPut()
	}
	req.send(err)
}

func (inst *ContentTypeController) handleDelete(c *gin.Context) {
	req := &myContentTypeRequest{
		gc:              c,
		controller:      inst,
		wantRequestID:   true,
		wantRequestBody: false,
	}
	err := req.open()
	if err == nil {
		err = req.doDelete()
	}
	req.send(err)
}

////////////////////////////////////////////////////////////////////////////////

type myContentTypeRequest struct {
	gc         *gin.Context
	controller *ContentTypeController

	wantRequestID   bool
	wantRequestBody bool

	id    dxo.ContentTypeID
	body1 vo.ContentType
	body2 vo.ContentType
}

func (inst *myContentTypeRequest) open() error {

	c := inst.gc

	if inst.wantRequestID {
		id := c.Param("id")
		n, err := strconv.Atoi(id)
		if err != nil {
			return err
		}
		inst.id = dxo.ContentTypeID(n)
	}

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *myContentTypeRequest) send(err error) {
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

func (inst *myContentTypeRequest) doGetOne() error {
	id := inst.id
	ctx := inst.gc
	ser := inst.controller.ContentTypeService
	o, err := ser.Find(ctx, id)
	if err != nil {
		return err
	}
	inst.body2.ContentTypes = []*dto.ContentType{o}
	return nil
}

func (inst *myContentTypeRequest) doGetList() error {
	ctx := inst.gc
	ser := inst.controller.ContentTypeService
	list, err := ser.ListAll(ctx)
	if err != nil {
		return err
	}
	inst.body2.ContentTypes = list
	return nil
}

func (inst *myContentTypeRequest) doPost() error {
	ctx := inst.gc
	ser := inst.controller.ContentTypeService
	o1 := inst.body1.ContentTypes[0]
	o2, err := ser.Insert(ctx, o1)
	if err != nil {
		return err
	}
	inst.body2.ContentTypes = []*dto.ContentType{o2}
	return nil
}

func (inst *myContentTypeRequest) doPut() error {
	ctx := inst.gc
	ser := inst.controller.ContentTypeService
	id := inst.id
	o1 := inst.body1.ContentTypes[0]
	o2, err := ser.Update(ctx, id, o1)
	if err != nil {
		return err
	}
	inst.body2.ContentTypes = []*dto.ContentType{o2}
	return nil
}

func (inst *myContentTypeRequest) doDelete() error {
	ctx := inst.gc
	ser := inst.controller.ContentTypeService
	id := inst.id
	return ser.Remove(ctx, id)
}
