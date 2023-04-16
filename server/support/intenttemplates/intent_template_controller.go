package intenttemplates

import (
	"net/http"
	"strconv"

	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/utils"
	"github.com/bitwormhole/wpm/server/web/dto"
	"github.com/bitwormhole/wpm/server/web/vo"
	"github.com/gin-gonic/gin"
)

// IntentTemplateController IntentTemplate 控制器
type IntentTemplateController struct {
	markup.RestController `class:"rest-controller"`

	IntentTemplateService service.IntentTemplateService `inject:"#IntentTemplateService"`
	Responder             glass.MainResponder           `inject:"#glass-main-responder"`
}

func (inst *IntentTemplateController) _Impl() glass.Controller {
	return inst
}

// Init 初始化
func (inst *IntentTemplateController) Init(ec glass.EngineConnection) error {

	ec = ec.RequestMapping("intent-templates")

	ec.Handle(http.MethodGet, "", inst.handleGetList)
	ec.Handle(http.MethodGet, ":id", inst.handleGetOne)
	ec.Handle(http.MethodGet, "macro/properties", inst.handleGetMacroProps)

	ec.Handle(http.MethodPost, "", inst.handlePostInsert)
	ec.Handle(http.MethodPost, "import-preset", inst.handlePostImportPreset)

	ec.Handle(http.MethodPut, ":id", inst.handlePut)

	ec.Handle(http.MethodDelete, "", inst.handleDeleteByIDs)
	ec.Handle(http.MethodDelete, ":id", inst.handleDeleteOne)

	return nil
}

func (inst *IntentTemplateController) handleGetList(c *gin.Context) {
	req := &myIntentTemplateRequest{
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

func (inst *IntentTemplateController) handleGetOne(c *gin.Context) {
	req := &myIntentTemplateRequest{
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

func (inst *IntentTemplateController) handleGetMacroProps(c *gin.Context) {
	req := &myIntentTemplateRequest{
		gc:              c,
		controller:      inst,
		wantRequestID:   false,
		wantRequestBody: false,
	}
	err := req.open()
	if err == nil {
		err = req.doGetMacroProps()
	}
	req.send(err)
}

func (inst *IntentTemplateController) handlePostInsert(c *gin.Context) {
	req := &myIntentTemplateRequest{
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

func (inst *IntentTemplateController) handlePostImportPreset(c *gin.Context) {
	req := &myIntentTemplateRequest{
		gc:              c,
		controller:      inst,
		wantRequestID:   false,
		wantRequestBody: true,
	}
	err := req.open()
	if err == nil {
		err = req.doPostImportPreset()
	}
	req.send(err)
}

func (inst *IntentTemplateController) handlePut(c *gin.Context) {
	req := &myIntentTemplateRequest{
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

func (inst *IntentTemplateController) handleDeleteOne(c *gin.Context) {
	req := &myIntentTemplateRequest{
		gc:              c,
		controller:      inst,
		wantRequestID:   true,
		wantRequestBody: false,
	}
	err := req.open()
	if err == nil {
		id := req.id
		err = req.doDelete(id)
	}
	req.send(err)
}

func (inst *IntentTemplateController) handleDeleteByIDs(c *gin.Context) {
	req := &myIntentTemplateRequest{
		gc:              c,
		controller:      inst,
		wantRequestIDs:  true,
		wantRequestBody: false,
	}
	err := req.open()
	if err == nil {
		ids := req.ids
		err = req.doDelete(ids...)
	}
	req.send(err)
}

////////////////////////////////////////////////////////////////////////////////

type myIntentTemplateRequest struct {
	gc         *gin.Context
	controller *IntentTemplateController

	wantRequestID   bool
	wantRequestIDs  bool
	wantRequestBody bool

	id    dxo.IntentTemplateID
	ids   []dxo.IntentTemplateID
	body1 vo.IntentTemplate
	body2 vo.IntentTemplate
}

func (inst *myIntentTemplateRequest) open() error {

	c := inst.gc

	if inst.wantRequestID {
		id := c.Param("id")
		n, err := strconv.Atoi(id)
		if err != nil {
			return err
		}
		inst.id = dxo.IntentTemplateID(n)
	}

	if inst.wantRequestIDs {
		idsstr := c.Query("ids")
		ids, err := inst.parseIds(idsstr)
		if err != nil {
			return err
		}
		inst.ids = ids
	}

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *myIntentTemplateRequest) parseIds(str string) ([]dxo.IntentTemplateID, error) {
	list := make([]dxo.IntentTemplateID, 0)
	err := (&utils.GinUtils{}).ParseIDs(str, ".", func(n int) {
		list = append(list, dxo.IntentTemplateID(n))
	})
	return list, err
}

func (inst *myIntentTemplateRequest) send(err error) {
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

func (inst *myIntentTemplateRequest) doGetOne() error {
	id := inst.id
	ctx := inst.gc
	ser := inst.controller.IntentTemplateService
	o, err := ser.Find(ctx, id)
	if err != nil {
		return err
	}
	inst.body2.Templates = []*dto.IntentTemplate{o}
	return nil
}

func (inst *myIntentTemplateRequest) doGetMacroProps() error {
	ctx := inst.gc
	ser := inst.controller.IntentTemplateService
	props, err := ser.ListMacroProperties(ctx)
	if err != nil {
		return err
	}
	inst.body2.MacroProps = props
	return nil
}

func (inst *myIntentTemplateRequest) doGetList() error {
	ctx := inst.gc
	ser := inst.controller.IntentTemplateService
	list, err := ser.ListAll(ctx)
	if err != nil {
		return err
	}
	inst.body2.Templates = list
	return nil
}

func (inst *myIntentTemplateRequest) doPostInsert() error {
	ctx := inst.gc
	ser := inst.controller.IntentTemplateService
	o1 := inst.body1.Templates[0]
	o2, err := ser.Insert(ctx, o1)
	if err != nil {
		return err
	}
	inst.body2.Templates = []*dto.IntentTemplate{o2}
	return nil
}

func (inst *myIntentTemplateRequest) doPostImportPreset() error {
	ctx := inst.gc
	ser := inst.controller.IntentTemplateService
	return ser.ImportPreset(ctx)
}

func (inst *myIntentTemplateRequest) doPut() error {
	ctx := inst.gc
	ser := inst.controller.IntentTemplateService
	id := inst.id
	o1 := inst.body1.Templates[0]
	o2, err := ser.Update(ctx, id, o1)
	if err != nil {
		return err
	}
	inst.body2.Templates = []*dto.IntentTemplate{o2}
	return nil
}

func (inst *myIntentTemplateRequest) doDelete(ids ...dxo.IntentTemplateID) error {
	ctx := inst.gc
	ser := inst.controller.IntentTemplateService
	elist := utils.ErrorList{}
	for _, id := range ids {
		err := ser.Remove(ctx, id)
		elist.Append(err)
	}
	return elist.First()
}
