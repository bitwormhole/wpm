package dao

import (
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/data/entity"
)

// SoftwarePackageDAO ...
type SoftwarePackageDAO interface {
	Find(id dxo.SoftwarePackageID) (*entity.SoftwarePackage, error)

	ListAll() ([]*entity.SoftwarePackage, error)

	Insert(o *entity.SoftwarePackage) (*entity.SoftwarePackage, error)

	Update(id dxo.SoftwarePackageID, o *entity.SoftwarePackage) (*entity.SoftwarePackage, error)

	Remove(id dxo.SoftwarePackageID) error
}
