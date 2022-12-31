package main

import (
	"os"

	srcmain "bitwormhole.com/wpm/wpm-windows-boot/src/main"
	"github.com/bitwormhole/starter"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/vlog"
)

func main() {

	mb := &application.ModuleBuilder{}
	mb.Name("wpm-windows-boot").Version("v1").Revision(1)
	mb.Resources(srcmain.ExportResources())
	mb.Dependency(starter.Module())

	looperTimes := os.Getenv("looper.times")
	looperStep := os.Getenv("looper.step")
	vlog.Debug("looper.times: ", looperTimes)
	vlog.Debug("looper.step:  ", looperStep)

	starter.InitApp().Use(mb.Create()).Run()
}
