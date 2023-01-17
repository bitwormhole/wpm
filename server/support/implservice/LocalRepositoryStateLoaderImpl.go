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

	Dao         dao.LocalRepositoryDAO `inject:"#LocalRepositoryDAO"`
	GitLibAgent store.LibAgent         `inject:"#git-lib-agent"`
}

func (inst *LocalRepositoryStateLoaderImpl) _Impl() service.LocalRepositoryStateLoader {
	return inst
}

// LoadState ...
func (inst *LocalRepositoryStateLoaderImpl) LoadState(ctx context.Context, repo *dto.LocalRepository) error {

	if ctx == nil || repo == nil {
		return fmt.Errorf("param is nil")
	}

	err := inst.loadStateFromFileSystem(ctx, repo)
	if err != nil {
		return err
	}

	if repo.State != dxo.FileStateReady {
		return nil
	}

	err = inst.loadStateFromDatabase(ctx, repo)
	if err != nil {
		return err
	}

	repo.Ready = (repo.State == dxo.FileStateReady)
	return nil
}

func (inst *LocalRepositoryStateLoaderImpl) findEntity(repo *dto.LocalRepository) (*entity.LocalRepository, error) {

	dao1 := inst.Dao
	id := repo.ID
	path := repo.Path
	// name := repo.Name

	ent, err := dao1.Find(id)
	if err == nil && ent != nil {
		return ent, nil
	}

	ent, err = dao1.FindByPath(path)
	if err == nil && ent != nil {
		return ent, nil
	}

	// ent, err = dao1.FindByName(name)
	// if err == nil && ent != nil {
	// 	return ent, nil
	// }

	return nil, fmt.Errorf("repo not found")
}

func (inst *LocalRepositoryStateLoaderImpl) loadStateFromDatabase(ctx context.Context, repo *dto.LocalRepository) error {

	ent, err := inst.findEntity(repo)
	if err != nil {
		repo.State = dxo.FileStateUntracked
		return nil
	}

	if ent.Path == repo.Path {
		repo.State = dxo.FileStateReady
	} else {
		repo.State = dxo.FileStateMoved
	}

	return nil
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
