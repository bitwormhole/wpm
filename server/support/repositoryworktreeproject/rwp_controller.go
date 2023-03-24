package repositoryworktreeproject

import (
	"net/http"

	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/vo"
	"github.com/gin-gonic/gin"
)

// RepoWorktreeProjectController RepoWorktreeProject 控制器
type RepoWorktreeProjectController struct {
	markup.RestController `class:"rest-controller"`

	RWPService service.RepositoryWorktreeProjectService `inject:"#RepositoryWorktreeProjectService"`
	Responder  glass.MainResponder                      `inject:"#glass-main-responder"`
}

func (inst *RepoWorktreeProjectController) _Impl() glass.Controller {
	return inst
}

// Init 初始化
func (inst *RepoWorktreeProjectController) Init(ec glass.EngineConnection) error {

	ec = ec.RequestMapping("repository-worktree-project")

	ec.Handle(http.MethodPost, "find", inst.handlePostFind)
	ec.Handle(http.MethodPost, "save", inst.handlePostSave)

	// ec.Handle(http.MethodGet, "", inst.handleGetList)
	// ec.Handle(http.MethodGet, ":id", inst.handleGetOne)
	// ec.Handle(http.MethodPut, ":id", inst.handlePut)
	// ec.Handle(http.MethodDelete, ":id", inst.handleDelete)

	return nil
}

func (inst *RepoWorktreeProjectController) handlePostFind(c *gin.Context) {
	req := &myRepoWorktreeProjectRequest{
		gc:              c,
		controller:      inst,
		wantRequestID:   false,
		wantRequestBody: true,
	}
	err := req.open()
	if err == nil {
		err = req.doPostFind()
	}
	req.send(err)
}

func (inst *RepoWorktreeProjectController) handlePostSave(c *gin.Context) {
	req := &myRepoWorktreeProjectRequest{
		gc:              c,
		controller:      inst,
		wantRequestID:   false,
		wantRequestBody: true,
	}
	err := req.open()
	if err == nil {
		err = req.doPostSave()
	}
	req.send(err)
}

////////////////////////////////////////////////////////////////////////////////

type myRepoWorktreeProjectRequest struct {
	gc         *gin.Context
	controller *RepoWorktreeProjectController

	wantRequestID   bool
	wantRequestBody bool

	// id    dxo.RepoWorktreeProjectID
	body1 vo.RepositoryWorktreeProject
	body2 vo.RepositoryWorktreeProject
}

func (inst *myRepoWorktreeProjectRequest) open() error {

	c := inst.gc

	if inst.wantRequestID {
		// id := c.Param("id")
		// n, err := strconv.Atoi(id)
		// if err != nil {
		// 	return err
		// }
		// inst.id = dxo.RepoWorktreeProjectID(n)
	}

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *myRepoWorktreeProjectRequest) send(err error) {
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

func (inst *myRepoWorktreeProjectRequest) doPostFind() error {
	ctx := inst.gc
	ser := inst.controller.RWPService
	o1 := &inst.body1
	o2, err := ser.Find(ctx, o1)
	if err != nil {
		return err
	}
	inst.body2 = *o2
	return nil
}

func (inst *myRepoWorktreeProjectRequest) doPostSave() error {
	ctx := inst.gc
	ser := inst.controller.RWPService
	o1 := &inst.body1
	o2, err := ser.Save(ctx, o1)
	if err != nil {
		return err
	}
	inst.body2 = *o2
	return nil
}
