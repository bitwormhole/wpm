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

	Path       string // this.Path == ConfigFile.parentDir
	ConfigFile string // the path of '.git/config' file

	Location dxo.LocationID `gorm:"unique"`

	// Bare        bool
	// RawPath     string
	// RepositoryPath string // the parent of ConfigFile
	// DotGitPath     string // can be empty
	// WorkingPath    string // can be empty
	// Location dxo.LocationID
	// Class    dxo.LocationClass
}

////////////////////////////////////////////////////////////////////////////////

// TableName 。。。
func (RemoteRepository) TableName() string {
	return getTableName("remote_repositories")
}

////////////////////////////////////////////////////////////////////////////////

// TableName 。。。
func (LocalRepository) TableName() string {
	return getTableName("local_repositories")
}

// ComputeRegularPath 。。。
func (inst *LocalRepository) ComputeRegularPath() dxo.RegularPath {
	path := inst.Path
	return dxo.NewRegularPath(path)
}

// ListPathFields ...
func (LocalRepository) ListPathFields() []string {
	return []string{"path", "config_file", "repository_path", "dot_git_path", "working_path", "regular_path"}
}

////////////////////////////////////////////////////////////////////////////////
