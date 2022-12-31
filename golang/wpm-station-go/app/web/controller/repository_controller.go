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

// RepositoryController 仓库控制器
type RepositoryController struct {
	markup.RestController `class:"rest-controller"`

	RepoService service.RepositoryService `inject:"#RepositoryService"`
	Responder   glass.MainResponder       `inject:"#glass-main-responder"`
}

func (inst *RepositoryController) _Impl() glass.Controller {
	return inst
}

// Init 初始化
func (inst *RepositoryController) Init(ec glass.EngineConnection) error {

	ec = ec.RequestMapping("repositories")

	ec.Handle(http.MethodGet, "", inst.handleGetList)
	ec.Handle(http.MethodPost, "", inst.handlePost)

	ec.Handle(http.MethodGet, ":id", inst.handleGetOne)
	ec.Handle(http.MethodPut, ":id", inst.handlePut)
	ec.Handle(http.MethodDelete, ":id", inst.handleDelete)

	return nil
}

func (inst *RepositoryController) handleGetList(c *gin.Context) {
	req := &myRepositoryRequest{
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

func (inst *RepositoryController) handleGetOne(c *gin.Context) {
	req := &myRepositoryRequest{
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

func (inst *RepositoryController) handlePost(c *gin.Context) {
	req := &myRepositoryRequest{
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

func (inst *RepositoryController) handlePut(c *gin.Context) {
	req := &myRepositoryRequest{
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

func (inst *RepositoryController) handleDelete(c *gin.Context) {
	req := &myRepositoryRequest{
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

type myRepositoryRequest struct {
	gc         *gin.Context
	controller *RepositoryController

	wantRequestID   bool
	wantRequestBody bool

	id    dxo.RepositoryID
	body1 vo.Repository
	body2 vo.Repository
}

func (inst *myRepositoryRequest) open() error {

	c := inst.gc

	if inst.wantRequestID {
		id := c.Param("id")
		n, err := strconv.Atoi(id)
		if err != nil {
			return err
		}
		inst.id = dxo.RepositoryID(n)
	}

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *myRepositoryRequest) send(err error) {
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

func (inst *myRepositoryRequest) doGetOne() error {
	id := inst.id
	ctx := inst.gc
	ser := inst.controller.RepoService
	o, err := ser.Find(ctx, id)
	if err != nil {
		return err
	}
	inst.body2.Repositories = []*dto.Repository{o}
	return nil
}

func (inst *myRepositoryRequest) doGetList() error {
	ctx := inst.gc
	ser := inst.controller.RepoService
	list, err := ser.ListAll(ctx)
	if err != nil {
		return err
	}
	inst.body2.Repositories = list
	return nil
}

func (inst *myRepositoryRequest) doPost() error {

	return nil
}

func (inst *myRepositoryRequest) doPut() error {
	return nil
}

func (inst *myRepositoryRequest) doDelete() error {

	return nil
}
