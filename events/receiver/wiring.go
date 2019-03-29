//+build wireinject

package receiver

import (
	"github.com/google/wire"
	"github.com/mvcatsifma/golang-module-walker/db"
)

func BuildModule(db db.IDatabase) *module {
	wire.Build(NewApi, NewC)
	return &module{}
}
