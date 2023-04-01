package namespaces

import (
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/utils"
	"github.com/bitwormhole/wpm/server/web/dto"
	"github.com/bitwormhole/wpm/server/web/vo"
	"github.com/gin-gonic/gin"
)

// NamespaceController 命名空间控制器
type NamespaceController struct {
	markup.RestController `class:"rest-controller"`

	NamespaceService service.NamespaceService `inject:"#NamespaceService"`
	Responder        glass.MainResponder      `inject:"#glass-main-responder"`
}

func (inst *NamespaceController) _Impl() glass.Controller {
	return inst
}

// Init 初始化
func (inst *NamespaceController) Init(ec glass.EngineConnection) error {

	ec = ec.RequestMapping("namespaces")

	ec.Handle(http.MethodGet, "", inst.handleGetList)
	ec.Handle(http.MethodPost, "", inst.handlePost)

	ec.Handle(http.MethodGet, ":id", inst.handleGetOne)
	ec.Handle(http.MethodPut, ":id", inst.handlePut)
	ec.Handle(http.MethodDelete, ":id", inst.handleDelete)

	return nil
}

func (inst *NamespaceController) handleGetList(c *gin.Context) {
	req := &myNamespaceRequest{
		gc:                 c,
		controller:         inst,
		wantRequestID:      false,
		wantRequestBody:    false,
		wantRequestIDs:     true,
		wantRequestOptions: true,
	}
	err := req.open()
	if err == nil {
		err = req.doGetAll()
	}
	req.send(err)
}

func (inst *NamespaceController) handleGetOne(c *gin.Context) {
	req := &myNamespaceRequest{
		gc:                 c,
		controller:         inst,
		wantRequestID:      true,
		wantRequestBody:    false,
		wantRequestOptions: true,
	}
	err := req.open()
	if err == nil {
		err = req.doGetOne()
	}
	req.send(err)
}

func (inst *NamespaceController) handlePost(c *gin.Context) {
	req := &myNamespaceRequest{
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

func (inst *NamespaceController) handlePut(c *gin.Context) {
	req := &myNamespaceRequest{
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

func (inst *NamespaceController) handleDelete(c *gin.Context) {
	req := &myNamespaceRequest{
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

type myNamespaceRequest struct {
	gc         *gin.Context
	controller *NamespaceController

	wantRequestID      bool
	wantRequestIDs     bool
	wantRequestBody    bool
	wantRequestOptions bool

	// options service.NamespaceOptions

	ids   []dxo.NamespaceID
	id    dxo.NamespaceID
	body1 vo.Namespace
	body2 vo.Namespace
}

func (inst *myNamespaceRequest) open() error {

	c := inst.gc

	if inst.wantRequestID {
		id := c.Param("id")
		n, err := strconv.Atoi(id)
		if err != nil {
			return err
		}
		inst.id = dxo.NamespaceID(n)
	}

	// if inst.wantRequestIDs {
	// 	idstr := c.Query("ids")
	// 	ids, err := inst.parseIds(idstr)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	inst.ids = ids
	// }

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	// if inst.wantRequestOptions {
	// 	inst.options.All = inst.hasFlag(c, "all")
	// 	inst.options.WithFileState = inst.hasFlag(c, "with-file-state")
	// 	inst.options.WithGitStatus = inst.hasFlag(c, "with-git-status")
	// 	// inst.options.WithNamespaces = inst.hasFlag(c, "with-Namespaces")
	// }

	return nil
}

func (inst *myNamespaceRequest) parseIds(str string) ([]dxo.NamespaceID, error) {
	if str == "" {
		return nil, nil
	}
	nlist := make([]int, 0)
	src := strings.Split(str, ".")
	for _, item := range src {
		if item == "" {
			continue
		}
		n, err := strconv.ParseInt(item, 10, 32)
		if err != nil {
			return nil, err
		}
		nlist = append(nlist, int(n))
	}
	sort.Ints(nlist)
	dst := make([]dxo.NamespaceID, 0)
	for _, n := range nlist {
		dst = append(dst, dxo.NamespaceID(n))
	}
	return dst, nil
}

func (inst *myNamespaceRequest) hasFlag(c *gin.Context, name string) bool {
	return (&utils.GinUtils{}).HasFlag(c, name)
}

func (inst *myNamespaceRequest) send(err error) {
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

func (inst *myNamespaceRequest) doGetOne() error {
	id := inst.id
	ctx := inst.gc
	ser := inst.controller.NamespaceService
	// opt := &inst.options
	o, err := ser.Find(ctx, id)
	if err != nil {
		return err
	}
	inst.body2.Namespaces = []*dto.Namespace{o}
	return nil
}

func (inst *myNamespaceRequest) doGetAll() error {
	ctx := inst.gc
	ser := inst.controller.NamespaceService
	// opt := &inst.options
	list, err := ser.ListAll(ctx)
	if err != nil {
		return err
	}
	inst.body2.Namespaces = list
	return nil
}

func (inst *myNamespaceRequest) doPost() error {
	ctx := inst.gc
	ser := inst.controller.NamespaceService
	item1 := inst.body1.Namespaces[0]
	item2, err := ser.Insert(ctx, item1)
	if err != nil {
		return err
	}
	inst.body2.Namespaces = []*dto.Namespace{item2}
	return nil
}

func (inst *myNamespaceRequest) doPut() error {
	ctx := inst.gc
	ser := inst.controller.NamespaceService
	id := inst.id
	item1 := inst.body1.Namespaces[0]
	item2, err := ser.Update(ctx, id, item1)
	if err != nil {
		return err
	}
	inst.body2.Namespaces = []*dto.Namespace{item2}
	return nil
}

func (inst *myNamespaceRequest) doDelete() error {
	ctx := inst.gc
	ser := inst.controller.NamespaceService
	id := inst.id
	return ser.Remove(ctx, id)
}
