package presets

import (
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/vo"
)

// ImpPresetService ...
type ImpPresetService struct {
	markup.Component `id:"PresetService"`

	Cache Cache `inject:"#PresetCache"`
}

func (inst *ImpPresetService) _Impl() service.PresetService {
	return inst
}

// GetPresets ...
func (inst *ImpPresetService) GetPresets() (*vo.Online, error) {
	return inst.Cache.Get()
}
