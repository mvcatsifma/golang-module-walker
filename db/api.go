package db

import (
	"fmt"
	"github.com/mvcatsifma/golang-module-walker/core"
)

type api struct {
	core.Api
}

func (*api) Open() error {
	fmt.Println("Opening database connnection")
	return nil
}

func (*api) Close() error {
	fmt.Println("Closing database")
	return nil
}

func NewApi() *api {
	return &api{}
}

type IDatabase interface {
	Open() error
	Close() error
}
