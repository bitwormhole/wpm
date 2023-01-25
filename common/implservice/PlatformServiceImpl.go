package implservice

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// PlatformServiceImpl ...
type PlatformServiceImpl struct {
	markup.Component `id:"PlatformService"`

	Providers []service.PlatformProviderRegistry `inject:".PlatformProviderRegistry"`

	cachedProvider service.PlatformProvider
}

func (inst *PlatformServiceImpl) _Impl() service.PlatformService {
	return inst
}

// GetProvider ...
func (inst *PlatformServiceImpl) GetProvider() (service.PlatformProvider, error) {

	provider := inst.cachedProvider
	if provider != nil {
		return provider, nil
	}

	pf, err := inst.GetPlatform()
	if err != nil {
		return nil, err
	}

	all := inst.Providers
	for _, reg1 := range all {
		reg2 := reg1.GetRegistration()
		provider := reg2.Provider
		if provider.Accept(pf) {
			inst.cachedProvider = provider
			return provider, nil
		}
	}

	arch := pf.Arch
	os := pf.OS
	return nil, fmt.Errorf("unsupported platform, os:%v arch:%v", os, arch)
}

// GetPlatform ...
func (inst *PlatformServiceImpl) GetPlatform() (*dto.Platform, error) {
	arch := runtime.GOARCH
	os := runtime.GOOS
	p := &dto.Platform{}
	p.Arch = strings.ToLower(arch)
	p.OS = strings.ToLower(os)
	return p, nil
}

////////////////////////////////////////////////////////////////////////////////

func getRequiredEnvironmentValues(keys ...string) (map[string]string, error) {
	dst := make(map[string]string)
	for _, key := range keys {
		value := os.Getenv(key)
		if value == "" {
			return nil, fmt.Errorf("no required env:%v", key)
		}
		dst[key] = value
	}
	return dst, nil
}
