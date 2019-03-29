package db

import (
	"fmt"
	"github.com/mvcatsifma/golang-module-walker/core"
)

type api struct {
	core.Api
}

func (*api) Open() error {
	fmt.Println("Open")
	return nil
}

func (*api) Close() error {
	fmt.Println("Close")
	return nil
}

func NewApi() *api {
	return &api{}
}

type IDatabase interface {
	Open() error
	Close() error
}
