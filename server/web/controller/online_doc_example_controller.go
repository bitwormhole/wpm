package controller

import (
	"net/http"

	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/web/dto"
	"github.com/bitwormhole/wpm/server/web/vo"
	"github.com/gin-gonic/gin"
)

// OnlineDocumentExampleController 控制器
type OnlineDocumentExampleController struct {
	markup.RestController `class:"rest-controller"`

	Responder glass.MainResponder `inject:"#glass-main-responder"`
}

func (inst *OnlineDocumentExampleController) _Impl() glass.Controller {
	return inst
}

// Init 初始化
func (inst *OnlineDocumentExampleController) Init(ec glass.EngineConnection) error {
	ec = ec.RequestMapping("online-doc-example")
	ec.Handle(http.MethodGet, "", inst.handleGet)
	return nil
}

func (inst *OnlineDocumentExampleController) handleGet(c *gin.Context) {
	req := &myOnlineDocumentExampleRequest{
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

////////////////////////////////////////////////////////////////////////////////

type myOnlineDocumentExampleRequest struct {
	gc         *gin.Context
	controller *OnlineDocumentExampleController

	wantRequestID   bool
	wantRequestBody bool

	// id    dxo.OnlineDocumentExampleID
	body1 vo.Online
	body2 vo.Online
}

func (inst *myOnlineDocumentExampleRequest) open() error {

	c := inst.gc

	if inst.wantRequestID {
		// id := c.Param("id")
		// n, err := strconv.Atoi(id)
		// if err != nil {
		// 	return err
		// }
		// inst.id = dxo.OnlineDocumentExampleID(n)
	}

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *myOnlineDocumentExampleRequest) send(err error) {
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

func (inst *myOnlineDocumentExampleRequest) doGet() error {

	o2 := &inst.body2

	o2.Executables = append(o2.Executables, inst.makeExecutable())

	o2.IntentTemplates = append(o2.IntentTemplates, inst.makeIntentTemplate())

	o2.Mediae = append(o2.Mediae, inst.makeMedia())

	return nil
}

func (inst *myOnlineDocumentExampleRequest) makeExecutable() *dto.Executable {
	return &dto.Executable{}
}

func (inst *myOnlineDocumentExampleRequest) makeIntentTemplate() *dto.IntentTemplate {
	return &dto.IntentTemplate{}
}

func (inst *myOnlineDocumentExampleRequest) makeMedia() *dto.Media {
	return &dto.Media{}
}
