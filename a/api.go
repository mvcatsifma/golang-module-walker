package a

import (
	"github.com/mvcatsifma/golang-module-walker/core"
	"github.com/mvcatsifma/golang-module-walker/db"
)

type api struct {
	core.Api
	db db.IDatabase
}

func NewApi(db db.IDatabase) *api {
	return &api{
		db: db,
	}
}
