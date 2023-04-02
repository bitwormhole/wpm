package dxo

import "strings"

// URN 以字符串形式表示一个类似 'urn:type:id' 格式的名称
type URN string

// ExecutableURN like 'urn:executable:{{namespace.name}}#{{simpleName}}'
type ExecutableURN URN

// NamespaceURN like 'urn:namespace:{{name}}'
type NamespaceURN URN

// ProjectTypeURN like 'urn:projecttype:{{name}}'
type ProjectTypeURN URN

////////////////////////////////////////////////////////////////////////////////

func (name URN) String() string {
	return string(name)
}

func (name ExecutableURN) String() string {
	return string(name)
}

func (name NamespaceURN) String() string {
	return string(name)
}

////////////////////////////////////////////////////////////////////////////////

// NewURN ...
func NewURN(typeName string, id string) URN {
	p1 := strings.TrimSpace(typeName)
	p2 := strings.TrimSpace(id)
	str := "urn:" + p1 + ":" + p2
	return URN(strings.ToLower(str))
}

// NewNamespaceURN ...
func NewNamespaceURN(name string) NamespaceURN {
	t := "namespace"
	urn := NewURN(t, name)
	return NamespaceURN(urn)
}

// NewExecutableURN ...
func NewExecutableURN(name string) ExecutableURN {
	t := "executable"
	urn := NewURN(t, name)
	return ExecutableURN(urn)
}

// NewProjectTypeURN ...
func NewProjectTypeURN(name string) ProjectTypeURN {
	t := "projecttype"
	urn := NewURN(t, name)
	return ProjectTypeURN(urn)
}