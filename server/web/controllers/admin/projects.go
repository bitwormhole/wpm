package admin

import (
	"strconv"

	"github.com/bitwormhole/wpm/server/classes/projects"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/web/dto"
	"github.com/bitwormhole/wpm/server/web/vo"
	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
)

// ProjectController ...
type ProjectController struct {

	//starter:component
	_as func(libgin.Controller) //starter:as(".")

	Sender  libgin.Responder //starter:inject("#")
	Service projects.Service //starter:inject("#")
}

func (inst *ProjectController) _impl() libgin.Controller {
	return inst
}

// Registration ...
func (inst *ProjectController) Registration() *libgin.ControllerRegistration {
	return &libgin.ControllerRegistration{Route: inst.route}
}

func (inst *ProjectController) route(rp libgin.RouterProxy) error {

	rp = rp.For("projects")

	rp.POST("", inst.handle)
	rp.PUT(":id", inst.handle)
	rp.DELETE(":id", inst.handle)

	rp.GET("", inst.handle)
	rp.GET(":id", inst.handleGetOne)

	return nil
}

func (inst *ProjectController) handle(c *gin.Context) {
	req := &myProjectRequest{
		context:       c,
		controller:    inst,
		wantRequestID: false,
	}
	req.execute(req.doNOP)
}

func (inst *ProjectController) handleGetOne(c *gin.Context) {
	req := &myProjectRequest{
		context:       c,
		controller:    inst,
		wantRequestID: true,
	}
	req.execute(req.doGetOne)
}

////////////////////////////////////////////////////////////////////////////////

type myProjectRequest struct {
	context    *gin.Context
	controller *ProjectController

	wantRequestID   bool
	wantRequestBody bool

	id    dxo.ProjectID
	body1 vo.Project
	body2 vo.Project
}

func (inst *myProjectRequest) open() error {

	c := inst.context

	if inst.wantRequestID {
		idstr := c.Param("id")
		idnum, err := strconv.ParseInt(idstr, 10, 64)
		if err != nil {
			return err
		}
		inst.id = dxo.ProjectID(idnum)
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
	data := &inst.body2
	code := inst.body2.Status
	resp := new(libgin.Response)
	resp.Context = inst.context
	resp.Error = err
	resp.Data = data
	resp.Status = code
	inst.controller.Sender.Send(resp)
}

func (inst *myProjectRequest) execute(fn func() error) {
	err := inst.open()
	if err == nil {
		err = fn()
	}
	inst.send(err)
}

func (inst *myProjectRequest) doNOP() error {
	return nil
}

func (inst *myProjectRequest) doGetOne() error {
	id := inst.id
	o1, err := inst.controller.Service.Find(nil, id)
	if err != nil {
		return err
	}
	o2 := &dto.Project{
		ID: o1.ID,
	}
	inst.body2.Projects = []*dto.Project{o2}
	return nil
}
