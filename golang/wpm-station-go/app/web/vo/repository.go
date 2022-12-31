package vo

import (
	"github.com/bitwormhole/wpm/app/web/dto"
)

// Repository 仓库VO
type Repository struct {
	Base

	// ParamID      string

	Repositories []*dto.Repository `json:"repositories"`
}
