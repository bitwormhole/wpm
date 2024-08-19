package irepositoryworktreeproject

import (
	"context"

	"github.com/bitwormhole/gitlib"
	"github.com/bitwormhole/gitlib/git/repositories"
	"github.com/bitwormhole/wpm/common/objects/dto"
	"github.com/bitwormhole/wpm/common/objects/dxo"
	"github.com/bitwormhole/wpm/common/objects/vo"
	"github.com/bitwormhole/wpm/server/classes/repositoryworktreeproject"
)

// RepositoryWorktreeProjectServiceImpl ...
type RepositoryWorktreeProjectServiceImpl struct {

	//starter:component
	_as func(repositoryworktreeproject.Service) //starter:as("#")

	GitLibAgent gitlib.Agent //starter:inject("#")

}

func (inst *RepositoryWorktreeProjectServiceImpl) _impl() repositoryworktreeproject.Service {
	return inst
}

// Find ...
func (inst *RepositoryWorktreeProjectServiceImpl) Find(ctx context.Context, want, have *vo.RepositoryWorktreeProject) error {

	lib := inst.GitLibAgent.GetLib()
	path := lib.FS().NewPath(want.Path)
	layout, err := lib.Locator().Locate(path)
	if err != nil {
		return err
	}

	have.Repository = inst.getRepository(layout)
	have.Worktree = inst.getWorktree(layout)

	return nil
}

func (inst *RepositoryWorktreeProjectServiceImpl) getWorktree(layout repositories.Layout) *dto.Worktree {

	path1 := layout.Workspace()
	path2 := layout.DotGit()
	if path1 == nil || path2 == nil {
		return nil
	}

	dst := &dto.Worktree{}
	dst.Path = path1.GetPath()
	dst.WorkingDir = path1.GetPath()
	dst.DotGitPath = path2.GetPath()
	dst.State = dxo.FileStateInit

	return dst
}

func (inst *RepositoryWorktreeProjectServiceImpl) getRepository(layout repositories.Layout) *dto.LocalRepository {

	path := layout.Repository()
	if path == nil {
		return nil
	}

	dst := &dto.LocalRepository{}
	dst.Path = path.GetPath()
	dst.State = dxo.FileStateInit

	return dst
}
