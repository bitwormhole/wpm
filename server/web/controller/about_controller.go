package controller

import (
	"fmt"
	"net/http"

	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/dto"
	"github.com/bitwormhole/wpm/server/web/vo"
	"github.com/gin-gonic/gin"
)

// AboutController About 控制器
type AboutController struct {
	markup.RestController `class:"rest-controller"`

	AboutService  service.AboutService       `inject:"#AboutService"`
	UpdateService service.CheckUpdateService `inject:"#CheckUpdateService"`
	Responder     glass.MainResponder        `inject:"#glass-main-responder"`
	Profile       string                     `inject:"${application.profiles.active}"`
}

func (inst *AboutController) _Impl() glass.Controller {
	return inst
}

// Init 初始化
func (inst *AboutController) Init(ec glass.EngineConnection) error {

	ec = ec.RequestMapping("about")

	ec.Handle(http.MethodGet, "", inst.handleGet)
	ec.Handle(http.MethodGet, "debug-mode", inst.handleGetDebug)
	ec.Handle(http.MethodGet, "latest-example", inst.handleGetLatestExample)

	ec.Handle(http.MethodPost, "ignore", inst.handlePostIgnorePackage)
	ec.Handle(http.MethodPost, "checkupdate", inst.handlePostCheckUpdate)

	return nil
}

func (inst *AboutController) handleGet(c *gin.Context) {
	req := &myAboutRequest{
		gc:              c,
		controller:      inst,
		wantRequestID:   false,
		wantRequestBody: false,
	}
	err := req.open()
	if err == nil {
		err = req.doGet()
	}
	req.send(err)
}

func (inst *AboutController) handleGetDebug(c *gin.Context) {
	req := &myAboutRequest{
		gc:              c,
		controller:      inst,
		wantRequestID:   false,
		wantRequestBody: false,
	}
	err := req.open()
	if err == nil {
		err = req.doGetDebug()
	}
	req.send(err)
}

func (inst *AboutController) handlePostCheckUpdate(c *gin.Context) {
	req := &myAboutRequest{
		gc:              c,
		controller:      inst,
		wantRequestID:   false,
		wantRequestBody: true,
	}
	err := req.open()
	if err == nil {
		err = req.doPostCheckUpdate()
	}
	req.send(err)
}

func (inst *AboutController) handlePostIgnorePackage(c *gin.Context) {
	req := &myAboutRequest{
		gc:              c,
		controller:      inst,
		wantRequestID:   false,
		wantRequestBody: true,
	}
	err := req.open()
	if err == nil {
		err = req.doPostIgnorePackage()
	}
	req.send(err)
}

func (inst *AboutController) handleGetLatestExample(c *gin.Context) {

	doc := &vo.SoftwarePackage{}
	p1 := &dto.SoftwarePackage{}
	p2 := &dto.SoftwarePackage{}

	doc.Packages = append(doc.Packages, p1)
	doc.Packages = append(doc.Packages, p2)

	c.JSON(http.StatusOK, doc)
}

////////////////////////////////////////////////////////////////////////////////

type myAboutRequest struct {
	gc         *gin.Context
	controller *AboutController

	wantRequestID   bool
	wantRequestBody bool

	// id    dxo.AboutID
	body1 vo.About
	body2 vo.About
}

func (inst *myAboutRequest) open() error {

	c := inst.gc

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *myAboutRequest) send(err error) {
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

func (inst *myAboutRequest) doGet() error {
	ctx := inst.gc
	ser := inst.controller.AboutService
	view, err := ser.GetInfo(ctx)
	if err != nil {
		return err
	}
	inst.body2 = *view
	return nil
}

func (inst *myAboutRequest) doGetDebug() error {
	ctl := inst.controller
	inst.body2.Profile = ctl.Profile
	return nil
}

func (inst *myAboutRequest) doPostCheckUpdate() error {
	ctx := inst.gc
	o := inst.body1.CheckUpdate
	if o == nil {
		o = &vo.AboutCheckUpdate{}
	}
	err := inst.controller.UpdateService.Check(ctx, o)
	if err != nil {
		return err
	}
	inst.body2.CheckUpdate = o
	return nil
}

func (inst *myAboutRequest) doPostIgnorePackage() error {
	ctx := inst.gc
	o := inst.body1.CheckUpdate
	if o == nil {
		return fmt.Errorf("bad request data")
	}
	err := inst.controller.UpdateService.Ignore(ctx, o)
	if err != nil {
		return err
	}
	inst.body2.CheckUpdate = o
	return nil
}
