package vo

import "github.com/bitwormhole/wpm/common/objects/dto"

// Platform ...
type Platform struct {
	Base

	Platform *dto.Platform `json:"platform"`
}
