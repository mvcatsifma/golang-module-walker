package nats

import (
	"fmt"
	"github.com/mvcatsifma/golang-module-walker/core"
)

type api struct {
	core.Api
}

func (a api) Subscribe() error {
	fmt.Println("Subscribe")
	return nil
}

func NewApi() *api {
	return &api{}
}

type ISubsciber interface {
	Subscribe() error
}
