package admin

import (
	"fmt"
	"mime/multipart"

	"github.com/bitwormhole/wpm/server/classes/media"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/web/dto"
	"github.com/bitwormhole/wpm/server/web/vo"
	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
)

// UploadController ...
type UploadController struct {

	//starter:component
	_as func(libgin.Controller) //starter:as(".")

	UploadSizeMax   int //starter:inject("${http.upload.content-length.max}")
	DownloadSizeMax int //starter:inject("${http.download.content-length.max}")

	Sender       libgin.Responder //starter:inject("#")
	MediaService media.Service    //starter:inject("#")
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
		context:               c,
		controller:            inst,
		wantRequestUploadFile: true,
	}
	req.execute(req.doPostUpload)
}

////////////////////////////////////////////////////////////////////////////////

type myUploadRequest struct {
	context    *gin.Context
	controller *UploadController

	wantRequestID         bool
	wantRequestBody       bool
	wantRequestUploadFile bool

	id    dxo.MediaID
	body1 vo.Media
	body2 vo.Media
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

	if inst.wantRequestUploadFile {
		inst.body1.Mediae = nil
		err := inst.receiveMultipartFormData()
		if err != nil {
			return err
		}
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

func (inst *myUploadRequest) receiveMultipartFormData() error {
	c := inst.context
	mf, err := c.MultipartForm()
	if err != nil {
		return err
	}
	files := mf.File["file"]
	for _, file := range files {
		err = inst.receiveFile(file)
		if err != nil {
			return err
		}
	}
	return nil
}

func (inst *myUploadRequest) receiveFile(fh *multipart.FileHeader) error {

	file, err := fh.Open()
	if err != nil {
		return err
	}
	defer func() {
		file.Close()
	}()

	size := fh.Size
	limitSize := inst.controller.UploadSizeMax
	if size > int64(limitSize) {
		const f = "file size is too large to upload, limit:%d, have:%d"
		return fmt.Errorf(f, limitSize, size)
	}

	m1 := &dto.Media{}
	m1.Name = fh.Filename
	m1.FileSize = size
	m1.ContentType = fh.Header.Get("Content-Type")
	m1.Label = fh.Filename
	m1.Bucket = ""
	m1.URL = ""
	m1.Source = ""

	ctx := inst.context
	ser := inst.controller.MediaService
	m2, err := ser.ImportMediaStream(ctx, m1, file)
	if err != nil {
		return err
	}

	list := inst.body1.Mediae
	list = append(list, m2)
	inst.body1.Mediae = list
	return nil
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

func (inst *myUploadRequest) doPostUpload() error {
	list := inst.body1.Mediae
	inst.body2.Mediae = list
	return nil
}
