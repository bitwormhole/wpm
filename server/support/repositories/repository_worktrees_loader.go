package repositories

import (
	"fmt"

	"github.com/bitwormhole/gitlib/git/store"
	"github.com/bitwormhole/starter/vlog"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/web/dto"
)

type myWorktreesLoader struct {
	lib store.Lib
}

func (inst *myWorktreesLoader) Load(path string, owner dxo.LocalRepositoryID) ([]*dto.Worktree, error) {

	at := inst.lib.FS().NewPath(path)
	repo, err := inst.lib.RepositoryLoader().LoadWithPath(at)
	if err != nil {
		return nil, err
	}

	src := repo.Worktrees().List()
	dst := make([]*dto.Worktree, 0)

	wt0, err := inst.loadDefaultWorktree(repo)
	if err == nil && wt0 != nil {
		wt0.OwnerRepository = owner
		dst = append(dst, wt0)
	}

	for _, item1 := range src {
		name := item1.Name()
		item2 := &dto.Worktree{
			Name:            name,
			OwnerRepository: owner,
		}
		err := inst.loadItem(item1, item2)
		if err == nil {
			dst = append(dst, item2)
		} else {
			vlog.Warn(err)
		}
	}

	return dst, nil
}

func (inst *myWorktreesLoader) loadDefaultWorktree(repo store.Repository) (*dto.Worktree, error) {

	errBare := fmt.Errorf("repository is bare")
	layout := repo.Layout()
	if layout.IsBare() {
		return nil, errBare
	}
	dg := layout.DotGit()
	if dg == nil {
		return nil, errBare
	}
	if !dg.IsDirectory() {
		return nil, errBare
	}
	wkdir := dg.GetParent()

	dst := &dto.Worktree{}
	dst.DotGitPath = dg.GetPath()
	dst.WorkingDir = wkdir.GetPath()
	dst.Path = wkdir.GetPath()
	dst.IsDefault = true

	return dst, nil
}

func (inst *myWorktreesLoader) loadItem(src store.Worktree, dst *dto.Worktree) error {

	if src == nil {
		return fmt.Errorf("the worktree is nil")
	}
	if !src.Exists() {
		return fmt.Errorf("the worktree is not exists")
	}

	workspace := src.Workspace()
	dotgit := workspace.DotGit().GetPath()
	wkdir := workspace.WorkingDirectory().GetPath()

	dst.DotGitPath = dotgit
	dst.WorkingDir = wkdir
	dst.Path = wkdir
	dst.IsSubmodule = false
	dst.IsDefault = false

	return nil
}
