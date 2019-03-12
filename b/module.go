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
	m.Modules = make([]core.IModule, 0)
	// root api links
	m.RootLinks = append(m.RootLinks, core.Link{"rel1": "/b/a"})
	m.RootLinks = append(m.RootLinks, core.Link{"rel2": "/b/a"})
	return m
}
