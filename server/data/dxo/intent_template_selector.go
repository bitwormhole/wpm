package dxo

import (
	"fmt"
	"strings"
)

////////////////////////////////////////////////////////////////////////////////

// IntentTemplateSelector 用字符串表示一个 Intent 模板选择器
type IntentTemplateSelector string

func (value IntentTemplateSelector) String() string {
	return string(value)
}

// NewIntentTemplateSelector make a new IntentTemplateSelector
func NewIntentTemplateSelector(sel string) IntentTemplateSelector {
	return IntentTemplateSelector(sel)
}

////////////////////////////////////////////////////////////////////////////////

// IntentTemplateSelectorBuilder 用来创建 IntentTemplateSelector
type IntentTemplateSelectorBuilder struct {
	Method string // [open|insert|update|edit|...]
	Type   string // [project-type|content-type|target-type]
	With   string // [exe-name]
}

func (inst *IntentTemplateSelectorBuilder) normalizeName(name string) string {
	name = strings.TrimSpace(name)
	name = strings.ToLower(name)
	return name
}

// Create build a new IntentTemplateSelector
func (inst *IntentTemplateSelectorBuilder) Create() IntentTemplateSelector {

	method := inst.normalizeName(inst.Method)
	atype := inst.normalizeName(inst.Type)
	with := inst.normalizeName(inst.With)

	if method == "" {
		method = "open"
	}

	sel := fmt.Sprintf(".%v(%v).with(%v)", method, atype, with)
	sel = strings.ToLower(sel)
	return IntentTemplateSelector(sel)
}
