package service

import (
	"context"

	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/web/dto"
	"github.com/bitwormhole/wpm/server/web/vo"
)

// PipeService ...
type PipeService interface {
	Insert(ctx context.Context, o *dto.PipeInfo) (*dto.PipeInfo, error)

	Remove(ctx context.Context, id dxo.PipeID) error

	Pull(ctx context.Context, id dxo.PipeID, view *vo.Pipe, timeout int) error

	Push(ctx context.Context, id dxo.PipeID, view *vo.Pipe) error

	ListAll(ctx context.Context) ([]*dto.PipeInfo, error)

	// Find(ctx context.Context, id dxo.PipeID) (*dto.Pipe, error)
	// Update(ctx context.Context, id dxo.PipeID, o *dto.Pipe) (*dto.Pipe, error)

}
