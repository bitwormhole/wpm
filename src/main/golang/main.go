package main

import (
	"github.com/bitwormhole/starter"
	"github.com/bitwormhole/wpm"
)

func main() {
	i := starter.InitApp()
	i.Use(wpm.ServerModule())
	i.Run()
}
