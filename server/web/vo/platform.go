package vo

import "github.com/bitwormhole/wpm/server/web/dto"

// Platform ...
type Platform struct {
	Base

	Platform *dto.Platform `json:"platform"`
}
