package service

import (
	"context"

	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// WorktreeOptions ...
type WorktreeOptions struct {
	All           bool
	WithFileState bool
	WithGitStatus bool
}

// WorktreeService ...
type WorktreeService interface {

	// getter

	Find(ctx context.Context, id dxo.WorktreeID, opt *WorktreeOptions) (*dto.Worktree, error)

	FindByPath(ctx context.Context, path string, opt *WorktreeOptions) (*dto.Worktree, error)

	ListAll(ctx context.Context, opt *WorktreeOptions) ([]*dto.Worktree, error)

	// setter

	Insert(ctx context.Context, o *dto.Worktree) (*dto.Worktree, error)
	Update(ctx context.Context, id dxo.WorktreeID, o *dto.Worktree) (*dto.Worktree, error)
	Remove(ctx context.Context, id dxo.WorktreeID) error

	InsertOrFetch(ctx context.Context, o *dto.Worktree, opt *WorktreeOptions) (*dto.Worktree, error)
}
