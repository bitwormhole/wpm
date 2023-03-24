package service

import (
	"context"

	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// WorktreeService ...
type WorktreeService interface {
	Find(ctx context.Context, id dxo.WorktreeID) (*dto.Worktree, error)

	ListAll(ctx context.Context) ([]*dto.Worktree, error)

	Insert(ctx context.Context, o *dto.Worktree) (*dto.Worktree, error)
	Update(ctx context.Context, id dxo.WorktreeID, o *dto.Worktree) (*dto.Worktree, error)
	Remove(ctx context.Context, id dxo.WorktreeID) error
}
