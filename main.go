package main

import (
	"fmt"
	"github.com/mvcatsifma/golang-module-walker/core"
	"github.com/mvcatsifma/golang-module-walker/root"
	"log"
	"os"
	"os/signal"
)

func main() {
	r := root.BuildModule()
	err := r.Start()
	if err != nil {
		log.Fatal(err)
	}

	var links []core.Link
	core.Inspect(r, func(module core.IModule) bool {
		if module == nil {
			return false
		}
		fmt.Printf("inspecting module: %v \n", module.GetName())
		for _, l := range module.GetRootLinks() {
			links = append(links, l)
		}
		return true
	})

	for _, l := range links {
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
