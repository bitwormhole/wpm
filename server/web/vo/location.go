package vo

import "github.com/bitwormhole/wpm/server/web/dto"

// Location ...
type Location struct {
	Base

	Locations []*dto.Location `json:"locations"`
}
