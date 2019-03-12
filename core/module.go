package core

import (
	"fmt"
	"log"
)

type IModule interface {
	Start() error
	Terminate() error
}

type Module struct {
	Name      string
	Modules   []IModule
}

func (this *Module) Start() error {
	fmt.Printf("Starting module %v\n", this.Name)

	for _, m := range this.Modules {
		err := m.Start()
		if err != nil {
			log.Fatal(err)
		}
	}

	return nil
}

func (this *Module) Terminate() error {
	fmt.Printf("Terminating module %v\n", this.Name)

	for _, m := range this.Modules {
		err := m.Terminate()
		if err != nil {
			log.Println(err)
		}
	}

	return nil
}
