package plugins

import (
	"net/http"
	"strconv"

	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/dto"
	"github.com/bitwormhole/wpm/server/web/vo"
	"github.com/gin-gonic/gin"
)

// SoftwarePackageController 软件包控制器
type SoftwarePackageController struct {
	markup.RestController `class:"rest-controller"`

	SoftwarePackageService service.SoftwarePackageService `inject:"#SoftwarePackageService"`
	Responder              glass.MainResponder            `inject:"#glass-main-responder"`
}

func (inst *SoftwarePackageController) _Impl() glass.Controller {
	return inst
}

// Init 初始化
func (inst *SoftwarePackageController) Init(ec glass.EngineConnection) error {

	ec = ec.RequestMapping("packages")

	ec.Handle(http.MethodGet, "", inst.handleGetList)
	ec.Handle(http.MethodGet, ":id", inst.handleGetOne)

	ec.Handle(http.MethodPost, "", inst.handlePostInsert)
	ec.Handle(http.MethodPost, "update", inst.handlePostUpdate)

	ec.Handle(http.MethodPut, ":id", inst.handlePut)

	ec.Handle(http.MethodDelete, ":id", inst.handleDelete)

	return nil
}

func (inst *SoftwarePackageController) handleGetList(c *gin.Context) {
	req := &mySoftwarePackageRequest{
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

func (inst *SoftwarePackageController) handleGetOne(c *gin.Context) {
	req := &mySoftwarePackageRequest{
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

func (inst *SoftwarePackageController) handlePostInsert(c *gin.Context) {
	req := &mySoftwarePackageRequest{
		gc:              c,
		controller:      inst,
		wantRequestID:   false,
		wantRequestBody: true,
	}
	err := req.open()
	if err == nil {
		err = req.doPostInsert()
	}
	req.send(err)
}

func (inst *SoftwarePackageController) handlePostUpdate(c *gin.Context) {
	req := &mySoftwarePackageRequest{
		gc:              c,
		controller:      inst,
		wantRequestID:   false,
		wantRequestBody: true,
	}
	err := req.open()
	if err == nil {
		err = req.doPostUpdate()
	}
	req.send(err)
}

func (inst *SoftwarePackageController) handlePut(c *gin.Context) {
	req := &mySoftwarePackageRequest{
		gc:              c,
		controller:      inst,
		wantRequestID:   true,
		wantRequestBody: true,
	}
	err := req.open()
	if err == nil {
		err = req.doPutUpdateItem()
	}
	req.send(err)
}

func (inst *SoftwarePackageController) handleDelete(c *gin.Context) {
	req := &mySoftwarePackageRequest{
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

type mySoftwarePackageRequest struct {
	gc         *gin.Context
	controller *SoftwarePackageController

	wantRequestID   bool
	wantRequestBody bool

	id    dxo.SoftwarePackageID
	body1 vo.SoftwarePackage
	body2 vo.SoftwarePackage
}

func (inst *mySoftwarePackageRequest) open() error {

	c := inst.gc

	if inst.wantRequestID {
		id := c.Param("id")
		n, err := strconv.Atoi(id)
		if err != nil {
			return err
		}
		inst.id = dxo.SoftwarePackageID(n)
	}

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *mySoftwarePackageRequest) send(err error) {
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

func (inst *mySoftwarePackageRequest) doGetOne() error {
	id := inst.id
	ctx := inst.gc
	ser := inst.controller.SoftwarePackageService
	o, err := ser.Find(ctx, id)
	if err != nil {
		return err
	}
	inst.body2.Packages = []*dto.SoftwarePackage{o}
	return nil
}

func (inst *mySoftwarePackageRequest) doGetList() error {
	ctx := inst.gc
	ser := inst.controller.SoftwarePackageService
	list, err := ser.ListAll(ctx)
	if err != nil {
		return err
	}
	inst.body2.Packages = list
	return nil
}

func (inst *mySoftwarePackageRequest) doPostUpdate() error {
	ctx := inst.gc
	ser := inst.controller.SoftwarePackageService
	return ser.UpdateList(ctx)
}

func (inst *mySoftwarePackageRequest) doPostInsert() error {
	ctx := inst.gc
	ser := inst.controller.SoftwarePackageService
	o1 := inst.body1.Packages[0]
	o2, err := ser.Insert(ctx, o1)
	if err != nil {
		return nil
	}
	inst.body2.Packages = append(inst.body2.Packages, o2)
	return nil
}

func (inst *mySoftwarePackageRequest) doPutUpdateItem() error {
	ctx := inst.gc
	ser := inst.controller.SoftwarePackageService
	o1 := inst.body1.Packages[0]
	id := inst.id
	o2, err := ser.UpdateItem(ctx, id, o1)
	if err != nil {
		return nil
	}
	inst.body2.Packages = append(inst.body2.Packages, o2)
	return nil
}

func (inst *mySoftwarePackageRequest) doDelete() error {
	ctx := inst.gc
	ser := inst.controller.SoftwarePackageService
	return ser.Remove(ctx, inst.id)
}
