package nats

import (
	"fmt"
	"github.com/mvcatsifma/golang-module-walker/core"
	"github.com/nats-io/go-nats"
)

type ISubscription interface {
	Unsubscribe() error
}

type IConnection interface {
	Close()
	Subscribe(subject string, cb nats.MsgHandler) (ISubscription, error)
}

type api struct {
	core.Api
}

func (a api) Subscribe(subject string, cb nats.MsgHandler) (ISubscription, error) {
	fmt.Println("Subscribe")
	return nil, nil
}

func (a api) Close() {}

func NewApi() *api {
	return &api{}
}
