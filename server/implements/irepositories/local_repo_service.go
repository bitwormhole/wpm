package irepositories

import (
	"context"

	"github.com/bitwormhole/gitlib"
	"github.com/bitwormhole/wpm/common/objects/dto"
	"github.com/bitwormhole/wpm/common/objects/dxo"
	"github.com/bitwormhole/wpm/server/classes/repositories"
	"github.com/bitwormhole/wpm/server/data/entity"
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

func (inst *LocalRepositoryService) prepareToInsert(ent *entity.LocalRepository) error {

	err := inst.computeLocationWithEntity(ent)
	if err != nil {
		return err
	}
	name := ent.Name

	if ent.DisplayName == "" {
		ent.DisplayName = name
	}

	if ent.Description == "" {
		ent.Description = name
	}

	return nil
}

// Insert ...
func (inst *LocalRepositoryService) Insert(ctx context.Context, o *dto.LocalRepository) (*dto.LocalRepository, error) {

	o2 := inst.convertor.D2E(o)
	err := inst.prepareToInsert(o2)
	if err != nil {
		return nil, err
	}

	o3, err := inst.Dao.Insert(nil, o2)
	if err != nil {
		return nil, err
	}

	o4 := inst.convertor.E2D(o3)
	return o4, nil
}

// Update ...
func (inst *LocalRepositoryService) Update(ctx context.Context, id dxo.LocalRepositoryID, o1 *dto.LocalRepository) (*dto.LocalRepository, error) {

	o2 := inst.convertor.D2E(o1)

	o3, err := inst.Dao.Update(nil, id, func(ent *entity.LocalRepository) {

		if inst.isLocationChanged(o2, ent) {
			err := inst.computeLocationWithEntity(o2)
			if err == nil {
				ent.Path = o2.Path
				ent.ConfigFile = o2.ConfigFile
				ent.Location = o2.Location
			}
		}

		ent.Name = o2.Name
		ent.DisplayName = o2.DisplayName
		ent.Description = o2.Description

		// ent.URN = o2.URN
	})

	if err != nil {
		return nil, err
	}

	o4 := inst.convertor.E2D(o3)
	return o4, nil
}

func (inst *LocalRepositoryService) computeLocationWithEntity(ent *entity.LocalRepository) error {

	lib := inst.LibAgent.GetLib()
	pwd := lib.FS().NewPath(ent.Path)
	layout, err := lib.Locator().Locate(pwd)
	if err != nil {
		return err
	}

	repoDir := layout.Repository()
	repo := inst.getPathString(repoDir)
	config := inst.getPathString(layout.Config())
	name := repoDir.GetName()

	if name == ".git" {
		name = repoDir.GetParent().GetName()
	}

	ent.Path = repo
	ent.ConfigFile = config
	ent.Location = ent.Location
	ent.Name = name

	return nil
}

func (inst *LocalRepositoryService) isLocationChanged(e1, e2 *entity.LocalRepository) bool {
	b1 := (e1.Path == e2.Path)
	b2 := (e1.Location == e2.Location)
	b3 := (e1.ConfigFile == e2.ConfigFile)
	if b1 && b2 && b3 {
		return false
	}
	return true
}

// Remove ...
func (inst *LocalRepositoryService) Remove(ctx context.Context, id dxo.LocalRepositoryID) error {
	return inst.Dao.Remove(nil, id)
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
