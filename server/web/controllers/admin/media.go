package admin

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/bitwormhole/wpm/server/classes/media"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/web"
	"github.com/bitwormhole/wpm/server/web/dto"
	"github.com/bitwormhole/wpm/server/web/vo"
	"github.com/gin-gonic/gin"
	"github.com/starter-go/afs"
	"github.com/starter-go/base/lang"
	"github.com/starter-go/libgin"
)

// MediaController ...
type MediaController struct {

	//starter:component
	_as func(libgin.Controller) //starter:as(".")

	UploadSizeMax   int //starter:inject("${http.upload.content-length.max}")
	DownloadSizeMax int //starter:inject("${http.download.content-length.max}")

	Sender       libgin.Responder //starter:inject("#")
	MediaService media.Service    //starter:inject("#")
}

func (inst *MediaController) _impl() libgin.Controller {
	return inst
}

// Registration ...
func (inst *MediaController) Registration() *libgin.ControllerRegistration {
	return &libgin.ControllerRegistration{Route: inst.route}
}

func (inst *MediaController) route(rp libgin.RouterProxy) error {

	rp = rp.For("media")

	rp.POST("", inst.handlePost)
	rp.PUT(":id", inst.handle)
	rp.DELETE(":id", inst.handleDelete)

	rp.GET("", inst.handleGetList)
	rp.GET(":id", inst.handleGetOne)

	rp.GET("/media/:type1/:type2/:sum/:name", inst.handleGetMediaContent)

	return nil
}

func (inst *MediaController) handle(c *gin.Context) {
	req := &myMediaRequest{
		context:       c,
		controller:    inst,
		wantRequestID: false,
	}
	req.execute(req.doNOP)
}

func (inst *MediaController) handlePost(c *gin.Context) {
	req := &myMediaRequest{
		context:         c,
		controller:      inst,
		wantRequestID:   false,
		wantRequestBody: true,
	}
	req.execute(req.doInsert)
}

func (inst *MediaController) handleGetOne(c *gin.Context) {
	req := &myMediaRequest{
		context:       c,
		controller:    inst,
		wantRequestID: true,
	}
	req.execute(req.doGetOne)
}

func (inst *MediaController) handleGetList(c *gin.Context) {
	req := &myMediaRequest{
		context:          c,
		controller:       inst,
		wantRequestID:    false,
		wantRequestQuery: true,
	}
	req.execute(req.doGetList)
}

func (inst *MediaController) handleGetMediaContent(c *gin.Context) {
	req := &myMediaRequest{
		context:    c,
		controller: inst,
	}

	t1 := c.Param("type1")
	t2 := c.Param("type2")
	sum := c.Param("sum")

	m1 := new(dto.Media)
	m1.ContentType = t1 + "/" + t2
	m1.SHA256SUM = lang.Hex(sum)
	m1.Bucket = "" // as default

	m2, file, err := req.doGetMediaContent(m1)
	if err == nil && file != nil && m2 != nil {
		if file.IsFile() {
			err = req.sendFile(file, m2.ContentType)
			if err == nil {
				return
			}
		}
	}

	req.body2.Status = http.StatusNotFound
	req.send(err)
}

func (inst *MediaController) handleDelete(c *gin.Context) {
	req := &myMediaRequest{
		context:       c,
		controller:    inst,
		wantRequestID: true,
	}
	req.execute(req.doDelete)
}

////////////////////////////////////////////////////////////////////////////////

type myMediaRequest struct {
	context    *gin.Context
	controller *MediaController

	wantRequestID    bool
	wantRequestBody  bool
	wantRequestQuery bool

	id    dxo.MediaID
	query media.Query

	body1 vo.Media
	body2 vo.Media
}

func (inst *myMediaRequest) open() error {

	c := inst.context

	if inst.wantRequestID {
		idstr := c.Param("id")
		idnum, err := strconv.ParseInt(idstr, 10, 64)
		if err != nil {
			return err
		}
		inst.id = dxo.MediaID(idnum)
	}

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	if inst.wantRequestQuery {
		getter := web.NewQueryGetter(c).Optional()
		inst.query.All = getter.GetBool("all")
		inst.query.Pagination = getter.GetPagination()
		err := getter.Error()
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *myMediaRequest) send(err error) {
	data := &inst.body2
	code := inst.body2.Status
	resp := new(libgin.Response)
	resp.Context = inst.context
	resp.Error = err
	resp.Data = data
	resp.Status = code
	inst.controller.Sender.Send(resp)
}

func (inst *myMediaRequest) sendFile(file afs.Path, contentType string) error {

	max := int64(inst.controller.DownloadSizeMax)
	size := file.GetInfo().Length()
	if size > max {
		return fmt.Errorf("media size is too large, max:%d have:%d", max, size)
	}

	data, err := file.GetIO().ReadBinary(nil)
	if err != nil {
		return err
	}

	if contentType == "" {
		contentType = "application/octet-stream"
	}

	c := inst.context
	c.Header("Content-Type", contentType)
	c.Header("Content-Length", strconv.FormatInt(size, 10))
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write(data)
	return nil
}

func (inst *myMediaRequest) execute(fn func() error) {
	err := inst.open()
	if err == nil {
		err = fn()
	}
	inst.send(err)
}

func (inst *myMediaRequest) doNOP() error {
	return nil
}

func (inst *myMediaRequest) doInsert() error {
	ctx := inst.context
	ser := inst.controller.MediaService
	src := inst.body1.Mediae
	dst := inst.body2.Mediae
	for _, item1 := range src {
		item2, err := ser.Insert(ctx, item1)
		if err != nil {
			return err
		}
		dst = append(dst, item2)
	}
	inst.body2.Mediae = dst
	return nil
}

func (inst *myMediaRequest) doGetOne() error {
	id := inst.id
	ctx := inst.context
	ser := inst.controller.MediaService
	o1, err := ser.Find(ctx, id)
	if err != nil {
		return err
	}
	o2 := &dto.Media{
		ID: o1.ID,
	}
	inst.body2.Mediae = []*dto.Media{o2}
	return nil
}

func (inst *myMediaRequest) doGetList() error {
	ctx := inst.context
	ser := inst.controller.MediaService
	q := &inst.query
	list, err := ser.List(ctx, q)
	if err != nil {
		return err
	}
	inst.body2.Mediae = list
	return nil
}

func (inst *myMediaRequest) doGetMediaContent(want *dto.Media) (*dto.Media, afs.Path, error) {
	ctx := inst.context
	ser := inst.controller.MediaService
	return ser.FindMediaFile(ctx, want)
}

func (inst *myMediaRequest) doDelete() error {
	ctx := inst.context
	ser := inst.controller.MediaService
	id := inst.id
	return ser.Remove(ctx, id)
}
