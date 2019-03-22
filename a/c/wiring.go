//+build wireinject

package c

import (
	"github.com/google/wire"
)

func BuildModule() *module {
	wire.Build(NewApi, NewC)
	return &module{}
}
