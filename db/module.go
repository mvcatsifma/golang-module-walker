package db

import (
	"fmt"
	"github.com/mvcatsifma/golang-module-walker/config"
	"github.com/mvcatsifma/golang-module-walker/core"
)

type module struct {
	core.Module
	Api *api
}

func NewDB(api *api, config *config.AppConfig) *module {
	fmt.Printf("foo: %v\n", config.Foo)
	m := &module{
		Module: core.MakeModule("DB", api),
	}
	return m
}
