package service

import (
	"context"

	"github.com/bitwormhole/wpm/server/web/vo"
)

// JSONBufferService ...
type JSONBufferService interface {
	Get(ctx context.Context) (*vo.Online, error)

	Reset(ctx context.Context) error

	Put(ctx context.Context, o *vo.Online) error
}
