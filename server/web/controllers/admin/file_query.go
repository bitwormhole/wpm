package admin

import (
	"github.com/bitwormhole/wpm/common/objects/vo"
	"github.com/bitwormhole/wpm/server/classes/filequery"
	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
)

// FileQueryController ...
type FileQueryController struct {

	//starter:component
	_as func(libgin.Controller) //starter:as(".")

	Sender           libgin.Responder  //starter:inject("#")
	FileQueryService filequery.Service //starter:inject("#")
}

func (inst *FileQueryController) _impl() libgin.Controller {
	return inst
}

// Registration ...
func (inst *FileQueryController) Registration() *libgin.ControllerRegistration {
	return &libgin.ControllerRegistration{Route: inst.route}
}

func (inst *FileQueryController) route(rp libgin.RouterProxy) error {

	rp = rp.For("file-query")

	rp.POST("", inst.handlePostQuery)

	// rp.PUT(":id", inst.handle)
	// rp.DELETE(":id", inst.handle)
	// rp.GET("", inst.handle)
	// rp.GET(":id", inst.handleGetOne)

	return nil
}

func (inst *FileQueryController) handle(c *gin.Context) {
	req := &myFileQueryRequest{
		context:       c,
		controller:    inst,
		wantRequestID: false,
	}
	req.execute(req.doNOP)
}

func (inst *FileQueryController) handlePostQuery(c *gin.Context) {
	req := &myFileQueryRequest{
		context:         c,
		controller:      inst,
		wantRequestID:   false,
		wantRequestBody: true,
	}
	req.execute(req.doPostQuery)
}

////////////////////////////////////////////////////////////////////////////////

type myFileQueryRequest struct {
	context    *gin.Context
	controller *FileQueryController

	wantRequestID   bool
	wantRequestBody bool

	// id int
	body1 vo.FileQuery
	body2 vo.FileQuery
}

func (inst *myFileQueryRequest) open() error {

	c := inst.context

	if inst.wantRequestID {
	}

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *myFileQueryRequest) send(err error) {
	data := &inst.body2
	code := inst.body2.Status
	resp := new(libgin.Response)
	resp.Context = inst.context
	resp.Error = err
	resp.Data = data
	resp.Status = code
	inst.controller.Sender.Send(resp)
}

func (inst *myFileQueryRequest) execute(fn func() error) {
	err := inst.open()
	if err == nil {
		err = fn()
	}
	inst.send(err)
}

func (inst *myFileQueryRequest) doNOP() error {
	return nil
}

func (inst *myFileQueryRequest) doPostQuery() error {
	ctx := inst.context
	ser := inst.controller.FileQueryService
	want := &inst.body1
	result, err := ser.Query(ctx, want)
	if err != nil {
		return err
	}
	inst.body2 = *result
	return nil
}
