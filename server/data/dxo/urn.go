package dxo

import "strings"

// URN 以字符串形式表示一个类似 'urn:type:id' 格式的名称
type URN string

// ExecutableURN like 'urn:executable:{{namespace.name}}#{{simpleName}}'
type ExecutableURN URN

// NamespaceURN like 'urn:namespace:{{name}}'
type NamespaceURN URN

// ContentTypeURN like 'urn:projecttype:{{name}}'
type ContentTypeURN URN

// SoftwarePackageURN like 'urn:package:{{name}}'
type SoftwarePackageURN URN

// SoftwareSetURN like 'urn:packageset:{{name}}'
type SoftwareSetURN URN

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

func (name SoftwarePackageURN) String() string {
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

// NewContentTypeURN ...
func NewContentTypeURN(name string) ContentTypeURN {
	t := "contenttype"
	urn := NewURN(t, name)
	return ContentTypeURN(urn)
}

// NewSoftwarePackageURN ...
func NewSoftwarePackageURN(name string) SoftwarePackageURN {
	t := "package"
	urn := NewURN(t, name)
	return SoftwarePackageURN(urn)
}
