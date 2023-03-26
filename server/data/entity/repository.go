package entity

import "github.com/bitwormhole/wpm/server/data/dxo"

// RemoteRepository ...
type RemoteRepository struct {
	ID dxo.RemoteRepositoryID `gorm:"primaryKey"`
	Base

	Name        string
	URL         string
	DisplayName string
}

// LocalRepository ...
type LocalRepository struct {
	ID dxo.LocalRepositoryID `gorm:"primaryKey"`
	Base

	// LocalRepositoryBase

	Name        string
	DisplayName string
	Description string
	Bare        bool

	ConfigFile     string
	RepositoryPath string // the parent of ConfigFile
	DotGitPath     string // can be empty
	WorkingPath    string // can be empty

	Path     string         // this.Path == Location.Path == ConfigFile
	Location dxo.LocationID `gorm:"index:,unique"`
	Class    dxo.LocationClass
}

// ListPathFields ...
func (LocalRepository) ListPathFields() []string {
	return []string{"path", "config_file", "repository_path", "dot_git_path", "working_path"}
}
