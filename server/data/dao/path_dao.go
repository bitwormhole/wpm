package dao

import (
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/data/entity"
)

// PathDAO ...
type PathDAO interface {
	Find(id dxo.PathID) (*entity.Path, error)

	FindByName(fullname string) (*entity.Path, error)

	ListAll() ([]*entity.Path, error)

	Insert(o *entity.Path) (*entity.Path, error)

	Update(id dxo.PathID, o *entity.Path) (*entity.Path, error)

	Remove(id dxo.PathID) error
}
