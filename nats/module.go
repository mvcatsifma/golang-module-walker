package nats

import (
	"github.com/mvcatsifma/golang-module-walker/core"
)

type module struct {
	core.Module
	Api *api
}

func NewDB(api *api) *module {
	m := &module{
		Module: core.MakeModule("NATS", api),
	}
	return m
}
