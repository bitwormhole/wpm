package filequery

import (
	"context"

	"github.com/bitwormhole/wpm/common/objects/vo"
)

// Service ...
type Service interface {
	Query(ctx context.Context, want *vo.FileQuery) (*vo.FileQuery, error)
}