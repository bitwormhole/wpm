package implservice

import (
	"fmt"

	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// LinuxPlatformServiceImpl ...
type LinuxPlatformServiceImpl struct {
	markup.Component `class:"PlatformServiceRegistry"`
}

func (inst *LinuxPlatformServiceImpl) _Impl() (service.PlatformServiceProvider, service.PlatformServiceRegistry) {
	return inst, inst
}

// GetRegistration ...
func (inst *LinuxPlatformServiceImpl) GetRegistration() *service.PlatformServiceRegistration {
	return &service.PlatformServiceRegistration{
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
func (inst *LinuxPlatformServiceImpl) GetProfile() (*dto.Profile, error) {
	const (
		keyUser = "USER"
		keyHome = "HOME"
	)
	kvs, err := getRequiredEnvironmentValues(keyHome, keyUser)
	if err != nil {
		return nil, err
	}
	p := &dto.Profile{}
	if p == nil {
		return nil, fmt.Errorf("param is nil")
	}
	p.Home = kvs[keyHome]
	p.User = kvs[keyUser]
	return p, nil
}

// GetPlatform ...
func (inst *LinuxPlatformServiceImpl) GetPlatform() (*dto.Platform, error) {
	return nil, fmt.Errorf("use: PlatformService.GetPlatform()")
}
