package wpm

import (
	"embed"

	"github.com/bitwormhole/wpm/gen/wpmappcfg"

	ginstarter "github.com/bitwormhole/starter-gin"
	gormstarter "github.com/bitwormhole/starter-gorm"
	sqlserverd "github.com/bitwormhole/starter-gorm-sqlserver"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
)

const (
	theModuleName     = "github.com/bitwormhole/wpm"
	theModuleVersion  = "v0.0.1"
	theModuleRevision = 1
)

//go:embed "src/main/resources"
var theModuleMainResFS embed.FS

// Module 定义模块:wpm
func Module() application.Module {

	mb := &application.ModuleBuilder{}
	mb.Name(theModuleName).Version(theModuleVersion).Revision(theModuleRevision)
	mb.OnMount(wpmappcfg.ExportConfig)
	mb.Resources(collection.LoadEmbedResources(&theModuleMainResFS, "src/main/resources"))

	mb.Dependency(ginstarter.Module())
	mb.Dependency(ginstarter.ModuleWithDevtools())

	mb.Dependency(gormstarter.Module())
	mb.Dependency(sqlserverd.DriverModule())

	return mb.Create()
}
