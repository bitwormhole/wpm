package service

import (
	"context"

	"github.com/bitwormhole/wpm/server/web/vo"
)

// FileQueryService ...
type FileQueryService interface {
	Query(ctx context.Context, q *vo.FileQuery) (*vo.FileQuery, error)
}
