package dao

import (
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/data/entity"
)

// WorktreeDAO ...
type WorktreeDAO interface {
	Find(id dxo.WorktreeID) (*entity.Worktree, error)

	FindByPath(path string) (*entity.Worktree, error)

	ListAll() ([]*entity.Worktree, error)

	Insert(o *entity.Worktree) (*entity.Worktree, error)

	Update(id dxo.WorktreeID, o *entity.Worktree) (*entity.Worktree, error)

	Remove(id dxo.WorktreeID) error
}
