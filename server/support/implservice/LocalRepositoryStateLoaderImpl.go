package implservice

import (
	"context"
	"fmt"

	"bitwormhole.com/starter/afs"
	"github.com/bitwormhole/gitlib/git/store"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/data/dao"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/data/entity"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// LocalRepositoryStateLoaderImpl ...
type LocalRepositoryStateLoaderImpl struct {
	markup.Component `id:"LocalRepositoryStateLoader"`

	LocalRepoService service.LocalRepositoryService `inject:"#LocalRepositoryService"`
	Dao              dao.LocalRepositoryDAO         `inject:"#LocalRepositoryDAO"`
	GitLibAgent      store.LibAgent                 `inject:"#git-lib-agent"`
}

func (inst *LocalRepositoryStateLoaderImpl) _Impl() service.LocalRepositoryStateLoader {
	return inst
}

// LoadState ...
func (inst *LocalRepositoryStateLoaderImpl) LoadState(ctx context.Context, repo0 *dto.LocalRepository) error {

	repo := repo0
	isInDB := false

	if ctx == nil || repo == nil {
		return fmt.Errorf("param is nil")
	}

	if inst.needForEntity(repo) {
		repo2, err := inst.loadEntityAsDTO(ctx, repo)
		if err == nil {
			repo = repo2
			isInDB = true
		}
	} else {
		isInDB = true
	}

	err := inst.loadStateFromFileSystem(ctx, repo)
	if err != nil {
		return err
	}

	if !isInDB {
		repo.State = dxo.FileStateUntracked
	}

	*repo0 = *repo
	repo0.Ready = (repo.State == dxo.FileStateReady)
	return nil
}

func (inst *LocalRepositoryStateLoaderImpl) needForEntity(repo *dto.LocalRepository) bool {
	id := repo.ID
	uuid := repo.UUID
	if id > 0 && uuid != "" {
		return false
	}
	return true
}

func (inst *LocalRepositoryStateLoaderImpl) findEntityByID(id dxo.LocalRepositoryID) (*entity.LocalRepository, error) {
	dao1 := inst.Dao
	ent, err := dao1.Find(id)
	if err == nil && ent != nil {
		return ent, nil
	}
	return nil, fmt.Errorf("repo not found")
}

func (inst *LocalRepositoryStateLoaderImpl) findEntityByUUID(uuid dxo.UUID) (*entity.LocalRepository, error) {
	dao1 := inst.Dao
	ent, err := dao1.FindByUUID(uuid)
	if err == nil && ent != nil {
		return ent, nil
	}
	return nil, fmt.Errorf("repo not found")
}

func (inst *LocalRepositoryStateLoaderImpl) findEntityByPath(path string) (*entity.LocalRepository, error) {
	dao1 := inst.Dao
	ent, err := dao1.FindByPath(path)
	if err == nil && ent != nil {
		return ent, nil
	}
	return nil, fmt.Errorf("repo not found")
}

func (inst *LocalRepositoryStateLoaderImpl) findEntityByWorkingDir(path string) (*entity.LocalRepository, error) {
	dao1 := inst.Dao
	ent, err := dao1.FindByWorkingDir(path)
	if err == nil && ent != nil {
		return ent, nil
	}
	return nil, fmt.Errorf("repo not found")
}

func (inst *LocalRepositoryStateLoaderImpl) findEntityByDotGit(path string) (*entity.LocalRepository, error) {
	dao1 := inst.Dao
	ent, err := dao1.FindByDotGit(path)
	if err == nil && ent != nil {
		return ent, nil
	}
	return nil, fmt.Errorf("repo not found")
}

func (inst *LocalRepositoryStateLoaderImpl) loadEntityAsDTO(ctx context.Context, repo *dto.LocalRepository) (*dto.LocalRepository, error) {
	ent, err := inst.loadEntity(ctx, repo)
	if err != nil {
		return nil, err
	}
	return inst.LocalRepoService.ConvertEntityToDto(ctx, ent)
}

func (inst *LocalRepositoryStateLoaderImpl) loadEntity(ctx context.Context, repo *dto.LocalRepository) (*entity.LocalRepository, error) {

	id := repo.ID
	uuid := repo.UUID
	path := repo.Path

	if id > 0 {
		ent, err := inst.findEntityByID(id)
		if err == nil {
			return ent, nil
		}
	}

	if uuid != "" {
		ent, err := inst.findEntityByUUID(uuid)
		if err == nil {
			return ent, nil
		}
	}

	errNoRepoInfo := fmt.Errorf("no repository info in database")
	if path == "" {
		return nil, errNoRepoInfo
	}

	ent, err := inst.findEntityByPath(path)
	if err == nil {
		return ent, nil
	}
	ent, err = inst.findEntityByWorkingDir(path)
	if err == nil {
		return ent, nil
	}
	ent, err = inst.findEntityByDotGit(path)
	if err == nil {
		return ent, nil
	}

	return nil, errNoRepoInfo
}

func (inst *LocalRepositoryStateLoaderImpl) loadStateFromFileSystem(ctx context.Context, repo *dto.LocalRepository) error {

	lib, err := inst.GitLibAgent.GetLib()
	if err != nil {
		return err
	}

	pwd := lib.FS().NewPath(repo.Path)
	layout, err := lib.RepositoryLocator().Locate(pwd)
	if err != nil {
		repo.State = dxo.FileStateOffline
		return nil
	}

	rPath, rName, err := inst.getRepoPathAndName(layout)
	if err != nil {
		repo.State = dxo.FileStateOffline
		return nil
	}

	wantPath := repo.Path
	havePath := rPath.GetPath()

	if wantPath == havePath {
		repo.State = dxo.FileStateReady
	} else {
		repo.State = dxo.FileStateMoved
	}
	repo.Name = rName

	return nil
}

func (inst *LocalRepositoryStateLoaderImpl) getRepoPathAndName(layout store.RepositoryLayout) (afs.Path, string, error) {

	path := layout.DotGit()
	if path != nil {
		name := path.GetParent().GetName()
		return path, name, nil
	}

	path = layout.Repository()
	if path != nil {
		name := path.GetName()
		return path, name, nil
	}

	return nil, "", fmt.Errorf("repository not found")
}
