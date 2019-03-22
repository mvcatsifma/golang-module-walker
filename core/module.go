package core

import (
	"fmt"
	"log"
)

type IModule interface {
	Start() error
	Terminate() error
	GetName() string
	GetApi() IApi
	GetChildren() []IModule
}

type Module struct {
	Name     string
	Children []IModule
	Api      IApi
}

func MakeModule(name string, api IApi) Module {
	return Module{
		Name: name,
		Api:  api,
	}
}

func (this *Module) GetApi() IApi {
	return this.Api
}

func (this *Module) GetName() string {
	return this.Name
}

func (this *Module) GetChildren() []IModule {
	return this.Children
}

func (this *Module) Start() error {
	for _, m := range this.Children {
		err := m.Start()
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Printf("Starting module %v\n", this.Name)

	return nil
}

func (this *Module) Terminate() error {
	for _, m := range this.Children {
		err := m.Terminate()
		if err != nil {
			log.Println(err)
		}
	}

	fmt.Printf("Terminating module %v\n", this.Name)

	return nil
}
