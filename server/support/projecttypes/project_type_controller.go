package projecttypes

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

// ProjectTypeController ProjectType 控制器
type ProjectTypeController struct {
	markup.RestController `class:"rest-controller"`

	ProjectTypeService service.ProjectTypeService `inject:"#ProjectTypeService"`
	Responder          glass.MainResponder        `inject:"#glass-main-responder"`
}

func (inst *ProjectTypeController) _Impl() glass.Controller {
	return inst
}

// Init 初始化
func (inst *ProjectTypeController) Init(ec glass.EngineConnection) error {

	ec = ec.RequestMapping("project-types")

	ec.Handle(http.MethodGet, "", inst.handleGetList)
	ec.Handle(http.MethodPost, "", inst.handlePost)
	ec.Handle(http.MethodGet, ":id", inst.handleGetOne)
	ec.Handle(http.MethodPut, ":id", inst.handlePut)
	ec.Handle(http.MethodDelete, ":id", inst.handleDelete)

	return nil
}

func (inst *ProjectTypeController) handleGetList(c *gin.Context) {
	req := &myProjectTypeRequest{
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

func (inst *ProjectTypeController) handleGetOne(c *gin.Context) {
	req := &myProjectTypeRequest{
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

func (inst *ProjectTypeController) handlePost(c *gin.Context) {
	req := &myProjectTypeRequest{
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

func (inst *ProjectTypeController) handlePut(c *gin.Context) {
	req := &myProjectTypeRequest{
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

func (inst *ProjectTypeController) handleDelete(c *gin.Context) {
	req := &myProjectTypeRequest{
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

type myProjectTypeRequest struct {
	gc         *gin.Context
	controller *ProjectTypeController

	wantRequestID   bool
	wantRequestBody bool

	id    dxo.ProjectTypeID
	body1 vo.ProjectType
	body2 vo.ProjectType
}

func (inst *myProjectTypeRequest) open() error {

	c := inst.gc

	if inst.wantRequestID {
		id := c.Param("id")
		n, err := strconv.Atoi(id)
		if err != nil {
			return err
		}
		inst.id = dxo.ProjectTypeID(n)
	}

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *myProjectTypeRequest) send(err error) {
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

func (inst *myProjectTypeRequest) doGetOne() error {
	id := inst.id
	ctx := inst.gc
	ser := inst.controller.ProjectTypeService
	o, err := ser.Find(ctx, id)
	if err != nil {
		return err
	}
	inst.body2.ProjectTypes = []*dto.ProjectType{o}
	return nil
}

func (inst *myProjectTypeRequest) doGetList() error {
	ctx := inst.gc
	ser := inst.controller.ProjectTypeService
	list, err := ser.ListAll(ctx)
	if err != nil {
		return err
	}
	inst.body2.ProjectTypes = list
	return nil
}

func (inst *myProjectTypeRequest) doPost() error {
	ctx := inst.gc
	ser := inst.controller.ProjectTypeService
	o1 := inst.body1.ProjectTypes[0]
	o2, err := ser.Insert(ctx, o1)
	if err != nil {
		return err
	}
	inst.body2.ProjectTypes = []*dto.ProjectType{o2}
	return nil
}

func (inst *myProjectTypeRequest) doPut() error {
	ctx := inst.gc
	ser := inst.controller.ProjectTypeService
	id := inst.id
	o1 := inst.body1.ProjectTypes[0]
	o2, err := ser.Update(ctx, id, o1)
	if err != nil {
		return err
	}
	inst.body2.ProjectTypes = []*dto.ProjectType{o2}
	return nil
}

func (inst *myProjectTypeRequest) doDelete() error {
	ctx := inst.gc
	ser := inst.controller.ProjectTypeService
	id := inst.id
	return ser.Remove(ctx, id)
}
