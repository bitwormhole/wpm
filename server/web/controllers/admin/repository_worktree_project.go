package admin

import (
	"github.com/bitwormhole/wpm/common/objects/vo"
	"github.com/bitwormhole/wpm/server/classes/repositoryworktreeproject"
	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
)

// RepositoryWorktreeProjectController ...
type RepositoryWorktreeProjectController struct {

	//starter:component
	_as func(libgin.Controller) //starter:as(".")

	Sender  libgin.Responder                  //starter:inject("#")
	Service repositoryworktreeproject.Service //starter:inject("#")

	// Dao     examples.DAO                      //starter:inject("#")

}

func (inst *RepositoryWorktreeProjectController) _impl() libgin.Controller {
	return inst
}

// Registration ...
func (inst *RepositoryWorktreeProjectController) Registration() *libgin.ControllerRegistration {
	return &libgin.ControllerRegistration{Route: inst.route}
}

func (inst *RepositoryWorktreeProjectController) route(rp libgin.RouterProxy) error {

	rp = rp.For("repository-worktree-project")

	// rp.POST("", inst.handle)
	// rp.PUT(":id", inst.handle)
	// rp.DELETE(":id", inst.handle)

	// rp.GET("", inst.handle)
	// rp.GET(":id", inst.handleGetOne)

	rp.POST("find", inst.handlePostFind)

	return nil
}

func (inst *RepositoryWorktreeProjectController) handle(c *gin.Context) {
	req := &myRepositoryWorktreeProjectRequest{
		context:       c,
		controller:    inst,
		wantRequestID: false,
	}
	req.execute(req.doNOP)
}

func (inst *RepositoryWorktreeProjectController) handlePostFind(c *gin.Context) {
	req := &myRepositoryWorktreeProjectRequest{
		context:         c,
		controller:      inst,
		wantRequestID:   false,
		wantRequestBody: true,
	}
	req.execute(req.doFind)
}

////////////////////////////////////////////////////////////////////////////////

type myRepositoryWorktreeProjectRequest struct {
	context    *gin.Context
	controller *RepositoryWorktreeProjectController

	wantRequestID   bool
	wantRequestBody bool

	// id    dxo.ExampleID
	body1 vo.RepositoryWorktreeProject
	body2 vo.RepositoryWorktreeProject
}

func (inst *myRepositoryWorktreeProjectRequest) open() error {

	c := inst.context

	if inst.wantRequestID {
		// idstr := c.Param("id")
		// idnum, err := strconv.ParseInt(idstr, 10, 64)
		// if err != nil {
		// 	return err
		// }
		// inst.id = dxo.ExampleID(idnum)
	}

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *myRepositoryWorktreeProjectRequest) send(err error) {
	data := &inst.body2
	code := inst.body2.Status
	resp := new(libgin.Response)
	resp.Context = inst.context
	resp.Error = err
	resp.Data = data
	resp.Status = code
	inst.controller.Sender.Send(resp)
}

func (inst *myRepositoryWorktreeProjectRequest) execute(fn func() error) {
	err := inst.open()
	if err == nil {
		err = fn()
	}
	inst.send(err)
}

func (inst *myRepositoryWorktreeProjectRequest) doNOP() error {
	return nil
}

func (inst *myRepositoryWorktreeProjectRequest) doFind() error {
	ctx := inst.context
	ser := inst.controller.Service
	b1 := &inst.body1
	b2 := &inst.body2
	return ser.Find(ctx, b1, b2)
}
