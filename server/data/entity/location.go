package entity

import "github.com/bitwormhole/wpm/common/objects/dxo"

// Location ...
type Location struct {
	ID dxo.LocationID `gorm:"primaryKey"`
	Base

	RawPath string
	Path    string `gorm:"index:,unique"` // the normalized full-path

	Target dxo.URN
	Class  dxo.LocationClass
	AsFile bool // 表示该路径可以是一个普通文件
	AsDir  bool // 表示该路径可以是一个目录

	// [已废弃] 以下字段仅当 Class==LocationProject 时有用
	// TypeID   dxo.ProjectTypeID
	// TypeKey  dxo.ProjectTypeKey
	// TypeName dxo.ProjectTypeName
}

// ListPathFields ...
func (Location) ListPathFields() []string {
	return []string{"path"}
}

// TableName 。。。
func (Location) TableName() string {
	return getTableName("locations")
}
