package projects

import (
	"net/http"

	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/vo"
	"github.com/gin-gonic/gin"
)

// ProjectImportController 仓库控制器
type ProjectImportController struct {
	markup.RestController `class:"rest-controller"`

	ProjectImportService service.ProjectImportService `inject:"#ProjectImportService"`
	Responder            glass.MainResponder          `inject:"#glass-main-responder"`
}

func (inst *ProjectImportController) _Impl() glass.Controller {
	return inst
}

// Init 初始化
func (inst *ProjectImportController) Init(ec glass.EngineConnection) error {

	ec = ec.RequestMapping("import-projects")

	ec.Handle(http.MethodPost, "", inst.handlePost)
	ec.Handle(http.MethodPut, "", inst.handlePut)

	return nil
}

func (inst *ProjectImportController) handlePost(c *gin.Context) {
	req := &myProjectImportRequest{
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

func (inst *ProjectImportController) handlePut(c *gin.Context) {
	req := &myProjectImportRequest{
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

////////////////////////////////////////////////////////////////////////////////

type myProjectImportRequest struct {
	gc         *gin.Context
	controller *ProjectImportController

	wantRequestID   bool
	wantRequestBody bool

	// id    dxo.ProjectID
	body1 vo.ProjectImport
	body2 vo.ProjectImport
}

func (inst *myProjectImportRequest) open() error {

	c := inst.gc

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *myProjectImportRequest) send(err error) {
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

func (inst *myProjectImportRequest) doPost() error {
	ctx := inst.gc
	ser := inst.controller.ProjectImportService
	result, err := ser.FindOrLocate(ctx, &inst.body1)
	if err != nil {
		return err
	}
	inst.body2 = *result
	return nil
}

func (inst *myProjectImportRequest) doPut() error {
	ctx := inst.gc
	ser := inst.controller.ProjectImportService
	result, err := ser.Save(ctx, &inst.body1)
	if err != nil {
		return err
	}
	inst.body2 = *result
	return nil
}
