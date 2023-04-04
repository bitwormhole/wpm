package caches

import (
	"fmt"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/components/caches"
	"github.com/bitwormhole/wpm/server/service"
)

// ImpCacheManager ...
type ImpCacheManager struct {
	markup.Component `id:"CacheService" class:"life"`

	ProviderRegistryList []caches.ProviderRegistry `inject:".wpm-cache-provider"`

	providers map[string]caches.Provider

	cache map[string]any
}

func (inst *ImpCacheManager) _Impl() (service.CacheService, caches.Manager, application.LifeRegistry) {
	return inst, inst, inst
}

// GetLifeRegistration ...
func (inst *ImpCacheManager) GetLifeRegistration() *application.LifeRegistration {
	return &application.LifeRegistration{
		OnInit: inst.init,
	}
}

func (inst *ImpCacheManager) init() error {

	if inst.cache == nil {
		inst.cache = make(map[string]any)
	}

	_, err := inst.getProviderTable()
	return err
}

func (inst *ImpCacheManager) reset() {
	inst.cache = make(map[string]any)
}

// GetManager ...
func (inst *ImpCacheManager) GetManager() caches.Manager {
	inst.init()
	return inst
}

// Clean ...
func (inst *ImpCacheManager) Clean() {
	inst.reset()
}

// Get 取缓存对象
func (inst *ImpCacheManager) Get(p caches.Provider) (any, error) {

	if p == nil {
		return nil, fmt.Errorf("provider is nil")
	}

	cname := p.ClassName()
	o := inst.cache[cname]
	if o != nil {
		return o, nil
	}

	o2, err := inst.load(cname)
	if err != nil {
		return nil, err
	}

	inst.cache[cname] = o2
	return o2, nil
}

func (inst *ImpCacheManager) load(cname string) (any, error) {
	p, err := inst.findProviderByName(cname)
	if err != nil {
		return nil, err
	}
	return p.New()
}

func (inst *ImpCacheManager) findProviderByName(cname string) (caches.Provider, error) {
	table, err := inst.getProviderTable()
	if err != nil {
		return nil, err
	}
	p := table[cname]
	if p == nil {
		return nil, fmt.Errorf("no cache provider with name: " + cname)
	}
	return p, nil
}

func (inst *ImpCacheManager) getProviderTable() (map[string]caches.Provider, error) {
	tab := inst.providers
	if tab != nil {
		return tab, nil
	}
	tab = make(map[string]caches.Provider)
	src := inst.ProviderRegistryList
	for _, r1 := range src {
		list := r1.ListProviderRegistrations()
		for _, r2 := range list {
			name := r2.ClassName
			old := tab[name]
			if old != nil {
				return nil, fmt.Errorf("the name of cache is duplicate: " + name)
			}
			tab[name] = r2.Provider
		}
	}
	inst.providers = tab
	return tab, nil
}
