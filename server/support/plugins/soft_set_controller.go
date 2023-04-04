package plugins

import (
	"net/http"

	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/vo"
	"github.com/gin-gonic/gin"
)

// SoftwareSetController 软件包控制器
type SoftwareSetController struct {
	markup.RestController `class:"rest-controller"`

	SoftwareSetService service.SoftwareSetService `inject:"#SoftwareSetService"`
	Responder          glass.MainResponder        `inject:"#glass-main-responder"`
}

func (inst *SoftwareSetController) _Impl() glass.Controller {
	return inst
}

// Init 初始化
func (inst *SoftwareSetController) Init(ec glass.EngineConnection) error {

	ec = ec.RequestMapping("software-sets")

	ec.Handle(http.MethodGet, "", inst.handleGetList)

	ec.Handle(http.MethodPost, "install", inst.handlePostInstall)
	ec.Handle(http.MethodPost, "re-install", inst.handlePostReInstall)
	ec.Handle(http.MethodPost, "uninstall", inst.handlePostUninstall)
	ec.Handle(http.MethodPost, "upgrade", inst.handlePostUpgrade)

	return nil
}

func (inst *SoftwareSetController) handleGetList(c *gin.Context) {
	req := &mySoftwareSetRequest{
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

func (inst *SoftwareSetController) handlePostInstall(c *gin.Context) {
	req := &mySoftwareSetRequest{
		gc:              c,
		controller:      inst,
		wantRequestID:   false,
		wantRequestBody: true,
	}
	err := req.open()
	if err == nil {
		err = req.doPostInstall()
	}
	req.send(err)
}

func (inst *SoftwareSetController) handlePostUninstall(c *gin.Context) {
	req := &mySoftwareSetRequest{
		gc:              c,
		controller:      inst,
		wantRequestID:   false,
		wantRequestBody: true,
	}
	err := req.open()
	if err == nil {
		err = req.doPostUninstall()
	}
	req.send(err)
}

func (inst *SoftwareSetController) handlePostReInstall(c *gin.Context) {
	req := &mySoftwareSetRequest{
		gc:              c,
		controller:      inst,
		wantRequestID:   false,
		wantRequestBody: true,
	}
	err := req.open()
	if err == nil {
		err = req.doPostReInstall()
	}
	req.send(err)
}

func (inst *SoftwareSetController) handlePostUpgrade(c *gin.Context) {
	req := &mySoftwareSetRequest{
		gc:              c,
		controller:      inst,
		wantRequestID:   false,
		wantRequestBody: true,
	}
	err := req.open()
	if err == nil {
		err = req.doPostUpgrade()
	}
	req.send(err)
}

////////////////////////////////////////////////////////////////////////////////

type mySoftwareSetRequest struct {
	gc         *gin.Context
	controller *SoftwareSetController

	wantRequestID   bool
	wantRequestBody bool

	id    dxo.SoftwareSetID
	body1 vo.SoftwareSet
	body2 vo.SoftwareSet
}

func (inst *mySoftwareSetRequest) open() error {

	c := inst.gc

	if inst.wantRequestID {
		// id := c.Param("id")
		// n, err := strconv.Atoi(id)
		// if err != nil {
		// 	return err
		// }
		// inst.id = dxo.SoftwareSetID(n)
	}

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *mySoftwareSetRequest) send(err error) {
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

func (inst *mySoftwareSetRequest) doGetList() error {
	ctx := inst.gc
	ser := inst.controller.SoftwareSetService
	list, err := ser.ListAll(ctx)
	if err != nil {
		return err
	}
	inst.body2.Sets = list
	return nil
}

func (inst *mySoftwareSetRequest) doPostInstall() error {
	ctx := inst.gc
	ser := inst.controller.SoftwareSetService
	o := inst.body1.Sets[0]
	return ser.Install(ctx, o)
}

func (inst *mySoftwareSetRequest) doPostUpgrade() error {
	ctx := inst.gc
	ser := inst.controller.SoftwareSetService
	o := inst.body1.Sets[0]
	return ser.Upgrade(ctx, o)
}

func (inst *mySoftwareSetRequest) doPostReInstall() error {
	ctx := inst.gc
	ser := inst.controller.SoftwareSetService
	o := inst.body1.Sets[0]
	return ser.ReInstall(ctx, o)
}

func (inst *mySoftwareSetRequest) doPostUninstall() error {
	ctx := inst.gc
	ser := inst.controller.SoftwareSetService
	o := inst.body1.Sets[0]
	return ser.Uninstall(ctx, o)
}
