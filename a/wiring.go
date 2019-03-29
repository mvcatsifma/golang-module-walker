//+build wireinject

package a

import (
	"github.com/google/wire"
	"github.com/mvcatsifma/golang-module-walker/db"
	"github.com/mvcatsifma/golang-module-walker/nats"
)

func BuildModule(database db.IDatabase, broker nats.ISubsciber) *module {
	wire.Build(NewApi, NewA)
	return &module{}
}
