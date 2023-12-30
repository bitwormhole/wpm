package wpm

import (
	"embed"

	"github.com/starter-go/application"
)

const (
	theModuleName     = "github.com/bitwormhole/wpm"
	theModuleVersion  = "v1.0.0"
	theModuleRevision = 1
)

////////////////////////////////////////////////////////////////////////////////

const (
	theMainModuleResPath   = "src/main/resources"
	theAgentModuleResPath  = "src/agent/resources"
	theServerModuleResPath = "src/server/resources"
	theTestModuleResPath   = "src/test/resources"
)

//go:embed "src/main/resources"
var theMainModuleResFS embed.FS

//go:embed "src/agent/resources"
var theAgentModuleResFS embed.FS

//go:embed "src/server/resources"
var theServerModuleResFS embed.FS

//go:embed "src/test/resources"
var theTestModuleResFS embed.FS

////////////////////////////////////////////////////////////////////////////////

// NewMainModule ...
func NewMainModule() *application.ModuleBuilder {
	mb := new(application.ModuleBuilder)
	mb.Name(theModuleName + "#main")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)
	mb.EmbedResources(theMainModuleResFS, theMainModuleResPath)
	return mb
}

// NewAgentModule ...
func NewAgentModule() *application.ModuleBuilder {
	mb := new(application.ModuleBuilder)
	mb.Name(theModuleName + "#agent")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)
	mb.EmbedResources(theAgentModuleResFS, theAgentModuleResPath)
	return mb
}

// NewServerModule ...
func NewServerModule() *application.ModuleBuilder {
	mb := new(application.ModuleBuilder)
	mb.Name(theModuleName + "#server")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)
	mb.EmbedResources(theServerModuleResFS, theServerModuleResPath)
	return mb
}

// NewTestModule ...
func NewTestModule() *application.ModuleBuilder {
	mb := new(application.ModuleBuilder)
	mb.Name(theModuleName + "#test")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)
	mb.EmbedResources(theTestModuleResFS, theTestModuleResPath)
	return mb
}
