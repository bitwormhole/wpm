package dao

import (
	"github.com/bitwormhole/wpm/app/data/dxo"
	"github.com/bitwormhole/wpm/app/data/entity"
)

// ProjectDAO ...
type ProjectDAO interface {
	Find(id dxo.ProjectID) (*entity.Project, error)
	FindByOwnerRepository(id dxo.RepositoryID) ([]*entity.Project, error)

	ListAll() ([]*entity.Project, error)

	Insert(o *entity.Project) (*entity.Project, error)
	Update(id dxo.ProjectID, o *entity.Project) (*entity.Project, error)
	Remove(id dxo.ProjectID) error
}
