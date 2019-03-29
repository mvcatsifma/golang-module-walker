package root

import (
	"github.com/mvcatsifma/golang-module-walker/a"
	"github.com/mvcatsifma/golang-module-walker/db"
	"github.com/mvcatsifma/golang-module-walker/core"
)

type Root struct {
	core.Module
	Links []core.Link
}

func NewRoot(api *api) *Root {
	m := &Root{
		Module: core.MakeModule("Root", api),
	}
	// sub modules
	m.Children = append(m.Children, a.BuildModule(nil))
	m.Children = append(m.Children, db.BuildModule())
	return m
}
