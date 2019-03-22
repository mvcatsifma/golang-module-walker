//+build wireinject

package db

import (
	"github.com/google/wire"
)

func BuildModule() *module {
	wire.Build(NewApi, NewDB)
	return &module{}
}
