package c

import (
	"github.com/mvcatsifma/golang-module-walker/core"
)

type module struct {
	core.Module
	Api *api
}

func NewC(api *api) *module {
	m := &module{
		Module: core.MakeModule("C", api),
		Api:    api,
	}
	m.Children = make([]core.IModule, 0)
	return m
}
