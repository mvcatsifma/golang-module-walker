//+build wireinject

package a

import (
	"github.com/google/wire"
	"github.com/mvcatsifma/golang-module-walker/db"
)

func BuildModule(database db.IDatabase) *module {
	wire.Build(NewApi, NewA)
	return &module{}
}
