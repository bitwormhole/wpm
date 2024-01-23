package vo

import "github.com/bitwormhole/wpm/common/objects/dto"

// RepositoryWorktreeProject ...
type RepositoryWorktreeProject struct {
	Base

	Path string `json:"path"`

	Repository *dto.LocalRepository `json:"repository"`
	Worktree   *dto.Worktree        `json:"worktree"`
	Project    *dto.Project         `json:"project"`
}
