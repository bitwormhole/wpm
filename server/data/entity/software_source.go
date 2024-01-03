package entity

import "github.com/bitwormhole/wpm/server/data/dxo"

// Namespace ...
type Namespace struct {
	ID  dxo.NamespaceID  `gorm:"primaryKey"`
	URN dxo.NamespaceURN `gorm:"index:,unique"`

	Base

	Name string // like 'domain/path/to/file' format

	URL string // like 'https://domain/path/to/file' format

	Arch string
	OS   string
}

// TableName 。。。
func (Namespace) TableName() string {
	return getTableName("namespaces")
}
