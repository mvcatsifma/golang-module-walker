package config

import (
	"github.com/mvcatsifma/golang-module-walker/core"
)

type AppConfig struct {
	Foo string
}

type module struct {
	core.Module
	Api    *api
	Config *AppConfig
}

func NewConfig(api *api) *module {
	m := &module{
		Module: core.MakeModule("Config", api),
		Config: &AppConfig{
			Foo: "bar",
		},
	}
	return m
}
