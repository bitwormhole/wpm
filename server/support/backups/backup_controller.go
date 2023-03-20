package backups

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

// WpmBackupController ...
type WpmBackupController struct {
	markup.Component `id:"" class:"rest-controller"`

	BackupService service.DatabaseBackupService `inject:"#DatabaseBackupService"`
	Responder     glass.MainResponder           `inject:"#glass-main-responder"`
}

func (inst *WpmBackupController) _Impl() glass.Controller {
	return inst
}

// Init 初始化
func (inst *WpmBackupController) Init(ec glass.EngineConnection) error {
	ec = ec.RequestMapping("backups")

	ec.Handle(http.MethodGet, "", inst.handleGetList)

	ec.Handle(http.MethodPost, "export", inst.handlePostExport)
	ec.Handle(http.MethodPost, "import", inst.handlePostImport)

	return nil
}

func (inst *WpmBackupController) handlePostImport(c *gin.Context) {
	req := &myBackupRequest{
		WpmBackupController: inst,
		gc:                  c,
		wantRequestBody:     true,
	}
	err := req.open()
	if err == nil {
		err = req.doImport()
	}
	req.send(err)
}

func (inst *WpmBackupController) handlePostExport(c *gin.Context) {
	req := &myBackupRequest{
		WpmBackupController: inst,
		gc:                  c,
		wantRequestBody:     true,
	}
	err := req.open()
	if err == nil {
		err = req.doExport()
	}
	req.send(err)
}

func (inst *WpmBackupController) handleGetList(c *gin.Context) {
	req := &myBackupRequest{
		WpmBackupController: inst,
		gc:                  c,
		wantRequestBody:     false,
	}
	err := req.open()
	if err == nil {
		err = req.doListAll()
	}
	req.send(err)
}

////////////////////////////////////////////////////////////////////////////////

type myBackupRequest struct {
	gc                  *gin.Context
	WpmBackupController *WpmBackupController

	wantRequestID   bool
	wantRequestBody bool

	id    dxo.BackupID
	body1 vo.Backup
	body2 vo.Backup
}

func (inst *myBackupRequest) open() error {

	c := inst.gc

	if inst.wantRequestID {
		id := c.Param("id")
		n, err := strconv.Atoi(id)
		if err != nil {
			return err
		}
		inst.id = dxo.BackupID(n)
	}

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *myBackupRequest) send(err error) {
	ctx := inst.gc
	data := &inst.body2
	status := data.Status
	resp := &glass.Response{
		Context: ctx,
		Data:    data,
		Error:   err,
		Status:  status,
	}
	inst.WpmBackupController.Responder.Send(resp)
}

func (inst *myBackupRequest) doImport() error {
	ctx := inst.gc
	ser := inst.WpmBackupController.BackupService
	o1 := inst.body1.Backups[0]
	o2, err := ser.Import(ctx, o1)
	if err != nil {
		return err
	}
	inst.body2.Backups = []*dto.Backup{o2}
	return nil
}

func (inst *myBackupRequest) doExport() error {
	ctx := inst.gc
	ser := inst.WpmBackupController.BackupService
	o1 := inst.body1.Backups[0]
	o2, err := ser.Export(ctx, o1)
	if err != nil {
		return err
	}
	inst.body2.Backups = []*dto.Backup{o2}
	return nil
}

func (inst *myBackupRequest) doListAll() error {
	ctx := inst.gc
	ser := inst.WpmBackupController.BackupService
	list, err := ser.ListAll(ctx)
	if err != nil {
		return err
	}
	inst.body2.Backups = list
	return nil
}
