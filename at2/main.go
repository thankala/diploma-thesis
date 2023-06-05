package main

import (
	"github.com/anthdm/hollywood/actor"
	"github.com/sirupsen/logrus"
	"github.com/thankala/diploma-thesis/common"
	"github.com/thankala/diploma-thesis/common/middlewares"
	"github.com/thankala/diploma-thesis/common/network"
	"github.com/thankala/diploma-thesis/common/stores"
)

func main() {
	var (
		store  = stores.NewRedisStore()
		server = network.NewKafkaServer(network.WithTopic("AT2"))
		e      = actor.NewEngine()
	)

	logrus.SetFormatter(&logrus.JSONFormatter{})

	e.Spawn(
		common.NewOrchestrator(store, store, server),
		"AT2",
		actor.WithMiddleware(middlewares.WithPersistence(store)),
	)

	<-make(chan struct{})
}
