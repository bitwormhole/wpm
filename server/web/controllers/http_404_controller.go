package controllers

import (
	"net/http"
	"strings"

	"github.com/bitwormhole/wpm/common/objects/vo"
	"github.com/gin-gonic/gin"
	"github.com/starter-go/application"
	"github.com/starter-go/libgin"
)

// HTTP404Controller ...
type HTTP404Controller struct {

	//starter:component
	_as func(libgin.Controller) //starter:as(".")

	AC application.Context //starter:inject("context")

	cachedHTMLContent string
}

func (inst *HTTP404Controller) _impl() libgin.Controller {
	return inst
}

// Registration ...
func (inst *HTTP404Controller) Registration() *libgin.ControllerRegistration {
	return &libgin.ControllerRegistration{Route: inst.route}
}

func (inst *HTTP404Controller) route(rp libgin.RouterProxy) error {
	rp.NoMethod(inst.handleNotFound)
	rp.NoRoute(inst.handleNotFound)
	return nil
}

func (inst *HTTP404Controller) handleNotFound(c *gin.Context) {
	if inst.isAPIRequest(c) {
		inst.sendErrorJSON(c)
	} else {
		inst.sendIndexPage(c)
	}
}

func (inst *HTTP404Controller) isAPIRequest(c *gin.Context) bool {
	path := c.Request.URL.Path
	return strings.HasPrefix(path, "/api/")
}

func (inst *HTTP404Controller) sendErrorJSON(c *gin.Context) {
	code := http.StatusNotFound
	obj := new(vo.Base)
	obj.Status = code
	obj.Message = http.StatusText(code)
	c.JSON(code, obj)
}

func (inst *HTTP404Controller) sendIndexPage(c *gin.Context) {
	code := http.StatusOK
	ctype := "text/html"
	content := inst.getIndexPageText()
	c.Data(code, ctype, []byte(content))
}

func (inst *HTTP404Controller) getIndexPageText() string {
	text := inst.cachedHTMLContent
	if text == "" {
		data, _ := inst.loadIndexPageText()
		if data == "" {
			b := new(strings.Builder)
			b.WriteString("<html><body>")
			b.WriteString("<p>HTTP 404 Not Found</p>")
			b.WriteString("</body></html>")
			data = b.String()
		}
		text = data
		inst.cachedHTMLContent = data
	}
	return text
}

func (inst *HTTP404Controller) loadIndexPageText() (string, error) {
	path := "res:///static/index.html"
	res, err := inst.AC.GetResources().GetResource(path)
	if err != nil {
		return "", err
	}
	return res.ReadText()
}

////////////////////////////////////////////////////////////////////////////////
