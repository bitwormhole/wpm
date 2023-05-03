package vo

import (
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// FileQuery ...
type FileQuery struct {
	Base

	BaseURL    string                `json:"base"` // 'file:///' or 'repository:///'
	Repository dxo.LocalRepositoryID `json:"repository"`

	Branch string `json:"branch"` // the git branch name
	Commit string `json:"commit"` // the git commit id
	Tag    string `json:"tag"`    // the git tag name
	Path   string `json:"path"`   // the path for filesystem or worktree

	Self dto.File `json:"self"`

	OptionWithContentType bool `json:"option-with-content-type"`

	Items []*dto.File        `json:"items"`
	Types []*dto.ContentType `json:"contenttypes"`
}
