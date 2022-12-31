package main

import (
	"github.com/bitwormhole/starter"
	"github.com/bitwormhole/wpm"
)

func main() {

	// fmt.Println("src/main/go")

	i := starter.InitApp()
	i.Use(wpm.Module())

	i.Run()
}
