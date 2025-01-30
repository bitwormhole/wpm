package locations

import (
	"github.com/bitwormhole/wpm/app/data/dxo"
	"github.com/bitwormhole/wpm/app/data/entity"
	"gorm.io/gorm"
)

// LocationDAO ...
type LocationDAO interface {
	Find(db *gorm.DB, id dxo.LocationID) (*entity.Location, error)

	ListWithCategory(db *gorm.DB, category dxo.Category, q *Query) ([]*entity.Location, error)

	Insert(db *gorm.DB, item *entity.Location) (*entity.Location, error)

	Update(db *gorm.DB, id dxo.LocationID, fn func(*entity.Location)) (*entity.Location, error)

	Remove(db *gorm.DB, id dxo.LocationID) error
}
