package main

import (
	"fmt"

	ginstarter "github.com/bitwormhole/starter-gin"
	"github.com/bitwormhole/wpm"
)

func main() {

	fmt.Println("src/main/go")

	i := ginstarter.InitGin()
	i.Use(ginstarter.ModuleWithDevtools())
	i.Use(wpm.Module())
	i.Run()
}
