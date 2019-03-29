package a

import (
	"github.com/mvcatsifma/golang-module-walker/core"
)

type api struct {
	core.Api
}

func NewApi() *api {
	return &api{
	}
}
