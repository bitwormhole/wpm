package repositories

import (
	"context"

	"github.com/bitwormhole/gitlib/git/store"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// MainRepositoryServiceImpl ...
type MainRepositoryServiceImpl struct {
	markup.Component `id:"MainRepositoryService"`

	// MainRepoPath string         `inject:"${wpm.system-main-repository.path}"`
	AppDataService service.AppDataService `inject:"#AppDataService"`
	GitLA          store.LibAgent         `inject:"#git-lib-agent"`

	// cachedInfo *dto.SystemMainRepository
	cachedRepo store.Repository
}

func (inst *MainRepositoryServiceImpl) _Impl() service.MainRepositoryService {
	return inst
}

// GetInfo ...
func (inst *MainRepositoryServiceImpl) GetInfo(ctx context.Context) (*dto.MainRepository, error) {
	path := inst.AppDataService.GetMainRepositoryPath()
	info := &dto.MainRepository{
		Path: path,
	}
	return info, nil
}

// GetRepository ...
func (inst *MainRepositoryServiceImpl) GetRepository(ctx context.Context) (store.Repository, error) {
	repo := inst.cachedRepo
	if repo != nil {
		return repo, nil
	}
	repo, err := inst.loadRepository(ctx)
	if err != nil {
		return nil, err
	}
	inst.cachedRepo = repo
	return repo, nil
}

func (inst *MainRepositoryServiceImpl) loadRepository(ctx context.Context) (store.Repository, error) {
	info, err := inst.GetInfo(ctx)
	if err != nil {
		return nil, err
	}
	lib, err := inst.GitLA.GetLib()
	if err != nil {
		return nil, err
	}
	path := lib.FS().NewPath(info.Path)
	return lib.RepositoryLoader().LoadWithPath(path)
}
