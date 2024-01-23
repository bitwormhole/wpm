package impauths

import (
	"context"
	"fmt"

	"github.com/bitwormhole/wpm/common/objects/dto"
	"github.com/bitwormhole/wpm/common/objects/dxo"
	"github.com/bitwormhole/wpm/server/classes/auths"
)

// ServiceImpl ...
type ServiceImpl struct {

	//starter:component
	_as func(auths.Service) //starter:as("#")

	Dao auths.DAO //starter:inject("#")
}

func (inst *ServiceImpl) _impl() auths.Service {
	return inst
}

// Insert ...
func (inst *ServiceImpl) Insert(ctx context.Context, o *dto.Example) (*dto.Example, error) {
	return nil, fmt.Errorf("no impl")
}

// Update ...
func (inst *ServiceImpl) Update(ctx context.Context, id dxo.ExampleID, o *dto.Example) (*dto.Example, error) {
	return nil, fmt.Errorf("no impl")
}

// Remove ...
func (inst *ServiceImpl) Remove(ctx context.Context, id dxo.ExampleID) error {
	return fmt.Errorf("no impl")
}

// Find ...
func (inst *ServiceImpl) Find(ctx context.Context, id dxo.ExampleID) (*dto.Example, error) {
	return nil, fmt.Errorf("no impl")
}

// List ...
func (inst *ServiceImpl) List(ctx context.Context, q *auths.Query) ([]*dto.Example, error) {
	return nil, fmt.Errorf("no impl")
}
