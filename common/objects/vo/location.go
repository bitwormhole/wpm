package vo

import "github.com/bitwormhole/wpm/common/objects/dto"

// Location ...
type Location struct {
	Base

	Locations []*dto.Location `json:"locations"`
}
