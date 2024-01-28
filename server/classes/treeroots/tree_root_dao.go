package treeroots

import (
	"github.com/bitwormhole/wpm/common/objects/dxo"
	"github.com/bitwormhole/wpm/server/data/entity"
	"gorm.io/gorm"
)

// DAO for TreeRoot
type DAO interface {
	Insert(db *gorm.DB, o *entity.TreeRoot) (*entity.TreeRoot, error)

	Update(db *gorm.DB, id dxo.TreeRootID, fn func(*entity.TreeRoot)) (*entity.TreeRoot, error)

	Remove(db *gorm.DB, id dxo.TreeRootID) error

	Find(db *gorm.DB, id dxo.TreeRootID) (*entity.TreeRoot, error)

	List(db *gorm.DB, q *Query) ([]*entity.TreeRoot, error)
}
