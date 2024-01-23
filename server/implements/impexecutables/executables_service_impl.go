package impexecutables

import (
	"context"
	"fmt"

	"github.com/bitwormhole/wpm/common/objects/dto"
	"github.com/bitwormhole/wpm/common/objects/dxo"
	"github.com/bitwormhole/wpm/server/classes/executables"
)

// ServiceImpl ...
type ServiceImpl struct {

	//starter:component
	_as func(executables.Service) //starter:as("#")

	Dao executables.DAO //starter:inject("#")

	convertor executables.Convertor
}

func (inst *ServiceImpl) _impl() executables.Service {
	return inst
}

// Insert ...
func (inst *ServiceImpl) Insert(ctx context.Context, o *dto.Executable) (*dto.Executable, error) {
	return nil, fmt.Errorf("no impl")
}

// Update ...
func (inst *ServiceImpl) Update(ctx context.Context, id dxo.ExecutableID, o *dto.Executable) (*dto.Executable, error) {
	return nil, fmt.Errorf("no impl")
}

// Remove ...
func (inst *ServiceImpl) Remove(ctx context.Context, id dxo.ExecutableID) error {
	return fmt.Errorf("no impl")
}

// Find ...
func (inst *ServiceImpl) Find(ctx context.Context, id dxo.ExecutableID) (*dto.Executable, error) {
	return nil, fmt.Errorf("no impl")
}

// List ...
func (inst *ServiceImpl) List(ctx context.Context, q *executables.Query) ([]*dto.Executable, error) {
	list1, err := inst.Dao.List(nil, q)
	if err != nil {
		return nil, err
	}
	list2 := inst.convertor.ConvertListE2D(list1)
	return list2, nil
}
