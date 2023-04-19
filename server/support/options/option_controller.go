package options

import (
	"net/http"

	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/vo"
	"github.com/gin-gonic/gin"
)

// TheOptionController 控制器
type TheOptionController struct {
	markup.RestController `class:"rest-controller"`

	Responder     glass.MainResponder   `inject:"#glass-main-responder"`
	OptionService service.OptionService `inject:"#OptionService"`
}

func (inst *TheOptionController) _Impl() glass.Controller {
	return inst
}

// Init 初始化
func (inst *TheOptionController) Init(ec glass.EngineConnection) error {
	ec = ec.RequestMapping("options")
	ec.Handle(http.MethodGet, "", inst.handleGetAll)
	return nil
}

func (inst *TheOptionController) handleGetAll(c *gin.Context) {
	req := &myOptionRequest{
		gc:         c,
		controller: inst,
	}
	err := req.open()
	if err == nil {
		err = req.doGetAll()
	}
	req.send(err)
}

////////////////////////////////////////////////////////////////////////////////

type myOptionRequest struct {
	gc         *gin.Context
	controller *TheOptionController

	body1 vo.Option
	body2 vo.Option
}

func (inst *myOptionRequest) open() error {

	// c := inst.gc
	// if inst.wantRequestBody {
	// 	err := c.BindJSON(&inst.body1)
	// 	if err != nil {
	// 		return err
	// 	}
	// }

	return nil
}

func (inst *myOptionRequest) send(err error) {
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

func (inst *myOptionRequest) doGetAll() error {
	o := &inst.body2
	return inst.controller.OptionService.GetOptions(o)
}
