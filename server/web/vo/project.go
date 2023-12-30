package vo

import "github.com/bitwormhole/wpm/server/web/dto"

// Project ...
type Project struct {
	Base

	Projects []*dto.Project `json:"projects"`
}
