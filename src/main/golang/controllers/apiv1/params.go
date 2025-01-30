package apiv1

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/starter-go/rbac"
)

// ParamGetter ...
type ParamGetter struct {
	gc *gin.Context
}

// GetQueryString ...
func (inst *ParamGetter) GetQueryString(name string, def string) string {
	val := inst.gc.Query(name)
	if val == "" {
		return def
	}
	return val
}

// GetQueryBool ...
func (inst *ParamGetter) GetQueryBool(name string, def bool) bool {
	val := inst.GetQueryString(name, "")
	b, err := strconv.ParseBool(val)
	if err == nil {
		return b
	}
	return def
}

// GetQueryPagination ...
func (inst *ParamGetter) GetQueryPagination(dst *rbac.Pagination) error {
	if dst == nil {
		return fmt.Errorf("no result struct")
	}
	page := inst.GetQueryInt64("page", 1)
	size := inst.GetQueryInt("size", 5)
	dst.Page = page
	dst.Size = size
	return nil
}

// GetQueryInt ...
func (inst *ParamGetter) GetQueryInt(name string, def int) int {
	val := inst.GetQueryString(name, "")
	n, err := strconv.Atoi(val)
	if err == nil {
		return n
	}
	return def
}

// GetQueryInt64 ...
func (inst *ParamGetter) GetQueryInt64(name string, def int64) int64 {
	const (
		base    = 10
		bitSize = 64
	)
	val := inst.GetQueryString(name, "")
	n, err := strconv.ParseInt(val, base, bitSize)
	if err == nil {
		return n
	}
	return def
}
