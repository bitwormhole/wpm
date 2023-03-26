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
	older, err := inst.findRepositoryOlder(ctx, repo)
	if older != nil && err == nil {
		return older, nil
	}

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
			wktree.Path = dotgitParent.GetPath()
		}
	}

	// try find old
	older, err := inst.findWorktreeOlder(ctx, wktree)
	if older != nil && err == nil {
		return older, nil
	}

	return wktree, nil
}

func (inst *RWPServiceImpl) findProject(ctx context.Context, path afs.Path) (*dto.Project, error) {

	opt := &service.ProjectOptions{WithFileState: true}

	// locate project
	project1 := &dto.Project{}
	project1.Path = path.GetPath()
	project2, err := inst.Projects.Locate(ctx, project1, opt)
	if err != nil {
		return nil, err
	}

	project2.State = dxo.FileStateUntracked

	// try find old
	older, err := inst.findProjectOlder(ctx, project2)
	if err == nil && older != nil {
		return older, nil
	}

	return project2, nil
}

func (inst *RWPServiceImpl) findProjectOlder(ctx context.Context, want *dto.Project) (*dto.Project, error) {
	path := want.Path
	opt := &service.ProjectOptions{
		WithFileState: true,
		// WithGitStatus: true,
	}
	return inst.Projects.FindByPath(ctx, path, opt)
}

func (inst *RWPServiceImpl) findRepositoryOlder(ctx context.Context, want *dto.LocalRepository) (*dto.LocalRepository, error) {
	path := want.ConfigFile
	opt := &service.LocalRepositoryOptions{
		WithFileState: true,
		WithGitStatus: true,
	}
	return inst.Repositories.FindByPath(ctx, path, opt)
}

func (inst *RWPServiceImpl) findWorktreeOlder(ctx context.Context, want *dto.Worktree) (*dto.Worktree, error) {
	path := want.DotGitPath
	opt := &service.WorktreeOptions{
		WithFileState: true,
		WithGitStatus: false,
	}
	return inst.Worktrees.FindByPath(ctx, path, opt)
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
func (inst *RWPServiceImpl) Save(c context.Context, o1 *vo.RepositoryWorktreeProject) (*vo.RepositoryWorktreeProject, error) {

	repo1 := o1.Repository
	worktree1 := o1.Worktree
	project1 := o1.Project
	o2 := &vo.RepositoryWorktreeProject{}

	if repo1 != nil {
		repo2, err := inst.fetchOrInsertRepository(c, repo1)
		if err != nil {
			return nil, err
		}
		o2.Repository = repo2
		repo1 = repo2
	} else {
		repo1 = &dto.LocalRepository{}
	}

	if worktree1 != nil {
		worktree1.OwnerRepository = repo1.ID
		worktree2, err := inst.fetchOrInsertWorktree(c, worktree1)
		if err != nil {
			return nil, err
		}
		o2.Worktree = worktree2
		worktree1 = worktree2
	} else {
		worktree1 = &dto.Worktree{}
	}

	if project1 != nil {
		project1.OwnerRepository = repo1.ID
		project2, err := inst.fetchOrInsertProject(c, project1)
		if err != nil {
			return nil, err
		}
		o2.Project = project2
	}

	return o2, nil
}

func (inst *RWPServiceImpl) fetchOrInsertRepository(c context.Context, o1 *dto.LocalRepository) (*dto.LocalRepository, error) {
	ser := inst.Repositories
	opt := &service.LocalRepositoryOptions{WithFileState: true}
	return ser.InsertOrFetch(c, o1, opt)
}

func (inst *RWPServiceImpl) fetchOrInsertWorktree(c context.Context, o1 *dto.Worktree) (*dto.Worktree, error) {
	ser := inst.Worktrees
	opt := &service.WorktreeOptions{WithFileState: true}
	return ser.InsertOrFetch(c, o1, opt)
}

func (inst *RWPServiceImpl) fetchOrInsertProject(c context.Context, o1 *dto.Project) (*dto.Project, error) {
	ser := inst.Projects
	opt := &service.ProjectOptions{WithFileState: true}
	return ser.InsertOrFetch(c, o1, opt)
}
