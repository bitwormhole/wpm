package filter

import (
	"net/http"

	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/vlog"
	"github.com/gin-gonic/gin"
)

// HostFilter 是 host 过滤器
type HostFilter struct {
	markup.RestController `class:"rest-controller"`

	// IntentService service.IntentService `inject:"#IntentService"`
	// Responder     glass.MainResponder   `inject:"#glass-main-responder"`
}

func (inst *HostFilter) _Impl() glass.Controller {
	return inst
}

// Init 初始化
func (inst *HostFilter) Init(ec glass.EngineConnection) error {
	ec.Filter(0, inst.filter)
	return nil
}

func (inst *HostFilter) filter(c *gin.Context) {
	if inst.accept(c) {
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusForbidden)
	}
}

func (inst *HostFilter) accept(c *gin.Context) bool {
	ip := c.ClientIP()
	vlog.Warn("http access with host:", ip)
	if ip == "127.0.0.1" {
		return true
	}
	return false
}
