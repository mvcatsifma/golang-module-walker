package events

import (
	"fmt"
	"github.com/mvcatsifma/golang-module-walker/core"
	"github.com/mvcatsifma/golang-module-walker/db"
	"github.com/mvcatsifma/golang-module-walker/events/receiver"
	"github.com/mvcatsifma/golang-module-walker/nats"
)

type module struct {
	core.Module
	db     db.IDatabase
	broker nats.IConnection
}

func (m *module) Start() error {
	err := m.Module.Start()
	fmt.Printf("Starting module %v (Custom)\n", m.Name)
	return err
}

func NewA(api *api, db db.IDatabase, broker nats.IConnection) *module {
	m := &module{
		Module: core.MakeModule("Events", api),
		db:     db,
		broker: broker,
	}
	m.Children = append(m.Children, receiver.BuildModule(db))
	return m
}
