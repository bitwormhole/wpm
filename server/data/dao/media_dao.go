package dao

import (
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/data/entity"
)

// MediaDAO ...
type MediaDAO interface {
	Find(id dxo.MediaID) (*entity.Media, error)

	FindByPath(path string) (*entity.Media, error)

	ListByIDs(ids []dxo.MediaID) ([]*entity.Media, error)

	ListAll() ([]*entity.Media, error)

	Insert(o *entity.Media) (*entity.Media, error)

	Update(id dxo.MediaID, o *entity.Media) (*entity.Media, error)

	Remove(id dxo.MediaID) error
}
