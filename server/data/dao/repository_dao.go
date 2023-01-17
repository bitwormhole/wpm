package dao

import (
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/data/entity"
)

// RemoteRepositoryDAO ...
type RemoteRepositoryDAO interface {
	Find(id dxo.RemoteRepositoryID) (*entity.RemoteRepository, error)
	FindByName(name string) (*entity.RemoteRepository, error)
	FindByPath(path string) (*entity.RemoteRepository, error)

	ListAll() ([]*entity.RemoteRepository, error)

	Insert(o *entity.RemoteRepository) (*entity.RemoteRepository, error)
	Update(id dxo.RemoteRepositoryID, o *entity.RemoteRepository) (*entity.RemoteRepository, error)
	Remove(id dxo.RemoteRepositoryID) error
}

// LocalRepositoryDAO ...
type LocalRepositoryDAO interface {
	Find(id dxo.LocalRepositoryID) (*entity.LocalRepository, error)
	FindByName(name string) (*entity.LocalRepository, error)
	FindByPath(path string) (*entity.LocalRepository, error)

	ListAll() ([]*entity.LocalRepository, error)

	Insert(o *entity.LocalRepository) (*entity.LocalRepository, error)
	Update(id dxo.LocalRepositoryID, o *entity.LocalRepository) (*entity.LocalRepository, error)
	Remove(id dxo.LocalRepositoryID) error
}

// UserMainRepositoryDAO ...
type UserMainRepositoryDAO interface {
	Find(id dxo.UserMainRepositoryID) (*entity.UserMainRepository, error)

	ListAll() ([]*entity.UserMainRepository, error)

	Insert(o *entity.UserMainRepository) (*entity.UserMainRepository, error)

	Update(id dxo.UserMainRepositoryID, o *entity.UserMainRepository) (*entity.UserMainRepository, error)

	Remove(id dxo.UserMainRepositoryID) error
}
