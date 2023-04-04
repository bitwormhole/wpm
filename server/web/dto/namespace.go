package dto

import "github.com/bitwormhole/wpm/server/data/dxo"

// Namespace 表示一个命令模板
type Namespace struct {
	ID  dxo.NamespaceID  `json:"id"`
	URN dxo.NamespaceURN `json:"urn"`

	Base

	URL  string `json:"url"` // 指向命名空间的包列表
	Name string `json:"name"`
	OS   string `json:"os"`   // 操作系统
	Arch string `json:"arch"` // 处理器架构

	Packages []*SoftwarePackage `json:"packages"`
}
