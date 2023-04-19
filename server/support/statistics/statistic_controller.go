package statistics

import (
	"net/http"

	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/data/dao"
	"github.com/bitwormhole/wpm/server/web/vo"
	"github.com/gin-gonic/gin"
)

// TheStatisticController Media 控制器
type TheStatisticController struct {
	markup.RestController `class:"rest-controller"`

	Responder glass.MainResponder `inject:"#glass-main-responder"`

	StatisticDAO dao.StatisticDAO `inject:"#StatisticDAO"`
}

func (inst *TheStatisticController) _Impl() glass.Controller {
	return inst
}

// Init 初始化
func (inst *TheStatisticController) Init(ec glass.EngineConnection) error {
	ec = ec.RequestMapping("statistics")
	ec.Handle(http.MethodGet, "", inst.handleGet)
	return nil
}

func (inst *TheStatisticController) handleGet(c *gin.Context) {
	req := &myStatisticRequest{
		gc:         c,
		controller: inst,
		// wantRequestID:      false,
		// wantRequestIDs:     true,
		// wantRequestBody:    false,
		// wantRequestOptions: true,
	}
	err := req.open()
	if err == nil {
		err = req.doGet()
	}
	req.send(err)
}

////////////////////////////////////////////////////////////////////////////////

type myStatisticRequest struct {
	gc         *gin.Context
	controller *TheStatisticController

	// wantRequestID      bool
	// wantRequestIDs     bool
	// wantRequestBody    bool
	// wantRequestOptions bool

	// options service.MediaOptions
	// id      dxo.MediaID
	// ids     []dxo.MediaID

	body1 vo.Statistic
	body2 vo.Statistic
}

func (inst *myStatisticRequest) open() error {

	// c := inst.gc

	// if inst.wantRequestIDs {
	// 	idstr := c.Query("ids")
	// 	ids, err := inst.parseIds(idstr)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	inst.ids = ids
	// }

	return nil
}

func (inst *myStatisticRequest) send(err error) {
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

func (inst *myStatisticRequest) doGet() error {

	src := inst.controller.StatisticDAO
	dst := &inst.body2.Statistics

	dst.Backups = src.CountBackups()
	dst.ContentTypes = src.CountContentTypes()
	dst.Executables = src.CountExecutables()
	dst.IntentTemplates = src.CountIntentTemplates()
	dst.Mediae = src.CountMediae()
	dst.PlugIns = src.CountPlugIns()
	dst.Projects = src.CountProjects()
	dst.Repositories = src.CountRepositories()

	return nil
}
