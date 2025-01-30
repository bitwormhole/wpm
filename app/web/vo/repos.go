package vo

import "github.com/bitwormhole/wpm-api/v1/dto"

// LocalRepository ...
type LocalRepository struct {
	Base

	Items []*dto.LocalRepository `json:"local_repositories"`
}

// ImportRepos ...
type ImportRepos struct {
	Base

	Action string                 `json:"action"`
	Items  []*dto.LocalRepository `json:"local_repositories"`
}
