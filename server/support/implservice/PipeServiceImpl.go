package implservice

import (
	"context"
	"errors"

	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// PipeServiceImpl ...
type PipeServiceImpl struct {
	markup.Component `id:"PipeService"`

	// PipeDAO dao.PipeDAO `inject:"#PipeDAO"`
}

func (inst *PipeServiceImpl) _Impl() service.PipeService {
	return inst
}

// ListAll ...
func (inst *PipeServiceImpl) ListAll(ctx context.Context) ([]*dto.Pipe, error) {

	// src, err := inst.PipeDAO.ListAll()
	// if err != nil {
	// 	return nil, err
	// }
	// dst := make([]*dto.Pipe, 0)
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
func (inst *PipeServiceImpl) Find(ctx context.Context, id dxo.PipeID) (*dto.Pipe, error) {
	return nil, errors.New("no impl")
}

// Insert ...
func (inst *PipeServiceImpl) Insert(ctx context.Context, o *dto.Pipe) (*dto.Pipe, error) {
	return nil, errors.New("no impl")
}

// Update ...
func (inst *PipeServiceImpl) Update(ctx context.Context, id dxo.PipeID, o *dto.Pipe) (*dto.Pipe, error) {
	return nil, errors.New("no impl")
}

// Remove ...
func (inst *PipeServiceImpl) Remove(ctx context.Context, id dxo.PipeID) error {
	return errors.New("no impl")
}
