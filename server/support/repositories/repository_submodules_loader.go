package repositories

import (
	"github.com/bitwormhole/gitlib/git/store"
	"github.com/bitwormhole/starter/vlog"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/web/dto"
)

type mySubmodulesLoader struct {
	lib store.Lib
}

func (inst *mySubmodulesLoader) Load(path string, owner dxo.LocalRepositoryID) ([]*dto.Submodule, error) {

	at := inst.lib.FS().NewPath(path)
	repo, err := inst.lib.RepositoryLoader().LoadWithPath(at)
	if err != nil {
		return nil, err
	}

	src := repo.Submodules().List()
	dst := make([]*dto.Submodule, 0)

	for _, item1 := range src {
		item2 := &dto.Submodule{}
		item2.OwnerRepository = owner
		err = inst.loadItem(item1, item2)
		if err == nil {
			dst = append(dst, item2)
		} else {
			vlog.Warn(err)
		}
	}

	return dst, nil
}

func (inst *mySubmodulesLoader) loadItem(src store.Submodule, dst *dto.Submodule) error {

	name := src.Name()
	url := src.URL()
	path := src.Path()
	active := src.IsActive()

	dst.Name = name
	dst.IsSubmodule = true
	dst.Active = active
	dst.URL = url
	dst.RawPath = path

	if src.Exists() && active {
		workspace := src.Workspace()
		dotgit := workspace.DotGit().GetPath()
		wkdir := workspace.WorkingDirectory().GetPath()
		dst.DotGitPath = dotgit
		dst.WorkingDir = wkdir
		dst.Path = wkdir
	}

	return nil
}
