package admin

import (
	"github.com/bitwormhole/wpm/server/classes/repositories"
	"github.com/bitwormhole/wpm/server/web/vo"
	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
)

// ImportRepositoryController ...
type ImportRepositoryController struct {

	//starter:component
	_as func(libgin.Controller) //starter:as(".")

	Sender  libgin.Responder                     //starter:inject("#")
	Service repositories.ImportRepositoryService //starter:inject("#")
}

func (inst *ImportRepositoryController) _impl() libgin.Controller {
	return inst
}

// Registration ...
func (inst *ImportRepositoryController) Registration() *libgin.ControllerRegistration {
	return &libgin.ControllerRegistration{Route: inst.route}
}

func (inst *ImportRepositoryController) route(rp libgin.RouterProxy) error {

	rp = rp.For("import-repositories")
	rp.POST("", inst.handlePost)
	return nil
}

func (inst *ImportRepositoryController) handle(c *gin.Context) {
	req := &myImportRepositoryRequest{
		context:       c,
		controller:    inst,
		wantRequestID: false,
	}
	req.execute(req.doNOP)
}

func (inst *ImportRepositoryController) handlePost(c *gin.Context) {
	req := &myImportRepositoryRequest{
		context:         c,
		controller:      inst,
		wantRequestID:   false,
		wantRequestBody: true,
	}
	req.execute(req.doPost)
}

////////////////////////////////////////////////////////////////////////////////

type myImportRepositoryRequest struct {
	context    *gin.Context
	controller *ImportRepositoryController

	wantRequestID   bool
	wantRequestBody bool

	// id    dxo.ImportRepositoryID
	body1 vo.RepositoryImport
	body2 vo.RepositoryImport
}

func (inst *myImportRepositoryRequest) open() error {

	c := inst.context

	if inst.wantRequestID {
		// idstr := c.Param("id")
		// idnum, err := strconv.ParseInt(idstr, 10, 64)
		// if err != nil {
		// 	return err
		// }
		// inst.id = dxo.ImportRepositoryID(idnum)
	}

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *myImportRepositoryRequest) send(err error) {
	data := &inst.body2
	code := inst.body2.Status
	resp := new(libgin.Response)
	resp.Context = inst.context
	resp.Error = err
	resp.Data = data
	resp.Status = code
	inst.controller.Sender.Send(resp)
}

func (inst *myImportRepositoryRequest) execute(fn func() error) {
	err := inst.open()
	if err == nil {
		err = fn()
	}
	inst.send(err)
}

func (inst *myImportRepositoryRequest) doNOP() error {
	return nil
}

func (inst *myImportRepositoryRequest) doPost() error {
	ctx := inst.context
	ser := inst.controller.Service
	want := &inst.body1
	have := &inst.body2
	return ser.Handle(ctx, want, have)
}
