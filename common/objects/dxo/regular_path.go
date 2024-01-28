package dxo

import (
	"runtime"
	"strings"
)

// RegularPath 用于在数据库中对路径排重
type RegularPath string

// NewRegularPath 为 path 取对应的 RegularPath 值
func NewRegularPath(path string) RegularPath {

	path = strings.TrimSpace(path)

	if runtime.GOOS == "windows" {
		path = strings.ToLower(path)
	}

	return RegularPath(path)
}
