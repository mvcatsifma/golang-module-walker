package a

import (
	"fmt"
	"github.com/mvcatsifma/golang-module-walker/a/c"
	"github.com/mvcatsifma/golang-module-walker/core"
	"log"
)

type module struct {
	Api     *api
	modules []core.IModule
}

func NewModule(api *api) *module {
	m := &module{Api: api}
	m.modules = append(m.modules, c.BuildModule())
	return m
}

func (this *module) Start() error {
	fmt.Println("a - start")

	for _, m := range this.modules {
		err := m.Start()
		if err != nil {
			log.Fatal(err)
		}
	}

	return nil
}

func (this *module) Terminate() error {
	fmt.Println("a - terminate")

	for _, m := range this.modules {
		err := m.Terminate()
		if err != nil {
			log.Println(err)
		}
	}

	return nil
}
