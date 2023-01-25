package implservice

import (
	"fmt"

	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// LinuxPlatformServiceImpl ...
type LinuxPlatformServiceImpl struct {
	markup.Component `class:"PlatformProviderRegistry"`
}

func (inst *LinuxPlatformServiceImpl) _Impl() (service.PlatformProvider, service.PlatformProviderRegistry) {
	return inst, inst
}

// GetRegistration ...
func (inst *LinuxPlatformServiceImpl) GetRegistration() *service.PlatformProviderRegistration {
	return &service.PlatformProviderRegistration{
		Provider: inst,
	}
}

// Accept ...
func (inst *LinuxPlatformServiceImpl) Accept(p *dto.Platform) bool {
	if p == nil {
		return false
	}
	return (p.OS == "linux")
}

// GetProfile ...
func (inst *LinuxPlatformServiceImpl) GetProfile(p *dto.Profile) error {
	const (
		keyUser = "USER"
		keyHome = "HOME"
	)
	if p == nil {
		return fmt.Errorf("param is nil")
	}
	kvs, err := getRequiredEnvironmentValues(keyHome, keyUser)
	if err != nil {
		return err
	}
	p.Home = kvs[keyHome]
	p.User = kvs[keyUser]
	return nil
}
