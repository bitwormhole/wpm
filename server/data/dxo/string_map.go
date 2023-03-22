package dxo

import "github.com/bitwormhole/starter/collection"

// StringMap 表示一个stringify 的 map[string]string，它的格式参考git.config
type StringMap string

func (value StringMap) String() string {
	return string(value)
}

// Map ...
func (value StringMap) Map() map[string]string {
	text := value.String()
	p, _ := collection.ParseProperties(text, nil)
	if p == nil {
		p = collection.CreateProperties()
	}
	return p.Export(nil)
}

////////////////////////////////////////////////////////////////////////////////

// NewStringMap ... 根据 src 创建新的 StringMap
func NewStringMap(src map[string]string) StringMap {
	if src == nil {
		return ""
	}
	p := collection.CreateProperties()
	p.Import(src)
	text := collection.FormatProperties(p)
	return StringMap(text)
}
