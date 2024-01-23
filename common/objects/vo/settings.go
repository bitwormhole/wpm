package vo

import "github.com/bitwormhole/wpm/common/objects/dto"

// Settings ...
type Settings struct {
	Base

	Settings *dto.Settings `json:"settings"`

	Table map[string]string `json:"table"`
}
