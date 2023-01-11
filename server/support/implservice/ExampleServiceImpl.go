package implservice

import (
	"context"
	"errors"

	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/data/entity"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// ExampleServiceImpl ...
type ExampleServiceImpl struct {
	markup.Component `id:"ExampleService"`

	// ExampleDAO dao.ExampleDAO `inject:"#ExampleDAO"`
}

func (inst *ExampleServiceImpl) _Impl() service.ExampleService {
	return inst
}

func (inst *ExampleServiceImpl) dto2entity(o1 *dto.Example) (*entity.Example, error) {
	o2 := &entity.Example{}
	o2.ID = o1.ID
	// o2.Executable = o1.Executable
	// todo ...
	return o2, nil
}

func (inst *ExampleServiceImpl) entity2dto(o1 *entity.Example) (*dto.Example, error) {
	o2 := &dto.Example{}
	o2.ID = o1.ID
	// o2.Executable = o1.Executable
	// todo ...
	return o2, nil
}

// ListAll ...
func (inst *ExampleServiceImpl) ListAll(ctx context.Context) ([]*dto.Example, error) {

	// src, err := inst.ExampleDAO.ListAll()
	// if err != nil {
	// 	return nil, err
	// }
	// dst := make([]*dto.Example, 0)
	// for _, item1 := range src {
	// 	item2, err := inst.entity2dto(item1)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	dst = append(dst, item2)
	// }
	// return dst, nil

	return nil, errors.New("no impl")
}

// Find ...
func (inst *ExampleServiceImpl) Find(ctx context.Context, id dxo.ExampleID) (*dto.Example, error) {
	return nil, errors.New("no impl")
}

// Insert ...
func (inst *ExampleServiceImpl) Insert(ctx context.Context, o *dto.Example) (*dto.Example, error) {
	return nil, errors.New("no impl")
}

// Update ...
func (inst *ExampleServiceImpl) Update(ctx context.Context, id dxo.ExampleID, o *dto.Example) (*dto.Example, error) {
	return nil, errors.New("no impl")
}

// Remove ...
func (inst *ExampleServiceImpl) Remove(ctx context.Context, id dxo.ExampleID) error {
	return errors.New("no impl")
}
