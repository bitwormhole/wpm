package repositories

import (
	"context"
	"fmt"

	"github.com/bitwormhole/gitlib/git/store"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// LocalRepositoryFinderImpl ...
type LocalRepositoryFinderImpl struct {
	markup.Component `id:"LocalRepositoryFinder"`

	GitLibAgent store.LibAgent `inject:"#git-lib-agent"`
}

func (inst *LocalRepositoryFinderImpl) _Impl() service.LocalRepositoryFinder {
	return inst
}

// Search ...
func (inst *LocalRepositoryFinderImpl) Search(ctx context.Context, path string, depthLimit int) ([]*dto.LocalRepository, error) {

	lib, err := inst.GitLibAgent.GetLib()
	if err != nil {
		return nil, err
	}

	pwd := lib.FS().NewPath(path)
	src, err := lib.RepositoryFinder().Find(pwd)
	if err != nil {
		return nil, err
	}

	dst := make([]*dto.LocalRepository, 0)
	for _, it1 := range src {
		it2 := &dto.LocalRepository{
			Path: it1.GetPath(),
		}
		dst = append(dst, it2)
	}

	return dst, nil
}

// Locate ...
func (inst *LocalRepositoryFinderImpl) Locate(ctx context.Context, path string) (*dto.LocalRepository, error) {

	layout, err := inst.LocateLayout(ctx, path)
	if err != nil {
		return nil, err
	}

	dotgit := layout.DotGit()
	config := layout.Config()
	bare := (dotgit == nil)
	errBadRepo := fmt.Errorf("it is not a git repository at " + path)

	if config == nil {
		return nil, errBadRepo
	}
	if !config.IsFile() {
		return nil, errBadRepo
	}

	repodir := config.GetParent()
	o := &dto.LocalRepository{}

	o.Path = repodir.GetPath()
	o.RepositoryPath = repodir.GetPath()
	o.ConfigFile = config.GetPath()
	o.Bare = bare
	if !bare {
		o.DotGitPath = dotgit.GetPath()
		o.WorkingPath = dotgit.GetParent().GetPath()
	}

	return o, nil
}

// LocateLayout ...
func (inst *LocalRepositoryFinderImpl) LocateLayout(ctx context.Context, path string) (store.RepositoryLayout, error) {
	lib, err := inst.GitLibAgent.GetLib()
	if err != nil {
		return nil, err
	}
	pwd := lib.FS().NewPath(path)
	return lib.RepositoryLocator().Locate(pwd)
}
