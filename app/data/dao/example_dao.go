package dao

import (
	"github.com/bitwormhole/wpm/app/data/dxo"
	"github.com/bitwormhole/wpm/app/data/entity"
	"gorm.io/gorm"
)

// ExampleDAO ...
type ExampleDAO interface {
	Find(db *gorm.DB, id dxo.ExampleID) (*entity.Example, error)
}
