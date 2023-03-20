package utils

import (
	"strings"

	"github.com/bitwormhole/starter/collection"
)

// PropertiesUtil ...
type PropertiesUtil struct{}

// GetNameList 读取 src 中，通过 prefix 和 suffix 指定类型的名单
func (inst *PropertiesUtil) GetNameList(src collection.Properties, prefix string, suffix string) []string {
	table := src.Export(nil)
	namelist := make([]string, 0)
	for key := range table {
		if strings.HasPrefix(key, prefix) && strings.HasSuffix(key, suffix) {
			name := key[len(prefix) : len(key)-len(suffix)]
			if name != "" {
				namelist = append(namelist, name)
			}
		}
	}
	return namelist
}
