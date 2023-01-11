package controller

import (
	"net/http"
	"strconv"

	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/vo"
	"github.com/gin-gonic/gin"
)

// MainRepositoryController 仓库控制器
type MainRepositoryController struct {
	markup.RestController `class:"rest-controller"`

	RepoService service.MainRepositoryService `inject:"#MainRepositoryService"`
	Responder   glass.MainResponder           `inject:"#glass-main-responder"`
}

func (inst *MainRepositoryController) _Impl() glass.Controller {
	return inst
}

// Init 初始化
func (inst *MainRepositoryController) Init(ec glass.EngineConnection) error {

	ec = ec.RequestMapping("main-repositories")

	ec.Handle(http.MethodGet, "", inst.handleGetList)
	ec.Handle(http.MethodPost, "", inst.handlePost)

	ec.Handle(http.MethodGet, ":id", inst.handleGetOne)
	ec.Handle(http.MethodPut, ":id", inst.handlePut)
	ec.Handle(http.MethodDelete, ":id", inst.handleDelete)

	return nil
}

func (inst *MainRepositoryController) handleGetList(c *gin.Context) {
	req := &myMainRepositoryRequest{
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

func (inst *MainRepositoryController) handleGetOne(c *gin.Context) {
	req := &myMainRepositoryRequest{
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

func (inst *MainRepositoryController) handlePost(c *gin.Context) {
	req := &myMainRepositoryRequest{
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

func (inst *MainRepositoryController) handlePut(c *gin.Context) {
	req := &myMainRepositoryRequest{
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

func (inst *MainRepositoryController) handleDelete(c *gin.Context) {
	req := &myMainRepositoryRequest{
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

type myMainRepositoryRequest struct {
	gc         *gin.Context
	controller *MainRepositoryController

	wantRequestID   bool
	wantRequestBody bool

	id    dxo.MainRepositoryID
	body1 vo.MainRepository
	body2 vo.MainRepository
}

func (inst *myMainRepositoryRequest) open() error {

	c := inst.gc

	if inst.wantRequestID {
		id := c.Param("id")
		n, err := strconv.Atoi(id)
		if err != nil {
			return err
		}
		inst.id = dxo.MainRepositoryID(n)
	}

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *myMainRepositoryRequest) send(err error) {
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

func (inst *myMainRepositoryRequest) doGetOne() error {
	id := inst.id
	ctx := inst.gc
	ser := inst.controller.RepoService
	o, err := ser.Find(ctx, id)
	if err != nil {
		return err
	}
	inst.body2.Repository = o
	return nil
}

func (inst *myMainRepositoryRequest) doGetList() error {
	ctx := inst.gc
	ser := inst.controller.RepoService
	list, err := ser.ListAll(ctx)
	if err != nil {
		return err
	}
	inst.body2.Repository = list[0]
	return nil
}

func (inst *myMainRepositoryRequest) doPost() error {

	return nil
}

func (inst *myMainRepositoryRequest) doPut() error {
	return nil
}

func (inst *myMainRepositoryRequest) doDelete() error {

	return nil
}
