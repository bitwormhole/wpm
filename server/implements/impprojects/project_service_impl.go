package impprojects

import (
	"context"
	"fmt"

	"github.com/bitwormhole/wpm/common/objects/dto"
	"github.com/bitwormhole/wpm/common/objects/dxo"
	"github.com/bitwormhole/wpm/server/classes/projects"
)

// ServiceImpl ...
type ServiceImpl struct {

	//starter:component
	_as func(projects.Service) //starter:as("#")

	Dao projects.DAO //starter:inject("#")
}

func (inst *ServiceImpl) _impl() projects.Service {
	return inst
}

// Insert ...
func (inst *ServiceImpl) Insert(ctx context.Context, o *dto.Project) (*dto.Project, error) {
	return nil, fmt.Errorf("no impl")
}

// Update ...
func (inst *ServiceImpl) Update(ctx context.Context, id dxo.ProjectID, o *dto.Project) (*dto.Project, error) {
	return nil, fmt.Errorf("no impl")
}

// Remove ...
func (inst *ServiceImpl) Remove(ctx context.Context, id dxo.ProjectID) error {
	return fmt.Errorf("no impl")
}

// Find ...
func (inst *ServiceImpl) Find(ctx context.Context, id dxo.ProjectID) (*dto.Project, error) {
	return nil, fmt.Errorf("no impl")
}

// List ...
func (inst *ServiceImpl) List(ctx context.Context, q *projects.Query) ([]*dto.Project, error) {
	list1, err := inst.Dao.List(nil, q)
	if err != nil {
		return nil, err
	}
	list2 := projects.NewConvertor().ConvertListE2D(list1)
	return list2, nil
}