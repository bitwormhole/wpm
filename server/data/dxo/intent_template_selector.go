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
	Method string // [open|insert|update|edit|...]
	Target string // [repository|worktree|project||]
	Type   string // [project-type|content-type]
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
	target := inst.normalizeName(inst.Target)
	as := inst.normalizeName(inst.Type)
	with := inst.normalizeName(inst.With)
	builder := strings.Builder{}

	if method == "" {
		method = "open"
	}

	const (
		keyAS   = "as"
		keyWith = "with"
	)
	keyTarget := method
	keys := []string{keyTarget, keyAS, keyWith}

	table := make(map[string]string)
	table[keyTarget] = target
	table[keyAS] = as
	table[keyWith] = with

	for _, key := range keys {
		value := table[key]
		if value == "" {
			value = "*"
		}
		builder.WriteString(".")
		builder.WriteString(key)
		builder.WriteString("(")
		builder.WriteString(value)
		builder.WriteString(")")
	}

	sel := builder.String()
	sel = strings.ToLower(sel)
	return IntentTemplateSelector(sel)
}
