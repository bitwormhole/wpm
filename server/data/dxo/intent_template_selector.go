package dxo

import "strings"

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
	Action string // [open|insert|update|edit|...]
	Target string // [repository|worktree|project||]
	With   string // [exe-name]
}

func (inst *IntentTemplateSelectorBuilder) normalizeName(name string) string {
	name = strings.TrimSpace(name)
	name = strings.ToLower(name)
	return name
}

// Create build a new IntentTemplateSelector
func (inst *IntentTemplateSelectorBuilder) Create() IntentTemplateSelector {

	action := inst.normalizeName(inst.Action)
	target := inst.normalizeName(inst.Target)
	with := inst.normalizeName(inst.With)
	builder := strings.Builder{}

	builder.WriteString(".")
	builder.WriteString(action)
	builder.WriteString("(")
	builder.WriteString(target)
	builder.WriteString(").with(")
	builder.WriteString(with)
	builder.WriteString(")")

	sel := builder.String()
	sel = strings.ToLower(sel)
	return IntentTemplateSelector(sel)
}
