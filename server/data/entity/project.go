package entity

import "github.com/bitwormhole/wpm/server/data/dxo"

// Project ...
type Project struct {
	ID dxo.ProjectID `gorm:"primaryKey"`
	Base

	Name            string
	PathInWorktree  string
	FullPath        string `gorm:"index:,unique"`
	ProjectDir      string
	Description     string
	IsFile          bool
	IsDir           bool
	ProjectTypeName string
	ConfigFileName  string
	OwnerRepository dxo.LocalRepositoryID

	TypeID   dxo.ProjectTypeID
	TypeKey  dxo.ProjectTypeKey
	TypeName dxo.ProjectTypeName
}
