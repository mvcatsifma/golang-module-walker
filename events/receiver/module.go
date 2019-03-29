package receiver

import (
	"github.com/mvcatsifma/golang-module-walker/core"
	"github.com/mvcatsifma/golang-module-walker/db"
)

type module struct {
	core.Module
	db db.IDatabase
}

func NewC(api *api, db db.IDatabase) *module {
	m := &module{
		Module: core.MakeModule("Receiver", api),
		db:     db,
	}
	m.Children = make([]core.IModule, 0)
	return m
}
