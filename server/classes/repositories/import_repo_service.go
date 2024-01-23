package repositories

import (
	"context"

	"github.com/bitwormhole/wpm/common/objects/vo"
)

// ImportRepositoryService ...
type ImportRepositoryService interface {
	Handle(c context.Context, want, have *vo.RepositoryImport) error
}
