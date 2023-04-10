package main

import (
	"github.com/bitwormhole/starter"
	"github.com/bitwormhole/wpm"
	"github.com/bitwormhole/wpm/server/service"
)

func main() {

	runModule1()
	runModule2()
}

func runModule1() {
	i := starter.InitApp()
	i.Use(wpm.ModuleBoot())
	rt, err := i.RunEx()
	if err != nil {
		panic(err)
	}
	err = rt.Loop()
	if err != nil {
		panic(err)
	}
}

func runModule2() {

	m := wpm.ModuleServer()
	service.SetAppModule(m)

	i := starter.InitApp()
	i.Use(m)
	i.Run()
}
