package service

import (
	"context"

	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// ExampleService ...
type ExampleService interface {
	Find(ctx context.Context, id dxo.ExampleID) (*dto.Example, error)

	ListAll(ctx context.Context) ([]*dto.Example, error)

	Insert(ctx context.Context, o *dto.Example) (*dto.Example, error)
	Update(ctx context.Context, id dxo.ExampleID, o *dto.Example) (*dto.Example, error)
	Remove(ctx context.Context, id dxo.ExampleID) error
}
