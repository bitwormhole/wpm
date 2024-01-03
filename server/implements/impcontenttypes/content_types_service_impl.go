package impcontenttypes

import (
	"context"
	"fmt"

	"github.com/bitwormhole/wpm/server/classes/contenttypes"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// ServiceImpl ...
type ServiceImpl struct {

	//starter:component
	_as func(contenttypes.Service) //starter:as("#")

	Dao contenttypes.DAO //starter:inject("#")
}

func (inst *ServiceImpl) _impl() contenttypes.Service {
	return inst
}

// Insert ...
func (inst *ServiceImpl) Insert(ctx context.Context, o *dto.ContentType) (*dto.ContentType, error) {
	return nil, fmt.Errorf("no impl")
}

// Update ...
func (inst *ServiceImpl) Update(ctx context.Context, id dxo.ContentTypeID, o *dto.ContentType) (*dto.ContentType, error) {
	return nil, fmt.Errorf("no impl")
}

// Remove ...
func (inst *ServiceImpl) Remove(ctx context.Context, id dxo.ContentTypeID) error {
	return fmt.Errorf("no impl")
}

// Find ...
func (inst *ServiceImpl) Find(ctx context.Context, id dxo.ContentTypeID) (*dto.ContentType, error) {
	return nil, fmt.Errorf("no impl")
}

// List ...
func (inst *ServiceImpl) List(ctx context.Context, q *contenttypes.Query) ([]*dto.ContentType, error) {
	return nil, fmt.Errorf("no impl")
}
