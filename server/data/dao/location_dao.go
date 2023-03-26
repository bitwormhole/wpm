package dao

import (
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/data/entity"
)

// LocationDAO ...
type LocationDAO interface {
	Find(id dxo.LocationID) (*entity.Location, error)

	FindByPath(path string) (*entity.Location, error)

	ListAll() ([]*entity.Location, error)

	Insert(o *entity.Location) (*entity.Location, error)

	Update(id dxo.LocationID, o *entity.Location) (*entity.Location, error)

	Remove(id dxo.LocationID) error
}
