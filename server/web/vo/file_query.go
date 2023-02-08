package vo

import (
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// FileQuery ...
type FileQuery struct {
	Base

	Repository dxo.LocalRepositoryID `json:"repository"`

	Branch string `json:"branch"` // the git branch name
	Commit string `json:"commit"` // the git commit id
	Tag    string `json:"tag"`    // the git tag name
	Path   string `json:"path"`   // the path for filesystem or worktree

	Self dto.File `json:"self"`

	Items []*dto.File `json:"items"`
}
