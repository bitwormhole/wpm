package entity

import "github.com/bitwormhole/wpm/app/data/dxo"

// Repository ...
type Repository struct {
	ID dxo.RepositoryID `gorm:"primaryKey"`
	Base

	Name        string
	Path        string
	ConfigFile  string
	DotGitPath  string
	WorkingPath string
	DisplayName string
}
