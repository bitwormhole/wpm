package main

import (
	"os"

	"github.com/bitwormhole/starter"
	"github.com/bitwormhole/wpm"
)

func main() {
	err := (&wpmboot{}).run()
	if err != nil {
		panic(err)
	}
}

////////////////////////////////////////////////////////////////////////////////

type wpmboot struct {
}

func (inst *wpmboot) run() error {

	args := os.Args
	isServerMode := inst.hasArgOption("-s", args)
	isClientMode := inst.hasArgOption("-c", args)
	if !isClientMode && !isServerMode {
		isClientMode = true
		isServerMode = true
	}

	if isServerMode {
		return inst.runServer()
	} else if isClientMode {
		return inst.runClient()
	}

	return nil
}

func (inst *wpmboot) hasArgOption(target string, args []string) bool {
	for _, item := range args {
		if item == target {
			return true
		}
	}
	return false
}

func (inst *wpmboot) runClient() error {
	i := starter.InitApp()
	i.Use(wpm.ClientModule())
	i.Run()
	return nil
}

func (inst *wpmboot) runServer() error {
	i := starter.InitApp()
	i.Use(wpm.ServerModule())
	i.Run()
	return nil
}
