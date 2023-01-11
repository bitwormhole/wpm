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

	Name        string
	Path        string
	ConfigFile  string
	DotGitPath  string
	WorkingPath string
	DisplayName string
	Description string
	IconURL     string
}

// MainRepository ...
type MainRepository struct {
	ID dxo.MainRepositoryID `gorm:"primaryKey"`
	Base

	Name        string
	Path        string
	ConfigFile  string
	DotGitPath  string
	WorkingPath string
	DisplayName string
}
