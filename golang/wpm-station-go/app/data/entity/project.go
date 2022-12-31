package entity

import "github.com/bitwormhole/wpm/app/data/dxo"

// Project ...
type Project struct {
	ID dxo.ProjectID `gorm:"primaryKey"`
	Base

	Name            string
	Path            string
	Description     string
	ProjectType     string
	OwnerRepository dxo.RepositoryID
}
