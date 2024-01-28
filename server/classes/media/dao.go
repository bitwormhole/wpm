package media

import (
	"github.com/bitwormhole/wpm/common/objects/dxo"
	"github.com/bitwormhole/wpm/server/data/entity"
	"gorm.io/gorm"
)

// DAO 表示该类型的 Data Access Object 接口
type DAO interface {
	Insert(db *gorm.DB, o *entity.Media) (*entity.Media, error)

	Update(db *gorm.DB, id dxo.MediaID, fn func(*entity.Media)) (*entity.Media, error)

	Remove(db *gorm.DB, id dxo.MediaID) error

	Find(db *gorm.DB, id dxo.MediaID) (*entity.Media, error)

	List(db *gorm.DB, q *Query) ([]*entity.Media, error)
}
