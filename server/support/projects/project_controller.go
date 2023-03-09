package projects

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

// ProjectController 仓库控制器
type ProjectController struct {
	markup.RestController `class:"rest-controller"`

	ProjectService service.ProjectService `inject:"#ProjectService"`
	Responder      glass.MainResponder    `inject:"#glass-main-responder"`
}

func (inst *ProjectController) _Impl() glass.Controller {
	return inst
}

// Init 初始化
func (inst *ProjectController) Init(ec glass.EngineConnection) error {

	ec = ec.RequestMapping("projects")

	ec.Handle(http.MethodGet, "", inst.handleGetList)
	ec.Handle(http.MethodPost, "", inst.handlePost)

	ec.Handle(http.MethodGet, ":id", inst.handleGetOne)
	ec.Handle(http.MethodPut, ":id", inst.handlePut)
	ec.Handle(http.MethodDelete, ":id", inst.handleDelete)

	return nil
}

func (inst *ProjectController) handleGetList(c *gin.Context) {
	req := &myProjectRequest{
		gc:                 c,
		controller:         inst,
		wantRequestID:      false,
		wantRequestBody:    false,
		wantRequestIDs:     true,
		wantRequestOptions: true,
	}
	err := req.open()
	if err == nil {
		wantAll := req.options.All
		if wantAll {
			err = req.doGetAll()
		} else {
			err = req.doGetListByIds()
		}
	}
	req.send(err)
}

func (inst *ProjectController) handleGetOne(c *gin.Context) {
	req := &myProjectRequest{
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

func (inst *ProjectController) handlePost(c *gin.Context) {
	req := &myProjectRequest{
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

func (inst *ProjectController) handlePut(c *gin.Context) {
	req := &myProjectRequest{
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

func (inst *ProjectController) handleDelete(c *gin.Context) {
	req := &myProjectRequest{
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

type myProjectRequest struct {
	gc         *gin.Context
	controller *ProjectController

	wantRequestID      bool
	wantRequestIDs     bool
	wantRequestBody    bool
	wantRequestOptions bool

	options service.ProjectOptions

	ids   []dxo.ProjectID
	id    dxo.ProjectID
	body1 vo.Project
	body2 vo.Project
}

func (inst *myProjectRequest) open() error {

	c := inst.gc

	if inst.wantRequestID {
		id := c.Param("id")
		n, err := strconv.Atoi(id)
		if err != nil {
			return err
		}
		inst.id = dxo.ProjectID(n)
	}

	if inst.wantRequestIDs {
		idstr := c.Query("ids")
		ids, err := inst.parseIds(idstr)
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

	if inst.wantRequestOptions {
		inst.options.All = inst.hasFlag(c, "all")
		inst.options.WithFileState = inst.hasFlag(c, "with-file-state")
		inst.options.WithGitStatus = inst.hasFlag(c, "with-git-status")
		// inst.options.WithProjects = inst.hasFlag(c, "with-projects")
	}

	return nil
}

func (inst *myProjectRequest) parseIds(str string) ([]dxo.ProjectID, error) {
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
	dst := make([]dxo.ProjectID, 0)
	for _, n := range nlist {
		dst = append(dst, dxo.ProjectID(n))
	}
	return dst, nil
}

func (inst *myProjectRequest) hasFlag(c *gin.Context, name string) bool {
	return (&utils.GinUtils{}).HasFlag(c, name)
}

func (inst *myProjectRequest) send(err error) {
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

func (inst *myProjectRequest) doGetOne() error {
	id := inst.id
	ctx := inst.gc
	ser := inst.controller.ProjectService
	opt := &inst.options
	o, err := ser.Find(ctx, id, opt)
	if err != nil {
		return err
	}
	inst.body2.Projects = []*dto.Project{o}
	return nil
}

func (inst *myProjectRequest) doGetListByIds() error {
	ctx := inst.gc
	ser := inst.controller.ProjectService
	opt := &inst.options
	ids := inst.ids
	list, err := ser.ListByIds(ctx, ids, opt)
	if err != nil {
		return err
	}
	inst.body2.Projects = list
	return nil
}

func (inst *myProjectRequest) doGetAll() error {
	ctx := inst.gc
	ser := inst.controller.ProjectService
	opt := &inst.options
	list, err := ser.ListAll(ctx, opt)
	if err != nil {
		return err
	}
	inst.body2.Projects = list
	return nil
}

func (inst *myProjectRequest) doPost() error {
	ctx := inst.gc
	ser := inst.controller.ProjectService
	item1 := inst.body1.Projects[0]
	item2, err := ser.Insert(ctx, item1)
	if err != nil {
		return err
	}
	inst.body2.Projects = []*dto.Project{item2}
	return nil
}

func (inst *myProjectRequest) doPut() error {
	ctx := inst.gc
	ser := inst.controller.ProjectService
	id := inst.id
	item1 := inst.body1.Projects[0]
	item2, err := ser.Update(ctx, id, item1)
	if err != nil {
		return err
	}
	inst.body2.Projects = []*dto.Project{item2}
	return nil
}

func (inst *myProjectRequest) doDelete() error {
	ctx := inst.gc
	ser := inst.controller.ProjectService
	id := inst.id
	return ser.Remove(ctx, id)
}
