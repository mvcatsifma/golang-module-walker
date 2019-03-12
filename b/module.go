package b

import "fmt"

type module struct {
	Api *api
}

func NewModule(api *api) *module {
	return &module{Api: api}
}

func (this *module) Start() error {
	fmt.Println("b - start")
	return nil
}

func (this *module) Terminate() error {
	fmt.Println("b - terminate")
	return nil
}
