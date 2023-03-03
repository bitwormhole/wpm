package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// GinUtils ...
type GinUtils struct{}

// HasFlag 判断传入的 query 中是否包含指定的标志位
func (inst *GinUtils) HasFlag(c *gin.Context, name string) bool {
	value := c.Query(name)
	if value == "1" {
		return true
	}
	b, _ := strconv.ParseBool(value)
	return b
}
