package ilocations

import (
	"context"
	"fmt"

	"github.com/bitwormhole/wpm/common/objects/dto"
	"github.com/bitwormhole/wpm/common/objects/dxo"
	"github.com/bitwormhole/wpm/server/classes/locations"
)

// ServiceImpl ...
type ServiceImpl struct {

	//starter:component
	_as func(locations.Service) //starter:as("#")

	Dao locations.DAO //starter:inject("#")
}

func (inst *ServiceImpl) _impl() locations.Service {
	return inst
}

// Insert ...
func (inst *ServiceImpl) Insert(ctx context.Context, o *dto.Location) (*dto.Location, error) {
	return nil, fmt.Errorf("no impl")
}

// Update ...
func (inst *ServiceImpl) Update(ctx context.Context, id dxo.LocationID, o *dto.Location) (*dto.Location, error) {
	return nil, fmt.Errorf("no impl")
}

// Remove ...
func (inst *ServiceImpl) Remove(ctx context.Context, id dxo.LocationID) error {
	return fmt.Errorf("no impl")
}

// Find ...
func (inst *ServiceImpl) Find(ctx context.Context, id dxo.LocationID) (*dto.Location, error) {
	return nil, fmt.Errorf("no impl")
}

// List ...
func (inst *ServiceImpl) List(ctx context.Context, q *locations.Query) ([]*dto.Location, error) {
	return nil, fmt.Errorf("no impl")
}
