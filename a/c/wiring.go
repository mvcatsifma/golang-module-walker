//+build wireinject

package c

import (
	"github.com/google/wire"
)

func BuildModule() *module {
	wire.Build(NewApi, NewModule)
	return &module{}
}
