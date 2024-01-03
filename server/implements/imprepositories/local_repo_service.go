package imprepositories

import (
	"context"
	"fmt"

	"github.com/bitwormhole/gitlib"
	"github.com/bitwormhole/wpm/server/classes/repositories"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/web/dto"
	"github.com/starter-go/afs"
)

// LocalRepositoryService ...
type LocalRepositoryService struct {

	//starter:component
	_as func(repositories.LocalRepositoryService) //starter:as("#")

	Dao      repositories.LocalRepositoryDAO //starter:inject("#")
	LibAgent gitlib.Agent                    //starter:inject("#")

	convertor repositories.LocalRepositoryConvertor
}

func (inst *LocalRepositoryService) _impl() repositories.LocalRepositoryService {
	return inst
}

func (inst *LocalRepositoryService) getPathString(path afs.Path) string {
	if path == nil {
		return ""
	}
	return path.GetPath()
}

func (inst *LocalRepositoryService) prepareToWrite(o *dto.LocalRepository) error {

	lib := inst.LibAgent.GetLib()
	pwd := lib.FS().NewPath(o.Path)
	layout, err := lib.Locator().Locate(pwd)
	if err != nil {
		return err
	}

	repoDir := layout.Repository()
	repo := inst.getPathString(repoDir)
	name := repoDir.GetName()

	if name == ".git" {
		name = repoDir.GetParent().GetName()
	}

	result := o
	result.Path = repo
	result.RegularPath = repo
	result.RepositoryPath = repo
	result.WorkingPath = inst.getPathString(layout.Workspace())
	result.DotGitPath = inst.getPathString(layout.DotGit())
	result.ConfigFile = inst.getPathString(layout.Config())

	if result.Name == "" {
		result.Name = name
	}

	if result.DisplayName == "" {
		result.DisplayName = name
	}

	if result.Description == "" {
		result.Description = name
	}

	return nil
}

// Insert ...
func (inst *LocalRepositoryService) Insert(ctx context.Context, o *dto.LocalRepository) (*dto.LocalRepository, error) {

	err := inst.prepareToWrite(o)
	if err != nil {
		return nil, err
	}

	o2 := inst.convertor.D2E(o)
	o3, err := inst.Dao.Insert(nil, o2)
	if err != nil {
		return nil, err
	}

	o4 := inst.convertor.E2D(o3)
	return o4, nil
}

// Update ...
func (inst *LocalRepositoryService) Update(ctx context.Context, id dxo.LocalRepositoryID, o *dto.LocalRepository) (*dto.LocalRepository, error) {
	return nil, fmt.Errorf("no impl")
}

// Remove ...
func (inst *LocalRepositoryService) Remove(ctx context.Context, id dxo.LocalRepositoryID) error {
	return fmt.Errorf("no impl")
}

// Find ...
func (inst *LocalRepositoryService) Find(ctx context.Context, id dxo.LocalRepositoryID) (*dto.LocalRepository, error) {
	item1, err := inst.Dao.Find(nil, id)
	if err != nil {
		return nil, err
	}
	item2 := inst.convertor.E2D(item1)
	return item2, nil
}

// List ...
func (inst *LocalRepositoryService) List(ctx context.Context, q *repositories.Query) ([]*dto.LocalRepository, error) {
	list1, err := inst.Dao.List(nil, q)
	if err != nil {
		return nil, err
	}
	list2 := inst.convertor.ListE2D(list1)
	return list2, nil
}
