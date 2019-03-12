package main

import (
	"github.com/mvcatsifma/golang-module-walker/a"
	"github.com/mvcatsifma/golang-module-walker/b"
	"github.com/mvcatsifma/golang-module-walker/core"
	"log"
	"os"
	"os/signal"
)

func main() {
	var modules []core.IModule

	modules = append(modules, a.BuildModule())
	modules = append(modules, b.BuildModule())

	// start modules
	for _, m := range modules {
		err := m.Start()
		if err != nil {
			log.Fatal(err)
		}
	}

	// create a channel to receive incoming OS interrupts (such as Ctrl-C):
	osInterruptChannel := make(chan os.Signal, 1)
	signal.Notify(osInterruptChannel, os.Interrupt)

	// block execution until an OS signal (such as Ctrl-C) is received:
	<-osInterruptChannel

	// terminate  modules
	for _, m := range modules {
		err := m.Terminate()
		if err != nil {
			log.Println(err)
		}
	}
}
