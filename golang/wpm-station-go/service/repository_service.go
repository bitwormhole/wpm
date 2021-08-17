package service

import (
	"github.com/bitwormhole/wpm/convert/entity2dto"
	"github.com/bitwormhole/wpm/data/dao"
	"github.com/bitwormhole/wpm/web/dto"
)

type RepositoryService interface {
	ListAll() []*dto.Repository
	FindByName(name string) (*dto.Repository, error)
	Update(name string, repo *dto.Repository) (*dto.Repository, error)
}

////////////////////////////////////////////////////////////////////////////////

type RepositoryServiceImpl struct {
	Store *dao.Store
}

func (inst *RepositoryServiceImpl) _Impl() RepositoryService {
	return inst
}

func (inst *RepositoryServiceImpl) ListAll() []*dto.Repository {
	all := inst.Store.RepositoryDAO.All()
	return entity2dto.ConvertRepositoryList(all)
}

func (inst *RepositoryServiceImpl) FindByName(name string) (*dto.Repository, error) {

	return nil, nil
}

func (inst *RepositoryServiceImpl) Update(name string, repo *dto.Repository) (*dto.Repository, error) {

	return nil, nil
}
