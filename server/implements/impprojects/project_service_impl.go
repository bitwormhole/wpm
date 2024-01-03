package impprojects

import (
	"context"
	"fmt"

	"github.com/bitwormhole/wpm/server/classes/projects"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/web/dto"
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
	return nil, fmt.Errorf("no impl")
}
