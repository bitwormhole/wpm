package utils

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

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

// ParseIDs 解析类似“1.2.3.4.5.”的参数
func (inst *GinUtils) ParseIDs(str string, sep string, fn func(n int)) error {
	if fn == nil {
		return fmt.Errorf("callback func is nil")
	}
	nlist := make([]int, 0)
	src := strings.Split(str, sep)
	for _, item := range src {
		item = strings.TrimSpace(item)
		if item == "" {
			continue
		}
		n, err := strconv.ParseInt(item, 10, 32)
		if err != nil {
			return err
		}
		nlist = append(nlist, int(n))
	}
	sort.Ints(nlist)
	for _, n := range nlist {
		fn(n)
	}
	return nil
}
