package service

import (
	"context"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/wpm/server/web/vo"
)

// AboutService ...
type AboutService interface {
	GetInfo(ctx context.Context) (*vo.About, error)
}

////////////////////////////////////////////////////////////////////////////////

var theAppModule application.Module

// SetAppModule ...
func SetAppModule(m application.Module) {
	if m == nil {
		return
	}
	theAppModule = m
}

// GetAppModule ...
func GetAppModule() application.Module {
	return theAppModule
}
