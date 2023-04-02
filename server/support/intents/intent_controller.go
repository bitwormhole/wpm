package intents

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/utils"
	"github.com/bitwormhole/wpm/server/web/vo"
	"github.com/gin-gonic/gin"
)

// RunIntentController RunIntent 控制器
type RunIntentController struct {
	markup.RestController `class:"rest-controller"`

	IntentService service.IntentService `inject:"#IntentService"`
	Responder     glass.MainResponder   `inject:"#glass-main-responder"`
}

func (inst *RunIntentController) _Impl() glass.Controller {
	return inst
}

// Init 初始化
func (inst *RunIntentController) Init(ec glass.EngineConnection) error {

	ec = ec.RequestMapping("intents")

	ec.Handle(http.MethodPost, "", inst.handlePost)

	// ec.Handle(http.MethodGet, "", inst.handleGetList)
	// ec.Handle(http.MethodGet, ":id", inst.handleGetOne)
	// ec.Handle(http.MethodPut, ":id", inst.handlePut)
	// ec.Handle(http.MethodDelete, ":id", inst.handleDelete)

	return nil
}

func (inst *RunIntentController) handlePost(c *gin.Context) {
	req := &myRunIntentRequest{
		gc:              c,
		controller:      inst,
		wantRequestID:   false,
		wantRequestBody: true,
	}
	err := req.open()
	if err == nil {
		err = req.doPost()
	}
	req.send(err)
}

func (inst *RunIntentController) handlePut(c *gin.Context) {
	req := &myRunIntentRequest{
		gc:              c,
		controller:      inst,
		wantRequestID:   true,
		wantRequestBody: true,
	}
	err := req.open()
	if err == nil {
		err = req.doPut()
	}
	req.send(err)
}

func (inst *RunIntentController) handleDelete(c *gin.Context) {
	req := &myRunIntentRequest{
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

type myRunIntentRequest struct {
	gc         *gin.Context
	controller *RunIntentController

	wantRequestID   bool
	wantRequestBody bool

	id    dxo.IntentID
	body1 vo.Intent
	body2 vo.Intent
}

func (inst *myRunIntentRequest) open() error {

	c := inst.gc

	if inst.wantRequestID {
		id := c.Param("id")
		n, err := strconv.Atoi(id)
		if err != nil {
			return err
		}
		inst.id = dxo.IntentID(n)
	}

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *myRunIntentRequest) send(err error) {
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

func (inst *myRunIntentRequest) doPost() error {

	ctx := inst.gc
	ser := inst.controller.IntentService
	olist1 := inst.body1.Intents
	errlist := utils.ErrorList{}
	errCnt := 0

	for _, o1 := range olist1 {
		_, err := ser.Run(ctx, o1)
		if err != nil {
			o1.Status = http.StatusInternalServerError
			o1.Error = err.Error()
			errlist.Append(err)
			errCnt++
		}
	}

	if errCnt == 1 {
		return errlist.First()
	}
	inst.body2.Intents = olist1
	return nil
}

func (inst *myRunIntentRequest) doPut() error {

	return fmt.Errorf("no impl")
}

func (inst *myRunIntentRequest) doDelete() error {

	return fmt.Errorf("no impl")
}
