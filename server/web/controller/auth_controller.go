package controller

import (
	"net/http"

	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/web/vo"
	"github.com/gin-gonic/gin"
)

// AuthController 控制器
type AuthController struct {
	markup.RestController `class:"rest-controller"`

	// RepoService service.RepositoryService `inject:"#RepositoryService"`
	Responder glass.MainResponder `inject:"#glass-main-responder"`
}

func (inst *AuthController) _Impl() glass.Controller {
	return inst
}

// Init 初始化
func (inst *AuthController) Init(ec glass.EngineConnection) error {

	ec = ec.RequestMapping("auths")

	ec.Handle(http.MethodGet, "", inst.handleGetList)
	ec.Handle(http.MethodPost, "", inst.handlePost)

	ec.Handle(http.MethodGet, ":id", inst.handleGetOne)
	ec.Handle(http.MethodPut, ":id", inst.handlePut)
	ec.Handle(http.MethodDelete, ":id", inst.handleDelete)

	return nil
}

func (inst *AuthController) handleGetList(c *gin.Context) {
	req := &myAuthRequest{
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

func (inst *AuthController) handleGetOne(c *gin.Context) {
	req := &myAuthRequest{
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

func (inst *AuthController) handlePost(c *gin.Context) {
	req := &myAuthRequest{
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

func (inst *AuthController) handlePut(c *gin.Context) {
	req := &myAuthRequest{
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

func (inst *AuthController) handleDelete(c *gin.Context) {
	req := &myAuthRequest{
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

type myAuthRequest struct {
	gc         *gin.Context
	controller *AuthController

	wantRequestID   bool
	wantRequestBody bool

	id    string
	body1 vo.Auth
	body2 vo.Auth
}

func (inst *myAuthRequest) open() error {

	c := inst.gc

	if inst.wantRequestID {
		inst.id = c.Param("id")
	}

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *myAuthRequest) send(err error) {
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

func (inst *myAuthRequest) doGetOne() error {

	return nil
}

func (inst *myAuthRequest) doGetList() error {

	return nil
}

func (inst *myAuthRequest) doPost() error {

	return nil
}

func (inst *myAuthRequest) doPut() error {
	return nil
}

func (inst *myAuthRequest) doDelete() error {

	return nil
}
