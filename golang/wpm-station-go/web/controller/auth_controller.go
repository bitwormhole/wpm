package controller

import (
	"context"
	"net/http"

	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/wpm/web/vo"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
}

func (inst *AuthController) _Impl() glass.Controller {
	return inst
}

func (inst *AuthController) Init(e glass.EngineConnection) error {

	e = e.RequestMapping("auth")

	e.Handle(http.MethodGet, "", func(c *gin.Context) { inst.doGet(c) })
	e.Handle(http.MethodPost, "", func(c *gin.Context) { inst.doPost(c) })

	return nil
}

////////////////////////
// do

func (inst *AuthController) doPost(c *gin.Context) {

	cc := glass.CommonContext(c)
	data := &vo.Auth{}
	c.BindJSON(data)

	resp, err := inst.handleAuthPost(cc, data)

	if err != nil {
		resp.Status = 500
		resp.Error = err.Error()
	} else {
		resp.Status = 200
	}
	c.JSON(resp.Status, resp)
}

func (inst *AuthController) doGet(c *gin.Context) {

	cc := glass.CommonContext(c)
	data := &vo.Auth{}
	// c.BindJSON(data)

	resp, err := inst.handleAuthGet(cc, data)

	if err != nil {
		resp.Status = 500
		resp.Error = err.Error()
	} else {
		resp.Status = 200
	}
	c.JSON(resp.Status, resp)
}

///////////////////////
// handle

func (inst *AuthController) handleAuthGet(cc context.Context, input *vo.Auth) (*vo.Auth, error) {

	return input, nil
}

func (inst *AuthController) handleAuthPost(cc context.Context, input *vo.Auth) (*vo.Auth, error) {

	return input, nil
}
