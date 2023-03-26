package repositories

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

// RemoteRepositoryController 仓库控制器
type RemoteRepositoryController struct {
	markup.RestController `class:"rest-controller"`

	RepoService service.RemoteRepositoryService `inject:"#RemoteRepositoryService"`
	Responder   glass.MainResponder             `inject:"#glass-main-responder"`
}

func (inst *RemoteRepositoryController) _Impl() glass.Controller {
	return inst
}

// Init 初始化
func (inst *RemoteRepositoryController) Init(ec glass.EngineConnection) error {

	ec = ec.RequestMapping("remote-repositories")

	ec.Handle(http.MethodGet, "", inst.handleGetList)
	ec.Handle(http.MethodPost, "", inst.handlePost)

	ec.Handle(http.MethodGet, ":id", inst.handleGetOne)
	ec.Handle(http.MethodPut, ":id", inst.handlePut)
	ec.Handle(http.MethodDelete, ":id", inst.handleDelete)

	return nil
}

func (inst *RemoteRepositoryController) handleGetList(c *gin.Context) {
	req := &myRemoteRepositoryRequest{
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

func (inst *RemoteRepositoryController) handleGetOne(c *gin.Context) {
	req := &myRemoteRepositoryRequest{
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

func (inst *RemoteRepositoryController) handlePost(c *gin.Context) {
	req := &myRemoteRepositoryRequest{
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

func (inst *RemoteRepositoryController) handlePut(c *gin.Context) {
	req := &myRemoteRepositoryRequest{
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

func (inst *RemoteRepositoryController) handleDelete(c *gin.Context) {
	req := &myRemoteRepositoryRequest{
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

type myRemoteRepositoryRequest struct {
	gc         *gin.Context
	controller *RemoteRepositoryController

	wantRequestID   bool
	wantRequestBody bool

	id    dxo.RemoteRepositoryID
	body1 vo.RemoteRepository
	body2 vo.RemoteRepository
}

func (inst *myRemoteRepositoryRequest) open() error {

	c := inst.gc

	if inst.wantRequestID {
		id := c.Param("id")
		n, err := strconv.Atoi(id)
		if err != nil {
			return err
		}
		inst.id = dxo.RemoteRepositoryID(n)
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

func (inst *myRemoteRepositoryRequest) doGetOne() error {
	id := inst.id
	ctx := inst.gc
	ser := inst.controller.RepoService
	o, err := ser.Find(ctx, id)
	if err != nil {
		return err
	}
	inst.body2.Repositories = []*dto.RemoteRepository{o}
	return nil
}

func (inst *myRemoteRepositoryRequest) doGetList() error {
	ctx := inst.gc
	ser := inst.controller.RepoService
	list, err := ser.ListAll(ctx)
	if err != nil {
		return err
	}
	inst.body2.Repositories = list
	return nil
}

func (inst *myRemoteRepositoryRequest) doPost() error {

	return nil
}

func (inst *myRemoteRepositoryRequest) doPut() error {
	return nil
}

func (inst *myRemoteRepositoryRequest) doDelete() error {

	return nil
}
