package root

import (
	"github.com/mvcatsifma/golang-module-walker/a"
	"github.com/mvcatsifma/golang-module-walker/db"
	"github.com/mvcatsifma/golang-module-walker/core"
	"github.com/mvcatsifma/golang-module-walker/nats"
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
	database := db.BuildModule()
	broker := nats.BuildModule()
	m.Children = append(m.Children, database)
	m.Children = append(m.Children, broker)
	m.Children = append(m.Children, a.BuildModule(database.Api, broker.Api))
	return m
}
