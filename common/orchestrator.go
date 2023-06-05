package common

import (
	"github.com/anthdm/hollywood/actor"
	"github.com/sirupsen/logrus"
	"github.com/thankala/diploma-thesis/common/components"
	"github.com/thankala/diploma-thesis/common/network"
	"github.com/thankala/diploma-thesis/common/stores"
)

type Orchestrator struct {
	storer stores.Storer
	locker components.Locker
	server network.Server
}

func NewOrchestrator(storer stores.Storer, locker components.Locker, server network.Server) actor.Producer {
	return func() actor.Receiver {
		return &Orchestrator{
			storer: storer,
			locker: locker,
			server: server,
		}
	}
}

func (o *Orchestrator) Receive(ctx *actor.Context) {
	switch msg := ctx.Message().(type) {
	case actor.Initialized:
		o.server.Initialise(ctx)
	case actor.Started:
		go o.server.Accept(ctx)
	case *network.Message:
		o.server.Send(ctx, *msg)
	case *network.Start:

	case *network.Stop:
	default:
		logrus.WithFields(map[string]interface{}{
			"Task":    ctx.PID().String(),
			"Message": msg,
		}).Warning("[ORCHESTRATOR] Unknown message received")
	}

}
