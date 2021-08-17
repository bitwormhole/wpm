package dao

import "github.com/bitwormhole/wpm/data/entity"

type Repository interface {
	Get(name string) (*entity.Repository, error)
	All() []*entity.Repository
}

type RepositoryImpl struct {
}

func (inst *RepositoryImpl) _Impl() Repository {
	return inst
}

func (inst *RepositoryImpl) Get(name string) (*entity.Repository, error) {
	return nil, nil
}

func (inst *RepositoryImpl) All() []*entity.Repository {
	return nil
}
