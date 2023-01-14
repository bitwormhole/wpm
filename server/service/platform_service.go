package service

import (
	"github.com/bitwormhole/wpm/server/web/dto"
)

// PlatformService ...
type PlatformService interface {
	Accept(p *dto.Platform) bool

	GetInfo(p *dto.Platform) (*dto.Platform, error)
}

// PlatformServiceRegistration ...
type PlatformServiceRegistration struct {
	Service PlatformService
}

// PlatformServiceRegistry ...
type PlatformServiceRegistry interface {
	GetRegistration() *PlatformServiceRegistration
}
