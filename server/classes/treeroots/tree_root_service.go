package treeroots

import (
	"context"

	"github.com/bitwormhole/wpm/common/objects/dto"
	"github.com/bitwormhole/wpm/common/objects/dxo"
)

// Service for TreeRoot
type Service interface {
	Insert(ctx context.Context, o *dto.TreeRoot) (*dto.TreeRoot, error)

	Update(ctx context.Context, id dxo.TreeRootID, o *dto.TreeRoot) (*dto.TreeRoot, error)

	Remove(ctx context.Context, id dxo.TreeRootID) error

	Find(ctx context.Context, id dxo.TreeRootID) (*dto.TreeRoot, error)

	List(ctx context.Context, q *Query) ([]*dto.TreeRoot, error)
}
