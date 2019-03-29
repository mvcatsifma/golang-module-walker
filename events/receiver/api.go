package receiver

import "github.com/mvcatsifma/golang-module-walker/core"

type api struct {
	core.Api
}

func NewApi() *api {
	links := []core.Link{
		{"rel3": "/c/a"},
		{"rel4": "/c/a"},
	}
	return &api{
		Api: core.MakeApi(links),
	}
}
