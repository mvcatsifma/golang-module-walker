package nats

import (
	"github.com/mvcatsifma/golang-module-walker/config"
	"github.com/mvcatsifma/golang-module-walker/core"
)

type module struct {
	core.Module
	Api *api
}

func NewDB(api *api, config *config.AppConfig) *module {
	m := &module{
		Module: core.MakeModule("NATS", api),
	}
	return m
}
