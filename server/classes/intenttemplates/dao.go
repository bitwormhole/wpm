package intenttemplates

import (
	"github.com/bitwormhole/wpm/common/objects/dxo"
	"github.com/bitwormhole/wpm/server/data/entity"
	"gorm.io/gorm"
)

// DAO 表示该类型的 Data Access Object 接口
type DAO interface {
	Insert(db *gorm.DB, o *entity.IntentTemplate) (*entity.IntentTemplate, error)
	Update(db *gorm.DB, id dxo.IntentTemplateID, fn func(*entity.IntentTemplate)) (*entity.IntentTemplate, error)

	Remove(db *gorm.DB, id dxo.IntentTemplateID) error

	Find(db *gorm.DB, id dxo.IntentTemplateID) (*entity.IntentTemplate, error)
	List(db *gorm.DB, q *Query) ([]*entity.IntentTemplate, error)
}
