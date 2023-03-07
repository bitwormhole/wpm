package dao

import (
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/data/entity"
)

// ProjectTypeDAO ...
type ProjectTypeDAO interface {
	Find(id dxo.ProjectTypeID) (*entity.ProjectType, error)

	ListAll() ([]*entity.ProjectType, error)

	Insert(o *entity.ProjectType) (*entity.ProjectType, error)

	Update(id dxo.ProjectTypeID, o *entity.ProjectType) (*entity.ProjectType, error)

	Remove(id dxo.ProjectTypeID) error
}
