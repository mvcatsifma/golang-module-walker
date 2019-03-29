package a

import (
	"github.com/mvcatsifma/golang-module-walker/a/c"
	"github.com/mvcatsifma/golang-module-walker/core"
	"github.com/mvcatsifma/golang-module-walker/db"
	"github.com/mvcatsifma/golang-module-walker/nats"
)

type module struct {
	core.Module
	db     db.IDatabase
	broker nats.IConnection
}

func NewA(api *api, db db.IDatabase, broker nats.IConnection) *module {
	m := &module{
		Module: core.MakeModule("A", api),
		db:     db,
		broker: broker,
	}
	m.Children = append(m.Children, c.BuildModule(db))
	return m
}
