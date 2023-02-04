package filter

import (
	"net/http"

	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/markup"
	"github.com/gin-gonic/gin"
)

// HTTP404Filter 是 http 404 error 过滤器
type HTTP404Filter struct {
	markup.RestController `class:"rest-controller"`

	Context application.Context `inject:"context"`

	pageData []byte
}

func (inst *HTTP404Filter) _Impl() glass.Controller {
	return inst
}

// Init 初始化
func (inst *HTTP404Filter) Init(ec glass.EngineConnection) error {
	ec.HandleNoResource(0, inst.handle)
	return nil
}

func (inst *HTTP404Filter) handle(c *gin.Context) {
	data := inst.getPage()
	c.Data(http.StatusNotFound, "text/html", data)
}

func (inst *HTTP404Filter) getPage() []byte {
	data := inst.pageData
	if data == nil {
		data = inst.loadPage()
		inst.pageData = data
	}
	return data
}

func (inst *HTTP404Filter) loadPage() []byte {
	name := "res:///static/" + "wpm-ui-http-404-index.html"
	d, err := inst.Context.GetResources().GetBinary(name)
	if err != nil {
		text := "<html><body><h1>HTTP 404</h1></body></html>"
		return []byte(text)
	}
	return d
}
