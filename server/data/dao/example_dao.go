package dao

import (
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/data/entity"
)

// ExampleDAO ...
type ExampleDAO interface {
	Find(id dxo.ExampleID) (*entity.Example, error)

	ListAll() ([]*entity.Example, error)

	Insert(o *entity.Example) (*entity.Example, error)

	Update(id dxo.ExampleID, o *entity.Example) (*entity.Example, error)

	Remove(id dxo.ExampleID) error
}
