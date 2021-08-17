package controller

import (
	"net/http"

	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/application"
	"github.com/gin-gonic/gin"
)

// 有了devtools，这个鬼东西可以废掉了

type DebugController struct {
	AppContext application.Context
}

func (inst *DebugController) _Impl() glass.Controller {
	return inst
}

func (inst *DebugController) Init(e glass.EngineConnection) error {

	e = e.RequestMapping("debug")

	e.Handle(http.MethodGet, "", func(c *gin.Context) { inst.doGetDebugInfo(c) })
	e.Handle(http.MethodPost, "", func(c *gin.Context) {})

	return nil
}

func (inst *DebugController) doGetDebugInfo(c *gin.Context) {

	h := gin.H{}
	app := inst.AppContext
	props := app.GetProperties().Export(nil)
	env := app.GetEnvironment().Export(nil)

	// properties
	h["properties"] = props
	h["env"] = env

	c.JSON(200, h)
}
