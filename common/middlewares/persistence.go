package middlewares

import (
	"encoding/json"
	"github.com/anthdm/hollywood/actor"
	"github.com/anthdm/hollywood/log"
	"github.com/sirupsen/logrus"
	"github.com/thankala/diploma-thesis/common/stores"
	"reflect"
)

func WithPersistence(store stores.Storer) func(actor.ReceiveFunc) actor.ReceiveFunc {
	return func(next actor.ReceiveFunc) actor.ReceiveFunc {
		return func(ctx *actor.Context) {
			switch ctx.Message().(type) {
			case actor.Initialized:
				logrus.WithFields(map[string]interface{}{
					"Task":  ctx.PID().ID,
					"Store": reflect.TypeOf(store).String(),
				}).Info("[MIDDLEWARE] Initializing Store")
				p, ok := ctx.Receiver().(Persister)
				if !ok {
					logrus.WithFields(map[string]interface{}{
						"Task":  ctx.PID().ID,
						"Store": reflect.TypeOf(store).String(),
					}).Warning("[MIDDLEWARE] ")
					next(ctx)
					return
				}
				b, err := store.Load(ctx.PID().String())
				if err != nil {
					logrus.WithFields(map[string]interface{}{
						"Task":  ctx.PID().ID,
						"Error": err,
						"Store": reflect.TypeOf(store).String(),
					}).Error("Unable to ping store")
					next(ctx)
					return
				}
				var data string
				if err := json.Unmarshal(b, &data); err != nil {
					logrus.WithFields(map[string]interface{}{
						"Task":  ctx.PID().ID,
						"Error": err,
						"Store": reflect.TypeOf(store).String(),
					}).Fatal("Unable to ping store")
					panic(err)
				}
				if err := p.LoadState(data); err != nil {
					log.Errorw("Load state error:", log.M{"Error": err})
					next(ctx)
				}
			case actor.Stopped:
				if p, ok := ctx.Receiver().(Persister); ok {
					s, err := p.State()
					if err != nil {
						log.Fatalw("Failed getting state from struct", log.M{"Error": err})
						next(ctx)
						return
					}
					if err := store.Store(ctx.PID().String(), s); err != nil {
						log.Errorw("Failed to store the state", log.M{"Error": err})
						next(ctx)
						return
					}
				}
			}
			next(ctx)
		}
	}
}
