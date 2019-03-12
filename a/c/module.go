package c

import (
	"fmt"
)

type module struct {
	Api *api
}

func NewModule(api *api) *module {
	return &module{Api: api}
}

func (this *module) Start() error {
	fmt.Println("c - start")
	return nil
}

func (this *module) Terminate() error {
	fmt.Println("c - terminate")
	return nil
}
