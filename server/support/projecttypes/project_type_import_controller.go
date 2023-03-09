package projecttypes

import (
	"fmt"
	"net/http"

	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/vo"
	"github.com/gin-gonic/gin"
)

// ProjectTypeImportController ProjectTypeImport 控制器
type ProjectTypeImportController struct {
	markup.RestController `class:"rest-controller"`

	ProjectTypeImportService service.ProjectTypeImportService `inject:"#ProjectTypeImportService"`
	Responder                glass.MainResponder              `inject:"#glass-main-responder"`
}

func (inst *ProjectTypeImportController) _Impl() glass.Controller {
	return inst
}

// Init 初始化
func (inst *ProjectTypeImportController) Init(ec glass.EngineConnection) error {

	ec = ec.RequestMapping("import-project-types")

	ec.Handle(http.MethodPost, "", inst.handlePost)

	// ec.Handle(http.MethodGet, "", inst.handleGetList)
	// ec.Handle(http.MethodGet, ":id", inst.handleGetOne)
	// ec.Handle(http.MethodPut, ":id", inst.handlePut)
	// ec.Handle(http.MethodDelete, ":id", inst.handleDelete)

	return nil
}

func (inst *ProjectTypeImportController) handleGetList(c *gin.Context) {
	req := &myProjectTypeImportRequest{
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

func (inst *ProjectTypeImportController) handleGetOne(c *gin.Context) {
	req := &myProjectTypeImportRequest{
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

func (inst *ProjectTypeImportController) handlePost(c *gin.Context) {
	req := &myProjectTypeImportRequest{
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

func (inst *ProjectTypeImportController) handlePut(c *gin.Context) {
	req := &myProjectTypeImportRequest{
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

func (inst *ProjectTypeImportController) handleDelete(c *gin.Context) {
	req := &myProjectTypeImportRequest{
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

type myProjectTypeImportRequest struct {
	gc         *gin.Context
	controller *ProjectTypeImportController

	wantRequestID   bool
	wantRequestBody bool

	// id    dxo.ProjectTypeImportID
	body1 vo.ProjectTypeImport
	body2 vo.ProjectTypeImport
}

func (inst *myProjectTypeImportRequest) open() error {

	c := inst.gc

	if inst.wantRequestID {
		// id := c.Param("id")
		// n, err := strconv.Atoi(id)
		// if err != nil {
		// 	return err
		// }
		// inst.id = dxo.ProjectTypeImportID(n)
	}

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *myProjectTypeImportRequest) send(err error) {
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

func (inst *myProjectTypeImportRequest) doGetOne() error {

	// id := inst.id
	// ctx := inst.gc
	// ser := inst.controller.ProjectTypeImportService
	// o, err := ser.Find(ctx, id)
	// if err != nil {
	// 	return err
	// }
	// inst.body2.ProjectTypeImports = []*dto.ProjectTypeImport{o}
	// return nil

	return fmt.Errorf("no impl")
}

func (inst *myProjectTypeImportRequest) doGetList() error {

	// ctx := inst.gc
	// ser := inst.controller.ProjectTypeImportService
	// list, err := ser.ListAll(ctx)
	// if err != nil {
	// 	return err
	// }
	// inst.body2.ProjectTypeImports = list
	// return nil

	return fmt.Errorf("no impl")
}

func (inst *myProjectTypeImportRequest) doPost() error {
	ctx := inst.gc
	ser := inst.controller.ProjectTypeImportService
	return ser.ImportTypesFromPreset(ctx)
}

func (inst *myProjectTypeImportRequest) doPut() error {

	return fmt.Errorf("no impl")
}

func (inst *myProjectTypeImportRequest) doDelete() error {

	return fmt.Errorf("no impl")
}
