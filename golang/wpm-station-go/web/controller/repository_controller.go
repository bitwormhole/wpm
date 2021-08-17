package controller

import (
	"context"
	"net/http"

	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/service"
	"github.com/bitwormhole/wpm/web/vo"
	"github.com/gin-gonic/gin"
)

// RepositoryController 仓库控制器
type RepositoryController struct {
	markup.Component `class:"rest-controller"`
	RepoService      service.RepositoryService `inject:"#repository-service"`
}

func (inst *RepositoryController) _Impl() glass.Controller {
	return inst
}

// Init 初始化
func (inst *RepositoryController) Init(ec glass.EngineConnection) error {

	ec = ec.RequestMapping("repositories")

	ec.Handle(http.MethodGet, "", func(c *gin.Context) { inst.doGetList(c) })
	ec.Handle(http.MethodGet, ":id", func(c *gin.Context) { inst.doGetOne(c) })

	return nil
}

////////////////////
// do

func (inst *RepositoryController) doGetOne(c *gin.Context) {

	data := &vo.Repository{}
	cc := glass.CommonContext(c)
	data.ParamID = c.Params.ByName("id")

	resp, err := inst.handleRepoGetOne(cc, data)

	if err == nil {
		resp.Status = 200
	} else {
		resp.Status = 500
		resp.Error = err.Error()
	}
	c.JSON(200, data)
}

func (inst *RepositoryController) doGetList(c *gin.Context) {

	data := &vo.Repository{}
	cc := glass.CommonContext(c)

	resp, err := inst.handleRepoGetList(cc, data)

	if err == nil {
		resp.Status = 200
	} else {
		resp.Status = 500
		resp.Error = err.Error()
	}
	c.JSON(200, data)
}

///////////////////////
// handle

func (inst *RepositoryController) handleRepoGetOne(cc context.Context, input *vo.Repository) (*vo.Repository, error) {

	return input, nil
}

func (inst *RepositoryController) handleRepoGetList(cc context.Context, input *vo.Repository) (*vo.Repository, error) {
	all := inst.RepoService.ListAll()
	resp := &vo.Repository{}
	resp.Repositories = all
	return resp, nil
}
