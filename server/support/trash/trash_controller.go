package trash

import (
	"net/http"

	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/data/dao"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/vo"
	"github.com/gin-gonic/gin"
)

// TheTrashController Trash 控制器
type TheTrashController struct {
	markup.RestController `class:"rest-controller"`

	TrashSer  service.TrashService `inject:"#TrashService"`
	TrashDao  dao.TrashDAO         `inject:"#TrashDAO"`
	Responder glass.MainResponder  `inject:"#glass-main-responder"`
}

func (inst *TheTrashController) _Impl() glass.Controller {
	return inst
}

// Init 初始化
func (inst *TheTrashController) Init(ec glass.EngineConnection) error {

	ec = ec.RequestMapping("trash")

	ec.Handle(http.MethodGet, "", inst.handleGetListAll)
	ec.Handle(http.MethodDelete, "", inst.handleDeleteItems)
	ec.Handle(http.MethodPost, "recover", inst.handlePostRecoverItems)
	ec.Handle(http.MethodPost, "clean", inst.handlePostClean)

	// ec.Handle(http.MethodPost, "", inst.handlePost)
	// ec.Handle(http.MethodGet, ":id", inst.handleGetOne)
	// ec.Handle(http.MethodPut, ":id", inst.handlePut)
	// ec.Handle(http.MethodDelete, ":id", inst.handleDelete)

	return nil
}

func (inst *TheTrashController) handleGetListAll(c *gin.Context) {
	req := &myTrashRequest{
		gc:              c,
		controller:      inst,
		wantRequestID:   false,
		wantRequestBody: false,
	}
	err := req.open()
	if err == nil {
		err = req.doGetListAll()
	}
	req.send(err)
}

func (inst *TheTrashController) handlePostRecoverItems(c *gin.Context) {
	req := &myTrashRequest{
		gc:              c,
		controller:      inst,
		wantRequestID:   false,
		wantRequestBody: true,
	}
	err := req.open()
	if err == nil {
		err = req.doPostRecoverItems()
	}
	req.send(err)
}

func (inst *TheTrashController) handlePostClean(c *gin.Context) {
	req := &myTrashRequest{
		gc:              c,
		controller:      inst,
		wantRequestID:   false,
		wantRequestBody: true,
	}
	err := req.open()
	if err == nil {
		err = req.doPostClean()
	}
	req.send(err)
}

// func (inst *TheTrashController) handlePut(c *gin.Context) {
// 	req := &myTrashRequest{
// 		gc:              c,
// 		controller:      inst,
// 		wantRequestID:   true,
// 		wantRequestBody: true,
// 	}
// 	err := req.open()
// 	if err == nil {
// 		err = req.doPut()
// 	}
// 	req.send(err)
// }

func (inst *TheTrashController) handleDeleteItems(c *gin.Context) {
	req := &myTrashRequest{
		gc:              c,
		controller:      inst,
		wantRequestID:   true,
		wantRequestBody: false,
	}
	err := req.open()
	if err == nil {
		err = req.doDeleteItems()
	}
	req.send(err)
}

////////////////////////////////////////////////////////////////////////////////

type myTrashRequest struct {
	gc         *gin.Context
	controller *TheTrashController

	wantRequestID   bool
	wantRequestBody bool

	// id    dxo.TrashID
	body1 vo.Trash
	body2 vo.Trash
}

func (inst *myTrashRequest) open() error {

	c := inst.gc

	if inst.wantRequestID {
		// id := c.Param("id")
		// n, err := strconv.Atoi(id)
		// if err != nil {
		// 	return err
		// }
		// inst.id = dxo.TrashID(n)
	}

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *myTrashRequest) send(err error) {
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

// func (inst *myTrashRequest) doGetOne() error {
// 	id := inst.id
// 	ctx := inst.gc
// 	ser := inst.controller.TrashService
// 	o, err := ser.Find(ctx, id)
// 	if err != nil {
// 		return err
// 	}
// 	inst.body2.Trashs = []*dto.Trash{o}
// 	return nil
// }

func (inst *myTrashRequest) doGetListAll() error {
	// ctx := inst.gc
	ser := inst.controller.TrashDao
	list, err := ser.ListAll()
	if err != nil {
		return err
	}
	inst.body2.TrashItems = list
	return nil
}

// func (inst *myTrashRequest) doPost() error {
// 	return fmt.Errorf("no impl")
// }

// func (inst *myTrashRequest) doPut() error {
// 	return fmt.Errorf("no impl")
// }

func (inst *myTrashRequest) doPostRecoverItems() error {
	ser := inst.controller.TrashDao
	list := inst.body1.TrashItems
	return ser.Recover(list...)
}

func (inst *myTrashRequest) doPostClean() error {
	ser := inst.controller.TrashSer
	return ser.Clean()
}

func (inst *myTrashRequest) doDeleteItems() error {
	ser := inst.controller.TrashDao
	list := inst.body1.TrashItems
	return ser.Delete(list...)
}
