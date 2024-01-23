package imprepositories

import (
	"context"
	"fmt"

	"github.com/bitwormhole/gitlib"
	"github.com/bitwormhole/wpm/common/objects/dto"
	"github.com/bitwormhole/wpm/common/objects/dxo"
	"github.com/bitwormhole/wpm/common/objects/vo"
	"github.com/bitwormhole/wpm/server/classes/repositories"
	"github.com/starter-go/afs"
)

// ImportRepositoryServiceImpl ...
type ImportRepositoryServiceImpl struct {

	//starter:component
	_as func(repositories.ImportRepositoryService) //starter:as("#")

	LibAgent    gitlib.Agent                        //starter:inject("#")
	RepoService repositories.LocalRepositoryService //starter:inject("#")

}

func (inst *ImportRepositoryServiceImpl) _impl() repositories.ImportRepositoryService {
	return inst
}

// Handle ...
func (inst *ImportRepositoryServiceImpl) Handle(c context.Context, want, have *vo.RepositoryImport) error {
	action := want.Action
	switch action {
	case vo.RepositoryImportActionFind:
		return inst.doFindAll(c, want, have)

	case vo.RepositoryImportActionLocate:
		return inst.doLocateAll(c, want, have)

	case vo.RepositoryImportActionSave:
		return inst.doSaveAll(c, want, have)

	default:
		return fmt.Errorf("bad RepositoryImportAction: %s", action)
	}
}

func (inst *ImportRepositoryServiceImpl) getPathString(path afs.Path) string {
	if path == nil {
		return ""
	}
	return path.GetPath()
}

func (inst *ImportRepositoryServiceImpl) doLocateOne(c context.Context, want *dto.LocalRepository) (*dto.LocalRepository, error) {

	lib := inst.LibAgent.GetLib()
	pwd := lib.FS().NewPath(want.Path)
	layout, err := lib.Locator().Locate(pwd)
	if err != nil {
		return nil, err
	}

	repoDir := layout.Repository()
	repo := inst.getPathString(repoDir)
	name := repoDir.GetName()

	if name == ".git" {
		name = repoDir.GetParent().GetName()
	}

	result := new(dto.LocalRepository)
	result.Path = repo
	result.RegularPath = repo
	result.RepositoryPath = repo
	result.WorkingPath = inst.getPathString(layout.Workspace())
	result.DotGitPath = inst.getPathString(layout.DotGit())
	result.ConfigFile = inst.getPathString(layout.Config())
	result.Name = name
	result.DisplayName = name
	result.State = dxo.FileStateUntracked

	return result, nil
}

func (inst *ImportRepositoryServiceImpl) doLocateAll(c context.Context, want, have *vo.RepositoryImport) error {
	src := want.Items
	dst := have.Items
	for _, item := range src {
		item2, err := inst.doLocateOne(c, item)
		if err != nil {
			return err
		}
		dst = append(dst, item2)
	}
	have.Items = dst
	return nil
}

func (inst *ImportRepositoryServiceImpl) doFindAt(c context.Context, want *dto.LocalRepository) ([]*dto.LocalRepository, error) {

	return nil, fmt.Errorf("no impl: ImportRepositoryServiceImpl.doFindAt")
}

func (inst *ImportRepositoryServiceImpl) doFindAll(c context.Context, want, have *vo.RepositoryImport) error {
	src := want.Items
	dst := have.Items
	for _, item := range src {
		list, err := inst.doFindAt(c, item)
		if err != nil {
			return err
		}
		dst = append(dst, list...)
	}
	have.Items = dst
	return nil
}

func (inst *ImportRepositoryServiceImpl) doSaveOne(c context.Context, want *dto.LocalRepository) (*dto.LocalRepository, error) {
	ser := inst.RepoService
	return ser.Insert(c, want)
}

func (inst *ImportRepositoryServiceImpl) doSaveAll(c context.Context, want, have *vo.RepositoryImport) error {
	src := want.Items
	dst := have.Items
	for _, item := range src {
		item2, err := inst.doSaveOne(c, item)
		if err != nil {
			return err
		}
		dst = append(dst, item2)
	}
	have.Items = dst
	return nil
}
