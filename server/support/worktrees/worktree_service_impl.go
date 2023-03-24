package worktrees

import (
	"context"
	"fmt"

	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// ImpWorktreeService ...
type ImpWorktreeService struct {
	markup.Component `id:"WorktreeService"`
}

func (inst *ImpWorktreeService) _Impl() service.WorktreeService {
	return inst
}

func (inst *ImpWorktreeService) Find(ctx context.Context, id dxo.WorktreeID) (*dto.Worktree, error) {
	return nil, fmt.Errorf("no impl")

}

func (inst *ImpWorktreeService) ListAll(ctx context.Context) ([]*dto.Worktree, error) {
	return nil, fmt.Errorf("no impl")

}

func (inst *ImpWorktreeService) Insert(ctx context.Context, o *dto.Worktree) (*dto.Worktree, error) {
	return nil, fmt.Errorf("no impl")

}

func (inst *ImpWorktreeService) Update(ctx context.Context, id dxo.WorktreeID, o *dto.Worktree) (*dto.Worktree, error) {
	return nil, fmt.Errorf("no impl")
}

func (inst *ImpWorktreeService) Remove(ctx context.Context, id dxo.WorktreeID) error {

	return fmt.Errorf("no impl")
}
