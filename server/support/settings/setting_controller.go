package settings

import (
	"net/http"
	"strconv"

	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/vo"
	"github.com/gin-gonic/gin"
)

// SettingController Settings 控制器
type SettingController struct {
	markup.RestController `class:"rest-controller"`

	SettingService service.SettingService `inject:"#SettingService"`
	Responder      glass.MainResponder    `inject:"#glass-main-responder"`
}

func (inst *SettingController) _Impl() glass.Controller {
	return inst
}

// Init 初始化
func (inst *SettingController) Init(ec glass.EngineConnection) error {
	ec = ec.RequestMapping("settings")
	ec.Handle(http.MethodGet, "", inst.handleGetList)
	return nil
}

func (inst *SettingController) handleGetList(c *gin.Context) {
	req := &mySettingRequest{
		gc:                 c,
		controller:         inst,
		wantRequestID:      false,
		wantRequestIDs:     false,
		wantRequestBody:    false,
		wantRequestOptions: false,
	}
	err := req.open()
	if err == nil {
		err = req.doGetList()
	}
	req.send(err)
}

////////////////////////////////////////////////////////////////////////////////

type mySettingRequest struct {
	gc         *gin.Context
	controller *SettingController

	wantRequestID      bool
	wantRequestIDs     bool
	wantRequestBody    bool
	wantRequestOptions bool

	id  dxo.SettingID
	ids []dxo.SettingID

	body1 vo.Settings
	body2 vo.Settings
}

func (inst *mySettingRequest) open() error {

	c := inst.gc

	if inst.wantRequestID {
		id := c.Param("id")
		n, err := strconv.Atoi(id)
		if err != nil {
			return err
		}
		inst.id = dxo.SettingID(n)
	}

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *mySettingRequest) send(err error) {
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

func (inst *mySettingRequest) doGetList() error {
	ser := inst.controller.SettingService

	all := ser.ListAll()
	inst.body2.Table = all

	o, err := ser.GetSettings()
	if err == nil {
		inst.body2.Settings = o
	}

	return nil
}
