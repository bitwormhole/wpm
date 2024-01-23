package settings

import (
	"github.com/bitwormhole/wpm/common/objects/dxo"
	"github.com/bitwormhole/wpm/server/data/entity"
	"gorm.io/gorm"
)

// DAO 表示该类型的 Data Access Object 接口
type DAO interface {
	Insert(db *gorm.DB, o *entity.Setting) (*entity.Setting, error)
	Update(db *gorm.DB, id dxo.SettingID, fn func(*entity.Setting)) (*entity.Setting, error)

	Remove(db *gorm.DB, id dxo.SettingID) error

	Find(db *gorm.DB, id dxo.SettingID) (*entity.Setting, error)
	List(db *gorm.DB, q *Query) ([]*entity.Setting, error)
}
