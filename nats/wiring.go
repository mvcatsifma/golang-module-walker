//+build wireinject

package nats

import (
	"github.com/google/wire"
)

func BuildModule() *module {
	wire.Build(NewApi, NewDB)
	return &module{}
}
