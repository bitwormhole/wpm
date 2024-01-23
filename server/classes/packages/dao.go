package packages

import (
	"github.com/bitwormhole/wpm/common/objects/dxo"
	"github.com/bitwormhole/wpm/server/data/entity"
	"gorm.io/gorm"
)

// DAO 表示该类型的 Data Access Object 接口
type DAO interface {
	Insert(db *gorm.DB, o *entity.SoftwarePackage) (*entity.SoftwarePackage, error)
	Update(db *gorm.DB, id dxo.SoftwarePackageID, fn func(*entity.SoftwarePackage)) (*entity.SoftwarePackage, error)

	Remove(db *gorm.DB, id dxo.SoftwarePackageID) error

	Find(db *gorm.DB, id dxo.SoftwarePackageID) (*entity.SoftwarePackage, error)
	List(db *gorm.DB, q *Query) ([]*entity.SoftwarePackage, error)
}
