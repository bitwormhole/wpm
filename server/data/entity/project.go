package entity

import "github.com/bitwormhole/wpm/server/data/dxo"

// Project ...
type Project struct {
	ID dxo.ProjectID `gorm:"primaryKey"`
	Base

	Name            string
	PathInWorktree  string
	FullPath        string
	Description     string
	ProjectType     string
	OwnerRepository dxo.LocalRepositoryID
}
