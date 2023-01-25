package implservice

import (
	"fmt"

	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// WindowsPlatformServiceImpl ...
type WindowsPlatformServiceImpl struct {
	markup.Component `class:"PlatformProviderRegistry"`
}

func (inst *WindowsPlatformServiceImpl) _Impl() (service.PlatformProvider, service.PlatformProviderRegistry) {
	return inst, inst
}

// GetRegistration ...
func (inst *WindowsPlatformServiceImpl) GetRegistration() *service.PlatformProviderRegistration {
	return &service.PlatformProviderRegistration{
		Provider: inst,
	}
}

// Accept ...
func (inst *WindowsPlatformServiceImpl) Accept(p *dto.Platform) bool {
	if p == nil {
		return false
	}
	return (p.OS == "windows")
}

// GetProfile ...
func (inst *WindowsPlatformServiceImpl) GetProfile(p *dto.Profile) error {
	const (
		keyUser = "USERNAME"
		keyHome = "USERPROFILE"
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
