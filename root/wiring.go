//+build wireinject

package root

import (
	"github.com/google/wire"
)

func BuildModule() *module {
	wire.Build(NewApi, NewModule)
	return &module{}
}
