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
	"github.com/bitwormhole/wpm/gen/wpmclient"
	"github.com/bitwormhole/wpm/gen/wpmcommon"
	"github.com/bitwormhole/wpm/gen/wpmserver"
)

const (
	theModuleName     = "github.com/bitwormhole/wpm"
	theModuleVersion  = "v0.0.2"
	theModuleRevision = 2
	theModuleResPath  = "src/main/resources"
)

//go:embed "src/main/resources"
var theModuleResFS embed.FS

// ServerModule 定义模块:wpm
func ServerModule() application.Module {

	mb := &application.ModuleBuilder{}
	mb.Name(theModuleName + "#server")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)

	mb.OnMount(wpmserver.ExportConfig)
	mb.Resources(collection.LoadEmbedResources(&theModuleResFS, theModuleResPath))

	mb.Dependency(ginstarter.Module())
	mb.Dependency(ginstarter.ModuleWithDevtools())
	mb.Dependency(CommonModule())

	mb.Dependency(gormstarter.Module())
	mb.Dependency(sqlserverd.DriverModule())
	mb.Dependency(mysqld.DriverModule())
	mb.Dependency(sqlited.DriverModule())

	return mb.Create()
}

// ClientModule 定义模块:wpm
func ClientModule() application.Module {

	mb := &application.ModuleBuilder{}
	mb.Name(theModuleName + "#client")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)

	mb.OnMount(wpmclient.ExportConfig)
	mb.Resources(collection.LoadEmbedResources(&theModuleResFS, theModuleResPath))

	mb.Dependency(starter.Module())
	mb.Dependency(CommonModule())

	return mb.Create()
}

// CommonModule 定义模块:wpm
func CommonModule() application.Module {

	mb := &application.ModuleBuilder{}
	mb.Name(theModuleName + "#common")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)

	mb.OnMount(wpmcommon.ExportConfig)
	mb.Resources(collection.LoadEmbedResources(&theModuleResFS, theModuleResPath))

	mb.Dependency(starter.Module())
	mb.Dependency(gitlib.Module())

	return mb.Create()
}
