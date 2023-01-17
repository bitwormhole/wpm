package implservice

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

	lib, err := inst.GitLibAgent.GetLib()
	if err != nil {
		return nil, err
	}

	pwd := lib.FS().NewPath(path)
	layout, err := lib.RepositoryLocator().Locate(pwd)
	if err != nil {
		return nil, err
	}

	targetPath := layout.DotGit()
	if targetPath == nil {
		targetPath = layout.Repository()
	}
	if targetPath == nil {
		return nil, fmt.Errorf("repo path not found")
	}

	result := &dto.LocalRepository{
		Path: targetPath.GetPath(),
	}

	return result, nil
}
