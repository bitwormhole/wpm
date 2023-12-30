package main

import (
	"os"

	"github.com/bitwormhole/wpm/modules/wpm"
	"github.com/starter-go/starter"
)

func main() {
	m := wpm.AgentModule()
	i := starter.Init(os.Args)
	i.MainModule(m)
	i.WithPanic(true).Run()
}
