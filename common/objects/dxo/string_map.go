package dxo

import (
	"github.com/starter-go/application/properties"
)

// StringMap 表示一个stringify 的 map[string]string，它的格式参考git.config
type StringMap string

func (value StringMap) String() string {
	return string(value)
}

// Map ...
func (value StringMap) Map() map[string]string {
	text := value.String()
	p, _ := properties.Parse(text, nil)
	if p == nil {
		p = properties.NewTable(nil)
	}
	return p.Export(nil)
}

////////////////////////////////////////////////////////////////////////////////

// NewStringMap ... 根据 src 创建新的 StringMap
func NewStringMap(src map[string]string) StringMap {
	if src == nil {
		return ""
	}
	p := properties.NewTable(nil)
	p.Import(src)
	text := properties.Format(p)
	return StringMap(text)
}
