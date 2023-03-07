package entity

import "github.com/bitwormhole/wpm/server/data/dxo"

// ProjectType ...
type ProjectType struct {
	ID dxo.ProjectTypeID `gorm:"primaryKey"`
	Base

	Name        string
	Type        string
	Label       string
	Description string

	AsFile bool
	AsDir  bool
}
