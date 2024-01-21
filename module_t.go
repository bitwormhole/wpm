package wpm

import (
	"embed"

	"github.com/starter-go/application"
)

const (
	theModuleName     = "github.com/bitwormhole/wpm"
	theModuleVersion  = "v1.0.0"
	theModuleRevision = 29
)

////////////////////////////////////////////////////////////////////////////////

const (
	theMainModuleResPath   = "src/main/resources"
	theAgentModuleResPath  = "src/agent/resources"
	theBootModuleResPath   = "src/boot/resources"
	theCliModuleResPath    = "src/cli/resources"
	theCommonModuleResPath = "src/common/resources"
	theGuiModuleResPath    = "src/gui/resources"
	theServerModuleResPath = "src/server/resources"
	theTestModuleResPath   = "src/test/resources"
)

//go:embed "src/main/resources"
var theMainModuleResFS embed.FS

//go:embed "src/agent/resources"
var theAgentModuleResFS embed.FS

//go:embed "src/boot/resources"
var theBootModuleResFS embed.FS

//go:embed "src/common/resources"
var theCommonModuleResFS embed.FS

//go:embed "src/server/resources"
var theServerModuleResFS embed.FS

//go:embed "src/test/resources"
var theTestModuleResFS embed.FS

//go:embed "src/cli/resources"
var theCliModuleResFS embed.FS

//go:embed "src/gui/resources"
var theGuiModuleResFS embed.FS

////////////////////////////////////////////////////////////////////////////////

// NewMainModule 主模块：默认情况下，调用 server-boot & gui
func NewMainModule() *application.ModuleBuilder {
	mb := new(application.ModuleBuilder)
	mb.Name(theModuleName + "#main")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)
	mb.EmbedResources(theMainModuleResFS, theMainModuleResPath)
	return mb
}

// NewAgentModule Agent 模块：向命令行输出 about 信息
func NewAgentModule() *application.ModuleBuilder {
	mb := new(application.ModuleBuilder)
	mb.Name(theModuleName + "#agent")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)
	mb.EmbedResources(theAgentModuleResFS, theAgentModuleResPath)
	return mb
}

// NewBootModule Boot 模块：先检查是否需要启动服务器，如需要就启动
func NewBootModule() *application.ModuleBuilder {
	mb := new(application.ModuleBuilder)
	mb.Name(theModuleName + "#boot")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)
	mb.EmbedResources(theBootModuleResFS, theBootModuleResPath)
	return mb
}

// NewCommonModule  Common 模块：提供一些通用的组件
func NewCommonModule() *application.ModuleBuilder {
	mb := new(application.ModuleBuilder)
	mb.Name(theModuleName + "#common")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)
	mb.EmbedResources(theCommonModuleResFS, theCommonModuleResPath)
	return mb
}

// NewServerModule 服务器模块：简单而纯粹地启动服务器
func NewServerModule() *application.ModuleBuilder {
	mb := new(application.ModuleBuilder)
	mb.Name(theModuleName + "#server")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)
	mb.EmbedResources(theServerModuleResFS, theServerModuleResPath)
	return mb
}

// NewGuiModule  GUI 模块：调用浏览器，显示用户界面
func NewGuiModule() *application.ModuleBuilder {
	mb := new(application.ModuleBuilder)
	mb.Name(theModuleName + "#gui")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)
	mb.EmbedResources(theGuiModuleResFS, theGuiModuleResPath)
	return mb
}

// NewCliModule  CLI 模块：提供命令行用户界面
func NewCliModule() *application.ModuleBuilder {
	mb := new(application.ModuleBuilder)
	mb.Name(theModuleName + "#cli")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)
	mb.EmbedResources(theCliModuleResFS, theCliModuleResPath)
	return mb
}

// NewTestModule  测试模块: 执行自动化测试
func NewTestModule() *application.ModuleBuilder {
	mb := new(application.ModuleBuilder)
	mb.Name(theModuleName + "#test")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)
	mb.EmbedResources(theTestModuleResFS, theTestModuleResPath)
	return mb
}
