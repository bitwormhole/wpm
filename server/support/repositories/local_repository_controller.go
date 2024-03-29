package repositories

import (
	"net/http"
	"strconv"

	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/utils"
	"github.com/bitwormhole/wpm/server/web/dto"
	"github.com/bitwormhole/wpm/server/web/vo"
	"github.com/gin-gonic/gin"
)

// LocalRepositoryController 仓库控制器
type LocalRepositoryController struct {
	markup.RestController `class:"rest-controller"`

	RepoService service.LocalRepositoryService `inject:"#LocalRepositoryService"`
	Responder   glass.MainResponder            `inject:"#glass-main-responder"`
}

func (inst *LocalRepositoryController) _Impl() glass.Controller {
	return inst
}

// Init 初始化
func (inst *LocalRepositoryController) Init(ec glass.EngineConnection) error {

	ec = ec.RequestMapping("local-repositories")

	ec.Handle(http.MethodGet, "", inst.handleGetList)
	ec.Handle(http.MethodPost, "", inst.handlePost)

	ec.Handle(http.MethodGet, ":id", inst.handleGetOne)
	ec.Handle(http.MethodPut, ":id", inst.handlePut)
	ec.Handle(http.MethodDelete, ":id", inst.handleDelete)

	return nil
}

func (inst *LocalRepositoryController) handleGetList(c *gin.Context) {
	req := &myLocalRepositoryRequest{
		gc:                 c,
		controller:         inst,
		wantRequestID:      false,
		wantRequestBody:    false,
		wantRequestIds:     true,
		wantRequestOptions: true,
	}
	err := req.open()
	if err == nil {
		err = req.doGetList()
	}
	req.send(err)
}

func (inst *LocalRepositoryController) handleGetOne(c *gin.Context) {
	req := &myLocalRepositoryRequest{
		gc:                 c,
		controller:         inst,
		wantRequestID:      true,
		wantRequestBody:    false,
		wantRequestOptions: true,
	}
	err := req.open()
	if err == nil {
		err = req.doGetOne()
	}
	req.send(err)
}

func (inst *LocalRepositoryController) handlePost(c *gin.Context) {
	req := &myLocalRepositoryRequest{
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

func (inst *LocalRepositoryController) handlePut(c *gin.Context) {
	req := &myLocalRepositoryRequest{
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

func (inst *LocalRepositoryController) handleDelete(c *gin.Context) {
	req := &myLocalRepositoryRequest{
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

type myLocalRepositoryRequest struct {
	gc         *gin.Context
	controller *LocalRepositoryController

	wantRequestID      bool
	wantRequestIds     bool
	wantRequestBody    bool
	wantRequestOptions bool

	options service.LocalRepositoryOptions
	id      dxo.LocalRepositoryID
	ids     []dxo.LocalRepositoryID
	body1   vo.LocalRepository
	body2   vo.LocalRepository
}

func (inst *myLocalRepositoryRequest) open() error {

	c := inst.gc

	if inst.wantRequestID {
		id := c.Param("id")
		n, err := strconv.Atoi(id)
		if err != nil {
			return err
		}
		inst.id = dxo.LocalRepositoryID(n)
	}

	if inst.wantRequestIds {
		idstr := c.Query("ids")
		ids, err := inst.parseIds(idstr)
		if err != nil {
			return err
		}
		inst.ids = ids
	}

	if inst.wantRequestOptions {
		inst.options.All = inst.hasFlag(c, "all")
		inst.options.WithFileState = inst.hasFlag(c, "with-file-state")
		inst.options.WithGitStatus = inst.hasFlag(c, "with-git-status")
		inst.options.WithProjects = inst.hasFlag(c, "with-projects")
		inst.options.WithSubmodules = inst.hasFlag(c, "with-submodules")
		inst.options.WithWorktrees = inst.hasFlag(c, "with-worktrees")
	}

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *myLocalRepositoryRequest) hasFlag(c *gin.Context, name string) bool {
	return (&utils.GinUtils{}).HasFlag(c, name)
}

func (inst *myLocalRepositoryRequest) parseIds(str string) ([]dxo.LocalRepositoryID, error) {
	list := make([]dxo.LocalRepositoryID, 0)
	err := (&utils.GinUtils{}).ParseIDs(str, ".", func(n int) {
		list = append(list, dxo.LocalRepositoryID(n))
	})
	return list, err
}

func (inst *myLocalRepositoryRequest) send(err error) {
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

func (inst *myLocalRepositoryRequest) doGetOne() error {
	id := inst.id
	ctx := inst.gc
	ser := inst.controller.RepoService
	opt := &inst.options
	o, err := ser.Find(ctx, id, opt)
	if err != nil {
		return err
	}
	inst.body2.Repositories = []*dto.LocalRepository{o}
	return nil
}

func (inst *myLocalRepositoryRequest) doGetList() error {
	ctx := inst.gc
	ser := inst.controller.RepoService
	ids := inst.ids
	opt := &inst.options
	var err error
	var list []*dto.LocalRepository
	if inst.options.All {
		list, err = ser.ListAll(ctx, opt)
	} else {
		list, err = ser.ListByIds(ctx, ids, opt)
	}
	if err != nil {
		return err
	}
	inst.body2.Repositories = list
	return nil
}

func (inst *myLocalRepositoryRequest) doPost() error {
	return nil
}

func (inst *myLocalRepositoryRequest) doPut() error {
	ctx := inst.gc
	ser := inst.controller.RepoService
	id := inst.id
	o1 := inst.body1.Repositories[0]
	o2, err := ser.Update(ctx, id, o1)
	if err != nil {
		return err
	}
	inst.body2.Repositories = []*dto.LocalRepository{o2}
	return nil
}

func (inst *myLocalRepositoryRequest) doDelete() error {
	ctx := inst.gc
	ser := inst.controller.RepoService
	id := inst.id
	return ser.Remove(ctx, id)
}
