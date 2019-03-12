package root

import (
	"github.com/mvcatsifma/golang-module-walker/a"
	"github.com/mvcatsifma/golang-module-walker/b"
	"github.com/mvcatsifma/golang-module-walker/core"
)

type module struct {
	*core.Module
	Api *api
}

func NewModule(api *api) *module {
	m := &module{
		Module: &core.Module{Name: "Root"},
		Api:    api,
	}
	// sub modules
	m.Modules = append(m.Modules, a.BuildModule())
	m.Modules = append(m.Modules, b.BuildModule())
	// root api links
	m.RootLinks = append(m.RootLinks, core.Link{"rel1": "/root/a"})
	m.RootLinks = append(m.RootLinks, core.Link{"rel2": "/root/a"})
	return m
}
