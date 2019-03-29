package logging

import (
	"github.com/mvcatsifma/golang-module-walker/config"
	"github.com/mvcatsifma/golang-module-walker/core"
)

type module struct {
	core.Module
	Api *api
}

func NewConfig(api *api, config *config.AppConfig) *module {
	m := &module{
		Module: core.MakeModule("Logging", api),
	}
	return m
}
