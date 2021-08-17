package controller

import (
	"net/http"

	"github.com/bitwormhole/gitlib/repository"
	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/wpm/web/vo"
	"github.com/gin-gonic/gin"
)

type TestGitController struct {
	AppContext application.Context
	RM         repository.Manager
}

func (inst *TestGitController) _Impl() glass.Controller {
	return inst
}

func (inst *TestGitController) Init(e glass.EngineConnection) error {

	e = e.RequestMapping("test")

	// e.Handle(http.MethodGet, "", func(c *gin.Context) { inst.doGetDebugInfo(c) })
	e.Handle(http.MethodPost, "git-repo-locator", func(c *gin.Context) { inst.doPostTestGitRepoLocator(c) })

	return nil
}

func (inst *TestGitController) doPostTestGitRepoLocator(c *gin.Context) {

	req := &vo.TestVO{}
	resp := &vo.TestVO{}
	c.BindJSON(req)

	pwd := fs.Default().GetPath(req.Locator.Path)

	driver, err := inst.RM.FindDriver(pwd.URI())
	if err != nil {
		resp.Error = err.Error()
		c.JSON(500, resp)
		return
	}

	location, err := driver.Locator().Locate(pwd.URI())
	if err != nil {
		resp.Error = err.Error()
		c.JSON(500, resp)
		return
	}

	kvs := make(map[string]string)
	resp.Locator = req.Locator
	resp.Locator.Results = kvs

	inst.fill(kvs, "pwd", location.PWD)
	inst.fill(kvs, ".git", location.DotGit)
	inst.fill(kvs, "working", location.WorkingDirectory)
	inst.fill(kvs, "submodule", location.SubmoduleDirectory)
	inst.fill(kvs, "worktree", location.WorktreeDirectory)
	inst.fill(kvs, "core", location.CoreDirectory)

	c.JSON(200, resp)
}

func (inst *TestGitController) fill(kvs map[string]string, key string, path fs.Path) {
	val := ""
	if path != nil {
		val = path.Path()
	}
	kvs[key] = val
}
