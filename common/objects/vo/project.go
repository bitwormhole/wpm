package vo

import "github.com/bitwormhole/wpm/common/objects/dto"

// Project ...
type Project struct {
	Base

	Projects []*dto.Project `json:"projects"`
}
