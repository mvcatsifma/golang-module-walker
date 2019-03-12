package core

import (
	"fmt"
	"log"
)

type IModule interface {
	Start() error
	Terminate() error
	GetName() string
	GetChildren() []IModule
	GetRootLinks() []Link
}

type Module struct {
	Name      string
	Modules   []IModule
	RootLinks []Link
}

func (this *Module) GetName() string {
	return this.Name
}

func (this *Module) GetChildren() []IModule {
	return this.Modules
}

func (this *Module) GetRootLinks() []Link {
	return this.RootLinks
}

func (this *Module) Start() error {
	for _, m := range this.Modules {
		err := m.Start()
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Printf("Starting module %v\n", this.Name)

	return nil
}

func (this *Module) Terminate() error {
	for _, m := range this.Modules {
		err := m.Terminate()
		if err != nil {
			log.Println(err)
		}
	}

	fmt.Printf("Terminating module %v\n", this.Name)

	return nil
}
