package service

import (
	"github.com/bitwormhole/wpm/server/web/dto"
)

// PlatformService ...
type PlatformService interface {
	GetPlatform() (*dto.Platform, error)
	GetProvider() (PlatformProvider, error)
}

// PlatformProvider ...
type PlatformProvider interface {
	Accept(p *dto.Platform) bool
	GetProfile(p *dto.Profile) error
}

// PlatformProviderRegistration ...
type PlatformProviderRegistration struct {
	Provider PlatformProvider
}

// PlatformProviderRegistry ...
type PlatformProviderRegistry interface {
	GetRegistration() *PlatformProviderRegistration
}
