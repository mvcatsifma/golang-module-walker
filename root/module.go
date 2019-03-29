package root

import (
	"github.com/mvcatsifma/golang-module-walker/events"
	"github.com/mvcatsifma/golang-module-walker/config"
	"github.com/mvcatsifma/golang-module-walker/db"
	"github.com/mvcatsifma/golang-module-walker/core"
	"github.com/mvcatsifma/golang-module-walker/logging"
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
	conf := config.BuildModule()
	var _ = logging.BuildModule(conf.Config) // todo: use in modules
	database := db.BuildModule(conf.Config)
	broker := nats.BuildModule(conf.Config)
	m.Children = append(m.Children, conf)
	m.Children = append(m.Children, database)
	m.Children = append(m.Children, broker)
	m.Children = append(m.Children, events.BuildModule(database.Api, broker.Api))
	return m
}
