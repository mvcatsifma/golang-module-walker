package a

import (
	"github.com/mvcatsifma/golang-module-walker/a/c"
	"github.com/mvcatsifma/golang-module-walker/core"
)

type module struct {
	core.Module
	Api     *api
}

func NewModule(api *api) *module {
	m := &module{
		Module: core.Module{Name: "A"},
		Api:    api,
	}
	m.Modules = append(m.Modules, c.BuildModule())
	return m
}
