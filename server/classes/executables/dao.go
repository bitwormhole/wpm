package executables

import (
	"github.com/bitwormhole/wpm/common/objects/dxo"
	"github.com/bitwormhole/wpm/server/data/entity"
	"gorm.io/gorm"
)

// DAO 表示该类型的 Data Access Object 接口
type DAO interface {
	Insert(db *gorm.DB, o *entity.Executable) (*entity.Executable, error)
	Update(db *gorm.DB, id dxo.ExecutableID, fn func(*entity.Executable)) (*entity.Executable, error)

	Remove(db *gorm.DB, id dxo.ExecutableID) error

	Find(db *gorm.DB, id dxo.ExecutableID) (*entity.Executable, error)
	List(db *gorm.DB, q *Query) ([]*entity.Executable, error)
}
