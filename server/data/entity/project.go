package entity

import "github.com/bitwormhole/wpm/server/data/dxo"

// Project ...
type Project struct {
	ID dxo.ProjectID `gorm:"primaryKey"`
	Base

	Name           string
	PathInWorktree string
	Description    string
	IsFile         bool
	IsDir          bool
	// ProjectTypeName string
	ConfigFileName  string
	OwnerRepository dxo.LocalRepositoryID

	// FullPath        string `gorm:"index:,unique"`  [已废弃] 用 Path 代替
	ProjectDir string

	Path     string         // this.Path == Location.Path
	Location dxo.LocationID `gorm:"index:,unique"`
	Class    dxo.LocationClass

	Type dxo.ProjectTypeURN
}

// ListPathFields 用于 FindByPath ...
func (Project) ListPathFields() []string {
	return []string{"path", "project_dir"}
}
