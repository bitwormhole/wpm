package controller

import (
	"errors"
	"net/http"

	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/app/service"
	"github.com/bitwormhole/wpm/app/web/vo"
	"github.com/gin-gonic/gin"
)

// RepositoryImportController 仓库控制器
type RepositoryImportController struct {
	markup.RestController `class:"rest-controller"`

	ImportService service.RepositoryImportService `inject:"#RepositoryImportService"`
	Responder     glass.MainResponder             `inject:"#glass-main-responder"`
}

func (inst *RepositoryImportController) _Impl() glass.Controller {
	return inst
}

// Init 初始化
func (inst *RepositoryImportController) Init(ec glass.EngineConnection) error {
	ec = ec.RequestMapping("import-repositories")
	ec.Handle(http.MethodPost, "", inst.handlePost)
	ec.Handle(http.MethodPut, "", inst.handlePost)
	return nil
}

func (inst *RepositoryImportController) handlePost(c *gin.Context) {
	req := &myRepositoryImportRequest{
		gc:               c,
		controller:       inst,
		wantRequestID:    false,
		wantRequestBody:  true,
		wantRequestQuery: true,
	}
	err := req.open()
	if err == nil {
		err = req.doPost()
	}
	req.send(err)
}

////////////////////////////////////////////////////////////////////////////////

type myRepositoryImportRequest struct {
	gc         *gin.Context
	controller *RepositoryImportController

	wantRequestID    bool
	wantRequestBody  bool
	wantRequestQuery bool

	// id    dxo.RepositoryID
	body1 vo.RepositoryImport
	body2 vo.RepositoryImport
}

func (inst *myRepositoryImportRequest) open() error {

	c := inst.gc

	if inst.wantRequestID {
		// id := c.Param("id")
		// n, err := strconv.Atoi(id)
		// if err != nil {
		// 	return err
		// }
		// inst.id = dxo.RepositoryID(n)
	}

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	if inst.wantRequestQuery {
		action := c.Query("action")
		if action != "" {
			inst.body1.Action = vo.RepositoryImportAction(action)
		}
	}

	return nil
}

func (inst *myRepositoryImportRequest) send(err error) {
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

func (inst *myRepositoryImportRequest) doPost() error {
	ctx := inst.gc
	ser := inst.controller.ImportService
	action := inst.body1.Action

	var err error
	var result *vo.RepositoryImport

	switch action {
	case vo.RepositoryImportActionFind:
		result, err = ser.Find(ctx, &inst.body1)
		break
	case vo.RepositoryImportActionLocate:
		result, err = ser.Locate(ctx, &inst.body1)
		break
	case vo.RepositoryImportActionSave:
		result, err = ser.Save(ctx, &inst.body1)
		break
	default:
		// result, err := ser.FindOrLocate(ctx, &inst.body1)
		err = errors.New("bad action:" + action.String())
		break
	}

	if err != nil {
		return err
	}

	inst.body2 = *result
	return nil
}
