package c

import (
	"github.com/mvcatsifma/golang-module-walker/core"
)

type module struct {
	core.Module
	Api *api
}

func NewModule(api *api) *module {
	m := &module{
		Module: core.Module{Name: "C"},
		Api:    api,
	}
	m.Modules = make([]core.IModule, 0)
	return m
}
