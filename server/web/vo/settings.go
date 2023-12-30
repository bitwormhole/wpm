package vo

import "github.com/bitwormhole/wpm/server/web/dto"

// Settings ...
type Settings struct {
	Base

	Settings *dto.Settings `json:"settings"`

	Table map[string]string `json:"table"`
}
