package vo

import "github.com/bitwormhole/wpm/app/web/dto"

// Project ...
type Project struct {
	Base

	Projects []*dto.Project `json:"projects"`
}
