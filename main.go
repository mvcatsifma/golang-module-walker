package main

import (
	"fmt"
	"github.com/mvcatsifma/golang-module-walker/core"
	"github.com/mvcatsifma/golang-module-walker/root"
	"log"
	"os"
	"os/signal"
)

// todo:
// - links should be in Api
// - register module interface for lookup
func main() {
	r := root.BuildModule()
	err := r.Start()
	if err != nil {
		log.Fatal(err)
	}

	core.Inspect(r, func(module core.IModule) bool {
		if module == nil {
			return false
		}

		fmt.Printf("inspecting module: %v \n", module.GetName())

		switch module.(type) {
		case *root.Root:
		default:
			for _, l := range module.GetApi().GetLinks() {
				r.Api.(core.LinkAdder).Add(l)
			}
		}
		return true
	})

	for _, l := range r.Api.GetLinks() {
		fmt.Println(l)
	}

	// create a channel to receive incoming OS interrupts (such as Ctrl-C):
	osInterruptChannel := make(chan os.Signal, 1)
	signal.Notify(osInterruptChannel, os.Interrupt)

	// block execution until an OS signal (such as Ctrl-C) is received:
	<-osInterruptChannel

	// terminate  modules
	err = r.Terminate()
	if err != nil {
		log.Println(err)
	}
}
