package root

import "github.com/mvcatsifma/golang-module-walker/core"

type Api struct {
	core.Api
}

func (this *Api) Add(link core.Link) {
	this.RootLinks = append(this.RootLinks, link)
}

func NewApi() *Api {
	links := []core.Link{
		{"rel1": "/root/a"},
		{"rel2": "/root/a"},
	}
	return &Api{
		Api: core.MakeApi(links),
	}
}
