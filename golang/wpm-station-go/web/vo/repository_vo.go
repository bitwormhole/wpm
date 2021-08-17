package vo

import (
	"github.com/bitwormhole/wpm/web/base"
	"github.com/bitwormhole/wpm/web/dto"
)

// Repository 仓库VO
type Repository struct {
	base.VO

	ParamID      string
	Repositories []*dto.Repository
}
