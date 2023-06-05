package main

import (
	"github.com/anthdm/hollywood/actor"
	"github.com/thankala/diploma-thesis/common/network"
)

type task struct {
}

func NewTask() actor.Producer {
	return func() actor.Receiver {
		return &task{}
	}
}

func (t *task) Receive(ctx *actor.Context) {
	switch msg := ctx.Message().(type) {
	case actor.Initialized:
		_ = msg
	case actor.Started:
	case *network.Start:
	case *network.Stop:

	}
}
