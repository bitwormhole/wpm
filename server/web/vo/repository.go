package vo

import (
	"github.com/bitwormhole/wpm/server/web/dto"
)

// RemoteRepository 仓库VO
type RemoteRepository struct {
	Base

	// ParamID      string

	Repositories []*dto.RemoteRepository `json:"remote_repositories"`
}

// LocalRepository 仓库VO
type LocalRepository struct {
	Base

	// ParamID      string

	Repositories []*dto.LocalRepository `json:"local_repositories"`
}

// MainRepository 仓库VO
type MainRepository struct {
	Base

	// ParamID      string

	Repository *dto.MainRepository `json:"main_repository"`
}
