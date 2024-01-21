package web

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/starter-go/rbac"
)

type getterInner struct {
	c         *gin.Context
	errlist   []error
	required  bool
	valueFunc func(name string) (string, error)
}

func (inst *getterInner) init(c *gin.Context) {
	inst.c = c
	inst.errlist = make([]error, 0)
	inst.required = false
	inst.valueFunc = inst.getQuery
}

func (inst *getterInner) handleError(err error) {
	if err == nil {
		return
	}
	inst.errlist = append(inst.errlist, err)
}

func (inst *getterInner) lastError() error {
	list := inst.errlist
	count := len(list)
	if count > 0 {
		return list[count-1]
	}
	return nil
}

func (inst *getterInner) getParam(name string) (string, error) {
	value := inst.c.Param(name)
	if value == "" && inst.required {
		return "", fmt.Errorf("no required HTTP path param with name '%s'", name)
	}
	return value, nil
}

func (inst *getterInner) getQuery(name string) (string, error) {
	value := inst.c.Query(name)
	if value == "" && inst.required {
		return "", fmt.Errorf("no required HTTP query with name '%s'", name)
	}
	return value, nil
}

func (inst *getterInner) value(name string, callback func(str string) error) {
	str, err := inst.valueFunc(name)
	if err == nil {
		if str == "" {
			return
		}
		err = callback(str)
		if err != nil {
			const format = "bad value for HTTP request field '%s'='%s', error:%s"
			err = fmt.Errorf(format, name, str, err.Error())
		}
	}
	inst.handleError(err)
}

////////////////////////////////////////////////////////////////////////////////

// Getter 用于从 gin.Context 中 获取查询参数
type Getter struct {
	gg getterInner
}

// Required 用于设置 Required = true，该参数指示接下来的字段是否为必须的
func (inst *Getter) Required() *Getter {
	inst.gg.required = true
	return inst
}

// Optional 用于设置 Required = false，该参数指示接下来的字段是否为必须的
func (inst *Getter) Optional() *Getter {
	inst.gg.required = false
	return inst
}

func (inst *Getter) Error() error {
	return inst.gg.lastError()
}

// GetString ...
func (inst *Getter) GetString(name string) string {
	value := ""
	inst.gg.value(name, func(str string) error {
		value = str
		return nil
	})
	return value
}

// GetInt ...
func (inst *Getter) GetInt(name string) int {
	value := 0
	inst.gg.value(name, func(str string) error {
		v, err := strconv.Atoi(str)
		value = v
		return err
	})
	return value
}

// GetBool ...
func (inst *Getter) GetBool(name string) bool {
	value := false
	inst.gg.value(name, func(str string) error {
		v, err := strconv.ParseBool(str)
		value = v
		return err
	})
	return value
}

// GetPagination ...
func (inst *Getter) GetPagination() rbac.Pagination {
	var value rbac.Pagination
	page := inst.GetInt("page")
	size := inst.GetInt("size")
	value.Page = int64(page)
	value.Size = size
	return value
}

////////////////////////////////////////////////////////////////////////////////

// NewQueryGetter ...
func NewQueryGetter(c *gin.Context) *Getter {
	g := new(Getter)
	g.gg.init(c)
	g.gg.valueFunc = g.gg.getQuery
	return g
}

// NewParamGetter ...
func NewParamGetter(c *gin.Context) *Getter {
	g := new(Getter)
	g.gg.init(c)
	g.gg.valueFunc = g.gg.getParam
	return g
}

////////////////////////////////////////////////////////////////////////////////
