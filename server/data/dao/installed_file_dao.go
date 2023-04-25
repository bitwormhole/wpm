package dao

import (
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/data/entity"
)

// InstalledFileDAO ...
type InstalledFileDAO interface {
	Find(id dxo.InstalledFileID) (*entity.InstalledFile, error)

	ListAll() ([]*entity.InstalledFile, error)

	Insert(o *entity.InstalledFile) (*entity.InstalledFile, error)

	Update(id dxo.InstalledFileID, o *entity.InstalledFile) (*entity.InstalledFile, error)

	Remove(id dxo.InstalledFileID) error
}
