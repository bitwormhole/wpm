package entity

import "github.com/bitwormhole/wpm/server/data/dxo"

// Project ...
type Project struct {
	ID dxo.ProjectID `gorm:"primaryKey"`
	Base
	URN dxo.URN

	Name           string
	PathInWorktree string
	Description    string
	IsFile         bool
	IsDir          bool
	// ProjectTypeName string

	ConfigFileName  string
	OwnerRepository dxo.LocalRepositoryID

	// FullPath        string `gorm:"index:,unique"`  [已废弃] 用 Path 代替
	ProjectDir  string
	RegularPath string `gorm:"index:,unique"`

	Path string // = RegularPath

	// Location dxo.LocationID `gorm:"index:,unique"`
	// Class    dxo.LocationClass

	Type dxo.ContentTypeName
}

// ListPathFields 用于 FindByPath ...
func (Project) ListPathFields() []string {
	return []string{"path", "project_dir"}
}

// TableName 。。。
func (Project) TableName() string {
	return getTableName("projects")
}
