package vo

import "github.com/bitwormhole/wpm/common/objects/dto"

// Profile ...
type Profile struct {
	Base

	Platform *dto.Platform `json:"platform"`
	Profile  *dto.Profile  `json:"profile"`
}
