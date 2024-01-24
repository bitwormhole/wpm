package entity

import "github.com/bitwormhole/wpm/common/objects/dxo"

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
	URN dxo.URN

	// LocalRepositoryBase

	Name        string
	DisplayName string
	Description string
	Bare        bool

	Path        string // this.Path == RegularPath.Path
	RawPath     string
	RegularPath string `gorm:"index:,unique"` // = ConfigFile/..

	// ConfigFile     string
	// RepositoryPath string // the parent of ConfigFile
	// DotGitPath     string // can be empty
	// WorkingPath    string // can be empty
	// Location dxo.LocationID
	// Class    dxo.LocationClass
}

// ListPathFields ...
func (LocalRepository) ListPathFields() []string {
	return []string{"path", "config_file", "repository_path", "dot_git_path", "working_path", "regular_path"}
}

// TableName 。。。
func (RemoteRepository) TableName() string {
	return getTableName("remote_repositories")
}

// TableName 。。。
func (LocalRepository) TableName() string {
	return getTableName("local_repositories")
}
