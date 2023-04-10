package service

import (
	"context"
	"strconv"
	"strings"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/vlog"
	"github.com/bitwormhole/wpm/server/web/vo"
)

// AboutService ...
type AboutService interface {
	GetInfo(ctx context.Context) (*vo.About, error)

	IsDebug() bool
	IsRelease() bool
}

////////////////////////////////////////////////////////////////////////////////

var theAppModule application.Module

// SetAppModule ...
func SetAppModule(m application.Module) {
	if m == nil {
		return
	}
	if theAppModule != nil {
		return
	}
	b := strings.Builder{}
	b.WriteString(m.GetName())
	b.WriteString("?ver=")
	b.WriteString(m.GetVersion())
	b.WriteString("&r=")
	b.WriteString(strconv.Itoa(m.GetRevision()))
	vlog.Info("set wpm.main.module = ", b.String())
	theAppModule = m
}

// GetAppModule ...
func GetAppModule() application.Module {
	m := theAppModule
	if m != nil {
		return m
	}
	mb := application.ModuleBuilder{}
	mb.Name("wpm.bitwormhole.com/no-module-info")
	mb.Version("v0.0.0")
	mb.Revision(0)
	m = mb.Create()
	return m
}
