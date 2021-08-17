package main

import (
	ginstarter "github.com/bitwormhole/starter-gin"
	"github.com/bitwormhole/starter/application"
	etcwpm "github.com/bitwormhole/wpm/etc/wpm"
	wpmsrcmain "github.com/bitwormhole/wpm/src/main"
)

const (
	myName     = "github.com/bitwormhole/wpm"
	myVersion  = "v0.0.1"
	myRevision = 1
)

// Module 定义模块:wpm
func Module() application.Module {
	return myModule()
}

func myModule() application.Module {

	mod := &application.DefineModule{
		Name:     myName,
		Version:  myVersion,
		Revision: myRevision,
	}

	mod.OnMount = func(cb application.ConfigBuilder) error { return etcwpm.ExportConfig(cb) }
	mod.Resources = wpmsrcmain.ExportResources()
	mod.AddDependency(ginstarter.Module())
	mod.AddDependency(ginstarter.ModuleWithDevtools())

	return mod
}
