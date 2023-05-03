package filequery

import (
	"net/http"

	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/vo"
	"github.com/gin-gonic/gin"
)

// FileQueryController FileQuery 控制器
type FileQueryController struct {
	markup.RestController `class:"rest-controller"`

	FileQueryService service.FileQueryService `inject:"#FileQueryService"`
	Responder        glass.MainResponder      `inject:"#glass-main-responder"`
}

func (inst *FileQueryController) _Impl() glass.Controller {
	return inst
}

// Init 初始化
func (inst *FileQueryController) Init(ec glass.EngineConnection) error {

	ec = ec.RequestMapping("file-query")
	ec.Handle(http.MethodPost, "", inst.handlePost)
	return nil
}

func (inst *FileQueryController) handlePost(c *gin.Context) {
	req := &myFileQueryRequest{
		gc:                 c,
		controller:         inst,
		wantRequestID:      false,
		wantRequestBody:    true,
		wantRequestOptions: true,
	}
	err := req.open()
	if err == nil {
		err = req.doPost()
	}
	req.send(err)
}

////////////////////////////////////////////////////////////////////////////////

type myFileQueryRequest struct {
	gc         *gin.Context
	controller *FileQueryController

	wantRequestID      bool
	wantRequestBody    bool
	wantRequestOptions bool

	// id    dxo.FileQueryID
	options service.FileQueryOptions
	body1   vo.FileQuery
	body2   vo.FileQuery
}

func (inst *myFileQueryRequest) open() error {

	c := inst.gc

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	if inst.wantRequestOptions {
		body := &inst.body1
		ops := &inst.options
		ops.WithContentType = body.OptionWithContentType
	}

	return nil
}

func (inst *myFileQueryRequest) send(err error) {
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

func (inst *myFileQueryRequest) doPost() error {
	ctx := inst.gc
	ser := inst.controller.FileQueryService
	ops := &inst.options
	q := &inst.body1
	res, err := ser.Query(ctx, q, ops)
	if err != nil {
		return err
	}
	inst.body2 = *res
	return nil
}
