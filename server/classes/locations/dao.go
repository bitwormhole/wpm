package locations

import (
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/data/entity"
	"gorm.io/gorm"
)

// DAO 表示该类型的 Data Access Object 接口
type DAO interface {
	Insert(db *gorm.DB, o *entity.Location) (*entity.Location, error)
	Update(db *gorm.DB, id dxo.LocationID, fn func(*entity.Location)) (*entity.Location, error)

	Remove(db *gorm.DB, id dxo.LocationID) error

	Find(db *gorm.DB, id dxo.LocationID) (*entity.Location, error)
	List(db *gorm.DB, q *Query) ([]*entity.Location, error)
}
