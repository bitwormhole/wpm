package controller

import (
	"net/http"
	"strconv"

	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/app/data/dxo"
	"github.com/bitwormhole/wpm/app/service"
	"github.com/bitwormhole/wpm/app/web/dto"
	"github.com/bitwormhole/wpm/app/web/vo"
	"github.com/gin-gonic/gin"
)

// ProjectController 仓库控制器
type ProjectController struct {
	markup.RestController `class:"rest-controller"`

	ProjectService service.ProjectService `inject:"#ProjectService"`
	Responder      glass.MainResponder    `inject:"#glass-main-responder"`
}

func (inst *ProjectController) _Impl() glass.Controller {
	return inst
}

// Init 初始化
func (inst *ProjectController) Init(ec glass.EngineConnection) error {

	ec = ec.RequestMapping("projects")

	ec.Handle(http.MethodGet, "", inst.handleGetList)
	ec.Handle(http.MethodPost, "", inst.handlePost)

	ec.Handle(http.MethodGet, ":id", inst.handleGetOne)
	ec.Handle(http.MethodPut, ":id", inst.handlePut)
	ec.Handle(http.MethodDelete, ":id", inst.handleDelete)

	return nil
}

func (inst *ProjectController) handleGetList(c *gin.Context) {
	req := &myProjectRequest{
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

func (inst *ProjectController) handleGetOne(c *gin.Context) {
	req := &myProjectRequest{
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

func (inst *ProjectController) handlePost(c *gin.Context) {
	req := &myProjectRequest{
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

func (inst *ProjectController) handlePut(c *gin.Context) {
	req := &myProjectRequest{
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

func (inst *ProjectController) handleDelete(c *gin.Context) {
	req := &myProjectRequest{
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

type myProjectRequest struct {
	gc         *gin.Context
	controller *ProjectController

	wantRequestID   bool
	wantRequestBody bool

	id    dxo.ProjectID
	body1 vo.Project
	body2 vo.Project
}

func (inst *myProjectRequest) open() error {

	c := inst.gc

	if inst.wantRequestID {
		id := c.Param("id")
		n, err := strconv.Atoi(id)
		if err != nil {
			return err
		}
		inst.id = dxo.ProjectID(n)
	}

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *myProjectRequest) send(err error) {
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

func (inst *myProjectRequest) doGetOne() error {
	id := inst.id
	ctx := inst.gc
	ser := inst.controller.ProjectService
	o, err := ser.Find(ctx, id)
	if err != nil {
		return err
	}
	inst.body2.Projects = []*dto.Project{o}
	return nil
}

func (inst *myProjectRequest) doGetList() error {
	ctx := inst.gc
	ser := inst.controller.ProjectService
	list, err := ser.ListAll(ctx)
	if err != nil {
		return err
	}
	inst.body2.Projects = list
	return nil
}

func (inst *myProjectRequest) doPost() error {

	return nil
}

func (inst *myProjectRequest) doPut() error {
	return nil
}

func (inst *myProjectRequest) doDelete() error {

	return nil
}
