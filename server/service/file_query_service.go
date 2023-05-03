package service

import (
	"context"

	"github.com/bitwormhole/wpm/server/web/vo"
)

type FileQueryOptions struct {
	WithContentType bool `json:"with-content-type"`
}

// FileQueryService ...
type FileQueryService interface {
	Query(ctx context.Context, q *vo.FileQuery, ops *FileQueryOptions) (*vo.FileQuery, error)
}
