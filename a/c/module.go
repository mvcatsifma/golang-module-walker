package c

import (
	"github.com/mvcatsifma/golang-module-walker/core"
)

type module struct {
	core.Module
	Api *api
}

func NewModule(api *api) *module {
	return &module{
		Module: core.Module{Name: "C"},
		Api:    api,
	}
}
