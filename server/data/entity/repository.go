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

// LocalRepositoryBase ...
type LocalRepositoryBase struct {
	Name        string
	DisplayName string
	Description string

	Path           string
	ConfigFile     string
	DotGitPath     string
	RepositoryPath string
	WorkingPath    string
}

// LocalRepository ...
type LocalRepository struct {
	ID dxo.LocalRepositoryID `gorm:"primaryKey"`
	Base
	LocalRepositoryBase
}

// // UserMainRepository ...
// type UserMainRepository struct {
// 	ID dxo.UserMainRepositoryID `gorm:"primaryKey"`
// 	Base
// 	LocalRepositoryBase
// }
