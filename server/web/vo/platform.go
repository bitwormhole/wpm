package vo

import "github.com/bitwormhole/wpm/server/web/dto"

// Platform ...
type Platform struct {
	Base

	Platforms []*dto.Platform `json:"platforms"`
}
