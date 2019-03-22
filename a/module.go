package a

import (
	"github.com/mvcatsifma/golang-module-walker/a/c"
	"github.com/mvcatsifma/golang-module-walker/core"
)

type module struct {
	core.Module
	Api     *api
}

func NewA(api *api) *module {
	m := &module{
		Module: core.MakeModule("A", api),
		Api:    api,
	}
	m.Children = append(m.Children, c.BuildModule())
	return m
}
