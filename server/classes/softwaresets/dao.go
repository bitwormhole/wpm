package softwaresets

import (
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/data/entity"
	"gorm.io/gorm"
)

// DAO 表示该类型的 Data Access Object 接口
type DAO interface {
	Insert(db *gorm.DB, o *entity.SoftwareSet) (*entity.SoftwareSet, error)
	Update(db *gorm.DB, id dxo.SoftwareSetID, fn func(*entity.SoftwareSet)) (*entity.SoftwareSet, error)

	Remove(db *gorm.DB, id dxo.SoftwareSetID) error

	Find(db *gorm.DB, id dxo.SoftwareSetID) (*entity.SoftwareSet, error)
	List(db *gorm.DB, q *Query) ([]*entity.SoftwareSet, error)
}
