package caches

import (
	"context"

	"github.com/starter-go/base/lang"
)

// Want 表示缓存请求参数
type Want struct {
	Context   context.Context
	ID        string
	ClassName string
	Class     Class
	Item      Item
	Loader    Loader
	Scope     Scope
	NotBefore lang.Time

	// Data any            // 数据对象
	attrs map[string]any // 其它扩展的参数
}

// SetAttr ...
func (inst *Want) SetAttr(name string, value any) {
	t := inst.attrs
	if t == nil {
		t = make(map[string]any)
		inst.attrs = t
	}
	t[name] = value
}

// GetAttr ...
func (inst *Want) GetAttr(name string) any {
	t := inst.attrs
	if t == nil {
		return nil
	}
	return t[name]
}
