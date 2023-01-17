package implservice

import (
	"context"

	"github.com/bitwormhole/gitlib/git/store"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// SystemMainRepositoryServiceImpl ...
type SystemMainRepositoryServiceImpl struct {
	markup.Component `id:"SystemMainRepositoryService"`

	SystemMainRepoPath string         `inject:"${wpm.system-main-repository.path}"`
	GitLA              store.LibAgent `inject:"#git-lib-agent"`

	// cachedInfo *dto.SystemMainRepository
	cachedRepo store.Repository
}

func (inst *SystemMainRepositoryServiceImpl) _Impl() service.SystemMainRepositoryService {
	return inst
}

// GetInfo ...
func (inst *SystemMainRepositoryServiceImpl) GetInfo(ctx context.Context) (*dto.SystemMainRepository, error) {
	path := inst.SystemMainRepoPath
	info := &dto.SystemMainRepository{
		Path: path,
	}
	return info, nil
}

// GetRepository ...
func (inst *SystemMainRepositoryServiceImpl) GetRepository(ctx context.Context) (store.Repository, error) {
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

func (inst *SystemMainRepositoryServiceImpl) loadRepository(ctx context.Context) (store.Repository, error) {
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
