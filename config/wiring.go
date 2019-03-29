//+build wireinject

package config

import (
	"github.com/google/wire"
)

func BuildModule() *module {
	wire.Build(NewApi, NewConfig)
	return &module{}
}
