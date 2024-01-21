package impmedia

import (
	"context"
	"fmt"

	"github.com/bitwormhole/wpm/server/classes/media"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// ServiceImpl ...
type ServiceImpl struct {

	//starter:component
	_as func(media.Service) //starter:as("#")

	Dao media.DAO //starter:inject("#")

	convertor media.Convertor
}

func (inst *ServiceImpl) _impl() media.Service {
	return inst
}

// Insert ...
func (inst *ServiceImpl) Insert(ctx context.Context, o *dto.Media) (*dto.Media, error) {
	return nil, fmt.Errorf("no impl")
}

// Update ...
func (inst *ServiceImpl) Update(ctx context.Context, id dxo.MediaID, o *dto.Media) (*dto.Media, error) {
	return nil, fmt.Errorf("no impl")
}

// Remove ...
func (inst *ServiceImpl) Remove(ctx context.Context, id dxo.MediaID) error {
	return fmt.Errorf("no impl")
}

// Find ...
func (inst *ServiceImpl) Find(ctx context.Context, id dxo.MediaID) (*dto.Media, error) {
	return nil, fmt.Errorf("no impl")
}

// List ...
func (inst *ServiceImpl) List(ctx context.Context, q *media.Query) ([]*dto.Media, error) {
	list1, err := inst.Dao.List(nil, q)
	if err != nil {
		return nil, err
	}
	list2 := inst.convertor.ConvertListE2D(list1)
	return list2, nil
}
