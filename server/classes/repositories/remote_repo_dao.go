package repositories

import (
	"github.com/bitwormhole/wpm/common/objects/dxo"
	"github.com/bitwormhole/wpm/server/data/entity"
	"gorm.io/gorm"
)

// RemoteRepositoryDAO 表示该类型的 Data Access Object 接口
type RemoteRepositoryDAO interface {
	Insert(db *gorm.DB, o *entity.RemoteRepository) (*entity.RemoteRepository, error)
	Update(db *gorm.DB, id dxo.RemoteRepositoryID, fn func(*entity.RemoteRepository)) (*entity.RemoteRepository, error)

	Remove(db *gorm.DB, id dxo.RemoteRepositoryID) error

	Find(db *gorm.DB, id dxo.RemoteRepositoryID) (*entity.RemoteRepository, error)
	List(db *gorm.DB, q *Query) ([]*entity.RemoteRepository, error)
}
