package repositories

import (
	"github.com/bitwormhole/wpm/common/objects/dxo"
	"github.com/bitwormhole/wpm/server/data/entity"
	"gorm.io/gorm"
)

// LocalRepositoryDAO 表示该类型的 Data Access Object 接口
type LocalRepositoryDAO interface {
	Insert(db *gorm.DB, o *entity.LocalRepository) (*entity.LocalRepository, error)
	Update(db *gorm.DB, id dxo.LocalRepositoryID, fn func(*entity.LocalRepository)) (*entity.LocalRepository, error)

	Remove(db *gorm.DB, id dxo.LocalRepositoryID) error

	Find(db *gorm.DB, id dxo.LocalRepositoryID) (*entity.LocalRepository, error)
	List(db *gorm.DB, q *Query) ([]*entity.LocalRepository, error)
}
