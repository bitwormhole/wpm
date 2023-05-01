package wpm

import (
	"embed"

	"github.com/bitwormhole/gitlib"
	"github.com/bitwormhole/starter"

	ginstarter "github.com/bitwormhole/starter-gin"
	gormstarter "github.com/bitwormhole/starter-gorm"
	mysqld "github.com/bitwormhole/starter-gorm-mysql"
	sqlited "github.com/bitwormhole/starter-gorm-sqlite"
	sqlserverd "github.com/bitwormhole/starter-gorm-sqlserver"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/wpm/gen/wpmboot"
	"github.com/bitwormhole/wpm/gen/wpmserver"
)

const (
	theModuleName     = "github.com/bitwormhole/wpm"
	theModuleVersion  = "v0.1.10"
	theModuleRevision = 28
	theModuleResPath  = "src/main/resources"
)

//go:embed "src/main/resources"
var theModuleResFS embed.FS

// ModuleServer 定义模块:wpm-server
func ModuleServer() application.Module {

	mb := &application.ModuleBuilder{}
	mb.Name(theModuleName + "#server")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)

	mb.OnMount(wpmserver.ExportConfig)
	mb.Resources(collection.LoadEmbedResources(&theModuleResFS, theModuleResPath))

	mb.Dependency(ginstarter.Module())
	mb.Dependency(ginstarter.ModuleWithDevtools())
	mb.Dependency(gormstarter.Module())
	mb.Dependency(sqlserverd.DriverModule())
	mb.Dependency(mysqld.DriverModule())
	mb.Dependency(sqlited.DriverModule())
	mb.Dependency(starter.Module())
	mb.Dependency(gitlib.Module())

	m := mb.Create()
	// service.SetAppModule(m)
	return m
}

// ModuleBoot 定义模块:wpm-boot
// 该模块用于预先加载配置，并根据配置启动对应的运行模式
func ModuleBoot() application.Module {

	mb := &application.ModuleBuilder{}
	mb.Name(theModuleName + "#boot")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)

	mb.OnMount(wpmboot.ExportConfig)
	mb.Resources(collection.LoadEmbedResources(&theModuleResFS, theModuleResPath))

	mb.Dependency(starter.Module())

	return mb.Create()
}
