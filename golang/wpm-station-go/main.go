package main

import (
	gitstarter "github.com/bitwormhole/gitlib/gitstarter"
	ginstarter "github.com/bitwormhole/starter-gin"
)

func main() {
	i := ginstarter.InitGin()
	i.Use(ginstarter.ModuleWithDevtools())
	i.Use(myModule())
	i.Use(gitstarter.Module())
	i.Run()
}
