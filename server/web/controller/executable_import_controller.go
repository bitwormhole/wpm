package controller

import (
	"net/http"

	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/vo"
	"github.com/gin-gonic/gin"
)

// ExecutableImportController 仓库控制器
type ExecutableImportController struct {
	markup.RestController `class:"rest-controller"`

	ExecutableImportService service.ExecutableImportService `inject:"#ExecutableImportService"`
	Responder               glass.MainResponder             `inject:"#glass-main-responder"`
}

func (inst *ExecutableImportController) _Impl() glass.Controller {
	return inst
}

// Init 初始化
func (inst *ExecutableImportController) Init(ec glass.EngineConnection) error {

	ec = ec.RequestMapping("import-executables")

	ec.Handle(http.MethodPost, "", inst.handlePost)
	ec.Handle(http.MethodPut, "", inst.handlePut)

	return nil
}

func (inst *ExecutableImportController) handlePost(c *gin.Context) {
	req := &myExecutableImportRequest{
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

func (inst *ExecutableImportController) handlePut(c *gin.Context) {
	req := &myExecutableImportRequest{
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

type myExecutableImportRequest struct {
	gc         *gin.Context
	controller *ExecutableImportController

	wantRequestID   bool
	wantRequestBody bool

	// id    dxo.ExecutableID

	body1 vo.ExecutableImport
	body2 vo.ExecutableImport
}

func (inst *myExecutableImportRequest) open() error {

	c := inst.gc

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *myExecutableImportRequest) send(err error) {
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

func (inst *myExecutableImportRequest) doPost() error {
	ctx := inst.gc
	ser := inst.controller.ExecutableImportService
	result, err := ser.Locate(ctx, &inst.body1)
	if err != nil {
		return err
	}
	inst.body2 = *result
	return nil
}

func (inst *myExecutableImportRequest) doPut() error {
	ctx := inst.gc
	ser := inst.controller.ExecutableImportService
	result, err := ser.Save(ctx, &inst.body1)
	if err != nil {
		return err
	}
	inst.body2 = *result
	return nil
}
