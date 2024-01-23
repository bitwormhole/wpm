package impintenttemplates

import (
	"context"
	"fmt"

	"github.com/bitwormhole/wpm/common/objects/dto"
	"github.com/bitwormhole/wpm/common/objects/dxo"
	"github.com/bitwormhole/wpm/server/classes/intenttemplates"
)

// ServiceImpl ...
type ServiceImpl struct {

	//starter:component
	_as func(intenttemplates.Service) //starter:as("#")

	Dao intenttemplates.DAO //starter:inject("#")

	convertor intenttemplates.Convertor
}

func (inst *ServiceImpl) _impl() intenttemplates.Service {
	return inst
}

// Insert ...
func (inst *ServiceImpl) Insert(ctx context.Context, o *dto.IntentTemplate) (*dto.IntentTemplate, error) {
	return nil, fmt.Errorf("no impl")
}

// Update ...
func (inst *ServiceImpl) Update(ctx context.Context, id dxo.IntentTemplateID, o *dto.IntentTemplate) (*dto.IntentTemplate, error) {
	return nil, fmt.Errorf("no impl")
}

// Remove ...
func (inst *ServiceImpl) Remove(ctx context.Context, id dxo.IntentTemplateID) error {
	return fmt.Errorf("no impl")
}

// Find ...
func (inst *ServiceImpl) Find(ctx context.Context, id dxo.IntentTemplateID) (*dto.IntentTemplate, error) {
	return nil, fmt.Errorf("no impl")
}

// List ...
func (inst *ServiceImpl) List(ctx context.Context, q *intenttemplates.Query) ([]*dto.IntentTemplate, error) {
	list1, err := inst.Dao.List(nil, q)
	if err != nil {
		return nil, err
	}
	list2 := inst.convertor.ConvertListE2D(list1)
	return list2, nil
}
