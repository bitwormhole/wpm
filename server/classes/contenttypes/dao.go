package contenttypes

import (
	"github.com/bitwormhole/wpm/common/objects/dxo"
	"github.com/bitwormhole/wpm/server/data/entity"
	"gorm.io/gorm"
)

// DAO 表示该类型的 Data Access Object 接口
type DAO interface {
	Insert(db *gorm.DB, o *entity.ContentType) (*entity.ContentType, error)
	Update(db *gorm.DB, id dxo.ContentTypeID, fn func(*entity.ContentType)) (*entity.ContentType, error)

	Remove(db *gorm.DB, id dxo.ContentTypeID) error

	Find(db *gorm.DB, id dxo.ContentTypeID) (*entity.ContentType, error)
	List(db *gorm.DB, q *Query) ([]*entity.ContentType, error)
}
