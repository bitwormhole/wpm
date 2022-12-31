package dao

import (
	"github.com/bitwormhole/wpm/app/data/dxo"
	"github.com/bitwormhole/wpm/app/data/entity"
)

// RepositoryDAO ...
type RepositoryDAO interface {
	Find(id dxo.RepositoryID) (*entity.Repository, error)
	FindByName(name string) (*entity.Repository, error)
	FindByPath(path string) (*entity.Repository, error)

	ListAll() ([]*entity.Repository, error)

	Insert(o *entity.Repository) (*entity.Repository, error)
	Update(id dxo.RepositoryID, o *entity.Repository) (*entity.Repository, error)
	Remove(id dxo.RepositoryID) error
}
