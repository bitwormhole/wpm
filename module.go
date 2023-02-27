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
	"github.com/bitwormhole/wpm/gen/wpmserver"
	"github.com/bitwormhole/wpm/server/service"
)

const (
	theModuleName     = "github.com/bitwormhole/wpm"
	theModuleVersion  = "v0.0.13"
	theModuleRevision = 13
	theModuleResPath  = "src/main/resources"
)

//go:embed "src/main/resources"
var theModuleResFS embed.FS

// Module 定义模块:wpm
func Module() application.Module {

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
	service.SetAppModule(m)
	return m
}
