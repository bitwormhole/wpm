package profiles

import (
	"github.com/bitwormhole/wpm/common/objects/dxo"
	"github.com/bitwormhole/wpm/server/data/entity"
	"gorm.io/gorm"
)

// DAO 表示该类型的 Data Access Object 接口
type DAO interface {
	Insert(db *gorm.DB, o *entity.Profile) (*entity.Profile, error)
	Update(db *gorm.DB, id dxo.ProfileID, fn func(*entity.Profile)) (*entity.Profile, error)

	Remove(db *gorm.DB, id dxo.ProfileID) error

	Find(db *gorm.DB, id dxo.ProfileID) (*entity.Profile, error)
	List(db *gorm.DB, q *Query) ([]*entity.Profile, error)
}
