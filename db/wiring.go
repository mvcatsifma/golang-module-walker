//+build wireinject

package db

import (
	"github.com/google/wire"
	"github.com/mvcatsifma/golang-module-walker/config"
)

func BuildModule(config *config.AppConfig) *module {
	wire.Build(NewApi, NewDB)
	return &module{}
}
