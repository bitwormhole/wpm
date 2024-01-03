package projects

import (
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/data/entity"
	"gorm.io/gorm"
)

// DAO 表示该类型的 Data Access Object 接口
type DAO interface {
	Insert(db *gorm.DB, o *entity.Project) (*entity.Project, error)
	Update(db *gorm.DB, id dxo.ProjectID, fn func(*entity.Project)) (*entity.Project, error)

	Remove(db *gorm.DB, id dxo.ProjectID) error

	Find(db *gorm.DB, id dxo.ProjectID) (*entity.Project, error)
	List(db *gorm.DB, q *Query) ([]*entity.Project, error)
}
