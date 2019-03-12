package b

import (
	"github.com/mvcatsifma/golang-module-walker/core"
)

type module struct {
	core.Module
	Api     *api
}

func NewModule(api *api) *module {
	m := &module{
		Module: core.Module{Name: "B"},
		Api:    api,
	}
	return m
}
