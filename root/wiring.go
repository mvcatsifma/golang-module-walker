//+build wireinject

package root

import (
	"github.com/google/wire"
)

func BuildModule() *Root {
	wire.Build(NewApi, NewRoot)
	return &Root{}
}
