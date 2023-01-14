package implservice

import (
	"fmt"

	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// WindowsPlatformServiceImpl ...
type WindowsPlatformServiceImpl struct {
	markup.Component `class:"PlatformServiceRegistry"`
}

func (inst *WindowsPlatformServiceImpl) _Impl() (service.PlatformService, service.PlatformServiceRegistry) {
	return inst, inst
}

// GetRegistration ...
func (inst *WindowsPlatformServiceImpl) GetRegistration() *service.PlatformServiceRegistration {
	return &service.PlatformServiceRegistration{
		Service: inst,
	}
}

// Accept ...
func (inst *WindowsPlatformServiceImpl) Accept(p *dto.Platform) bool {
	if p == nil {
		return false
	}
	return (p.OS == "windows")
}

// GetInfo ...
func (inst *WindowsPlatformServiceImpl) GetInfo(p *dto.Platform) (*dto.Platform, error) {
	const (
		keyUser = "USERNAME"
		keyHome = "USERPROFILE"
	)
	kvs, err := getRequiredEnvironmentValues(keyHome, keyUser)
	if err != nil {
		return nil, err
	}
	if p == nil {
		return nil, fmt.Errorf("param is nil")
	}
	p.Home = kvs[keyHome]
	p.User = kvs[keyUser]
	return p, nil
}
