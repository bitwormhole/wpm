package namespaces

import (
	"github.com/bitwormhole/wpm/common/objects/dxo"
	"github.com/bitwormhole/wpm/server/data/entity"
	"gorm.io/gorm"
)

// DAO 表示该类型的 Data Access Object 接口
type DAO interface {
	Insert(db *gorm.DB, o *entity.Namespace) (*entity.Namespace, error)
	Update(db *gorm.DB, id dxo.NamespaceID, fn func(*entity.Namespace)) (*entity.Namespace, error)

	Remove(db *gorm.DB, id dxo.NamespaceID) error

	Find(db *gorm.DB, id dxo.NamespaceID) (*entity.Namespace, error)
	List(db *gorm.DB, q *Query) ([]*entity.Namespace, error)
}
