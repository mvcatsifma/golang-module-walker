package root

import "github.com/mvcatsifma/golang-module-walker/core"

type api struct {
	core.Api
}

func (this *api) Add(link core.Link) {
	this.Links = append(this.Links, link)
}

func NewApi() *api {
	links := []core.Link{
		{"rel1": "/root/a"},
		{"rel2": "/root/a"},
	}
	return &api{
		Api: core.MakeApi(links),
	}
}
