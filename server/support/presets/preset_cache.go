package presets

import (
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/components/caches"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/vo"
)

// Cache ... [inject:"#PresetCache"]
type Cache interface {
	Get() (*vo.Online, error)
}

////////////////////////////////////////////////////////////////////////////////

// CacheProvider ... for Presets
type CacheProvider struct {
	markup.Component ` id:"PresetCache"  class:"wpm-cache-provider"`

	AC application.Context `inject:"context"`

	CS service.CacheService `inject:"#CacheService"`

	ListFileName string `inject:"${wpm.presets.list-file-name}"`
}

func (inst *CacheProvider) _Impl() (caches.ProviderRegistry, caches.Provider, Cache) {
	return inst, inst, inst
}

// ListProviderRegistrations ...
func (inst *CacheProvider) ListProviderRegistrations() []*caches.ProviderRegistration {
	cn := inst.ClassName()
	reg := &caches.ProviderRegistration{
		ClassName: cn,
		Enabled:   true,
		Provider:  inst,
	}
	return []*caches.ProviderRegistration{reg}
}

// ClassName 缓存类型名称
func (inst *CacheProvider) ClassName() string {
	return "presets"
}

// New ...
func (inst *CacheProvider) New() (any, error) {
	loader := &myPresetsLoader{
		ac:           inst.AC,
		listFileName: inst.ListFileName,
	}
	o, err := loader.Load()
	return o, err
}

// Get ...
func (inst *CacheProvider) Get() (*vo.Online, error) {
	o, err := inst.CS.GetManager().Get(inst)
	if err != nil {
		return nil, err
	}
	return o.(*vo.Online), nil
}
