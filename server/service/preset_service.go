package service

import "github.com/bitwormhole/wpm/server/web/vo"

// PresetService ...
type PresetService interface {
	GetPresets() (*vo.Online, error)
}
