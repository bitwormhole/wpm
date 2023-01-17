package service

import (
	"github.com/bitwormhole/wpm/server/web/dto"
)

// PlatformService ...
type PlatformService interface {
	GetPlatform() (*dto.Platform, error)

	GetProfile() (*dto.Profile, error)
}

// PlatformServiceProvider ...
type PlatformServiceProvider interface {
	PlatformService

	Accept(p *dto.Platform) bool
}

// PlatformServiceRegistration ...
type PlatformServiceRegistration struct {
	Provider PlatformServiceProvider
}

// PlatformServiceRegistry ...
type PlatformServiceRegistry interface {
	GetRegistration() *PlatformServiceRegistration
}
