package service

import (
	"context"

	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// PipeService ...
type PipeService interface {
	Find(ctx context.Context, id dxo.PipeID) (*dto.Pipe, error)

	ListAll(ctx context.Context) ([]*dto.Pipe, error)

	Insert(ctx context.Context, o *dto.Pipe) (*dto.Pipe, error)
	Update(ctx context.Context, id dxo.PipeID, o *dto.Pipe) (*dto.Pipe, error)
	Remove(ctx context.Context, id dxo.PipeID) error
}
