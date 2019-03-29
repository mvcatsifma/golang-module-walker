package db

import (
	"database/sql"
	"fmt"
	"github.com/mvcatsifma/golang-module-walker/core"
)

type IDatabase interface {
	GetDB() (*sql.DB, error)
	Open() (*sql.DB, error)
	Close() error
}

type api struct {
	core.Api
}

func (a *api) GetDB() (*sql.DB, error) {
	fmt.Println("GetDB")
	return nil, nil
}

func (*api) Open() (*sql.DB, error) {
	fmt.Println("Open")
	return nil, nil
}

func (*api) Close() error {
	fmt.Println("Close")
	return nil
}

func NewApi() *api {
	return &api{}
}