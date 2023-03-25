package repositoryworktreeproject

import (
	"context"
	"fmt"

	"bitwormhole.com/starter/afs"
	"github.com/bitwormhole/gitlib/git/store"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/dto"
	"github.com/bitwormhole/wpm/server/web/vo"
)

// RWPServiceImpl  ...
type RWPServiceImpl struct {
	markup.Component `id:"RepositoryWorktreeProjectService"`

	Repositories service.LocalRepositoryService `inject:"#LocalRepositoryService"`
	Worktrees    service.WorktreeService        `inject:"#WorktreeService"`
	Projects     service.ProjectService         `inject:"#ProjectService"`
	GitLib       store.LibAgent                 `inject:"#git-lib-agent"`
}

func (inst *RWPServiceImpl) _Impl() service.RepositoryWorktreeProjectService {
	return inst
}

func (inst *RWPServiceImpl) getgitlib() (store.Lib, error) {
	return inst.GitLib.GetLib()
}

func (inst *RWPServiceImpl) findRepository(ctx context.Context, layout store.RepositoryLayout) (*dto.LocalRepository, error) {

	if layout == nil {
		return nil, fmt.Errorf("no git-repository-layout")
	}

	config := layout.Config()
	dotgit := layout.DotGit()

	if config == nil {
		return nil, fmt.Errorf("no git-repository-config")
	}
	if !config.IsFile() {
		return nil, fmt.Errorf("no git-repository-config-file")
	}

	repodir := config.GetParent()
	repo := &dto.LocalRepository{}

	if dotgit != nil {
		if dotgit.Exists() {
			repo.DotGitPath = dotgit.GetPath()
			repo.WorkingPath = dotgit.GetParent().GetPath()
		}
	}

	name := repodir.GetName()
	if name == ".git" {
		name = repodir.GetParent().GetName()
	}
	repo.Name = name

	repo.ConfigFile = config.GetPath()
	repo.Path = repodir.GetPath()
	repo.RepositoryPath = repodir.GetPath()
	repo.State = dxo.FileStateUntracked

	// try find old
	// todo ...

	return repo, nil
}

func (inst *RWPServiceImpl) findWorktree(ctx context.Context, layout store.RepositoryLayout) (*dto.Worktree, error) {

	if layout == nil {
		return nil, fmt.Errorf("no git-repository-layout")
	}

	dotgit := layout.DotGit()
	wktree := &dto.Worktree{}

	wktree.State = dxo.FileStateUntracked

	if dotgit != nil {
		if dotgit.Exists() {
			dotgitParent := dotgit.GetParent()
			wktree.Name = dotgitParent.GetName()
			wktree.WorkingDir = dotgitParent.GetPath()
			wktree.DotGitPath = dotgit.GetPath()
		}
	}

	// try find old
	// todo ...

	return wktree, nil
}

func (inst *RWPServiceImpl) findProject(ctx context.Context, path afs.Path) (*dto.Project, error) {

	// locate project
	project1 := &dto.Project{}
	project1.FullPath = path.GetPath()
	project2, err := inst.Projects.Locate(ctx, project1)
	if err != nil {
		return nil, err
	}

	project2.State = dxo.FileStateUntracked

	// try find old
	// todo ...

	return project2, nil
}

// Find ...
func (inst *RWPServiceImpl) Find(ctx context.Context, o *vo.RepositoryWorktreeProject) (*vo.RepositoryWorktreeProject, error) {

	o.Repository = nil
	o.Worktree = nil
	o.Project = nil

	lib, err := inst.getgitlib()
	if err != nil {
		return nil, err
	}
	path := lib.FS().NewPath(o.Path)
	layout, err := lib.RepositoryLocator().Locate(path)
	if err != nil {
		layout = nil
	}

	// for git repo
	repo, err := inst.findRepository(ctx, layout)
	if err == nil {
		o.Repository = repo
	}

	// for worktree
	wktree, err := inst.findWorktree(ctx, layout)
	if err == nil {
		o.Worktree = wktree
	}

	// for project
	project, err := inst.findProject(ctx, path)
	if err == nil {
		o.Project = project
	}

	return o, nil
}

// Save ...
func (inst *RWPServiceImpl) Save(c context.Context, o *vo.RepositoryWorktreeProject) (*vo.RepositoryWorktreeProject, error) {
	return nil, fmt.Errorf("no impl")
}
