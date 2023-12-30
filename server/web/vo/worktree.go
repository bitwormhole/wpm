package vo

import "github.com/bitwormhole/wpm/server/web/dto"

// Worktree 表示对若干工作树的查询
type Worktree struct {
	Base

	Worktrees []*dto.Worktree `json:"worktrees"`
}
