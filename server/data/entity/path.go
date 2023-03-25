package entity

import "github.com/bitwormhole/wpm/server/data/dxo"

// Path ...
type Path struct {
	ID dxo.PathID `gorm:"primaryKey"`
	Base

	Name string `gorm:"index:,unique"` // the normalized full-path

	AsFile bool
	AsDir  bool

	TypeID  dxo.ProjectTypeID
	TypeKey dxo.ProjectTypeKey
	Type    dxo.ProjectTypeName
}
