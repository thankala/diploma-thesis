package network

import (
	"github.com/anthdm/hollywood/actor"
)

type Server interface {
	Initialise(ctx *actor.Context)
	Accept(ctx *actor.Context)
	Send(ctx *actor.Context, msg Message)
}
