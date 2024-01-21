package admin

import (
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/web/vo"
	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
)

// UploadController ...
type UploadController struct {

	//starter:component
	_as func(libgin.Controller) //starter:as(".")

	Sender libgin.Responder //starter:inject("#")
	// Dao    Uploads.DAO      //starter:inject("#")
}

func (inst *UploadController) _impl() libgin.Controller {
	return inst
}

// Registration ...
func (inst *UploadController) Registration() *libgin.ControllerRegistration {
	return &libgin.ControllerRegistration{Route: inst.route}
}

func (inst *UploadController) route(rp libgin.RouterProxy) error {

	rp.POST("/upload", inst.handlePostUpload)
	return nil
}

func (inst *UploadController) handle(c *gin.Context) {
	req := &myUploadRequest{
		context:       c,
		controller:    inst,
		wantRequestID: false,
	}
	req.execute(req.doNOP)
}

func (inst *UploadController) handlePostUpload(c *gin.Context) {
	req := &myUploadRequest{
		context:       c,
		controller:    inst,
		wantRequestID: true,
	}
	req.execute(req.doGetOne)
}

////////////////////////////////////////////////////////////////////////////////

type myUploadRequest struct {
	context    *gin.Context
	controller *UploadController

	wantRequestID   bool
	wantRequestBody bool

	id    dxo.ExampleID
	body1 vo.Example
	body2 vo.Example
}

func (inst *myUploadRequest) open() error {

	c := inst.context
	c.Query("")

	if inst.wantRequestID {
		// idstr := c.Param("id")
		// idnum, err := strconv.ParseInt(idstr, 10, 64)
		// if err != nil {
		// 	return err
		// }
		// inst.id = dxo.UploadID(idnum)
	}

	if inst.wantRequestBody {
		// err := c.BindJSON(&inst.body1)
		// if err != nil {
		// 	return err
		// }
	}

	return nil
}

func (inst *myUploadRequest) send(err error) {
	data := &inst.body2
	code := inst.body2.Status
	resp := new(libgin.Response)
	resp.Context = inst.context
	resp.Error = err
	resp.Data = data
	resp.Status = code
	inst.controller.Sender.Send(resp)
}

func (inst *myUploadRequest) execute(fn func() error) {
	err := inst.open()
	if err == nil {
		err = fn()
	}
	inst.send(err)
}

func (inst *myUploadRequest) doNOP() error {
	return nil
}

func (inst *myUploadRequest) doGetOne() error {
	return nil
}
