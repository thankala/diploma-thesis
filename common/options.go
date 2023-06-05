package common

import (
	"github.com/google/uuid"
	"strconv"
	"strings"
)

var (
	defaultId                   = uuid.New()
	defaultTaskName             = "AT1"
	defaultKafkaServerAddresses = []string{"localhost:9094"}
	defaultKafkaTopic           = "AT1"
	defaultKafkaGroupId         = "AT1"
	defaultKafkaPartition       = 0
	defaultStore                = "redis"
	defaultStoreAddr            = "localhost:6379"
	defaultStorePassword        = ""
	defaultStoreDb              = 0
)

type Opts struct {
	Id                   uuid.UUID
	TaskName             string
	KafkaServerAddresses []string
	KafkaTopic           string
	KafkaGroupId         string
	KafkaPartition       int
	Store                string
	StoreAddr            string
	StoreType            string
	StorePassword        string
	StoreDb              int
}

type OptFunc func(*Opts)

func NewOpts(opts ...OptFunc) *Opts {
	options := DefaultOpts()
	for _, opt := range opts {
		opt(options)
	}
	return options
}

func DefaultOpts() *Opts {
	return &Opts{
		defaultId,
		defaultTaskName,
		defaultKafkaServerAddresses,
		defaultKafkaTopic,
		defaultKafkaGroupId,
		defaultKafkaPartition,
		defaultStore,
		defaultStoreAddr,
		defaultTaskName,
		defaultStorePassword,
		defaultStoreDb,
	}
}

func WithId(id uuid.UUID) OptFunc {
	if id == uuid.Nil {
		return func(opts *Opts) {}
	} else {
		return func(opts *Opts) {
			opts.Id = id
		}
	}
}
func WithTaskName(taskName string) OptFunc {
	if taskName == "" {
		return func(opts *Opts) {}
	} else {
		return func(opts *Opts) {
			opts.TaskName = taskName
		}
	}
}

func WithKafkaServerAddresses(serverAddresses string) OptFunc {
	if serverAddresses == "" {
		return func(opts *Opts) {}
	} else {
		return func(opts *Opts) {
			opts.KafkaServerAddresses = strings.Split(serverAddresses, ",")
		}
	}
}

func WithKafkaTopic(topic string) OptFunc {
	if topic == "" {
		return func(opts *Opts) {}
	} else {
		return func(opts *Opts) {
			opts.KafkaTopic = topic
		}
	}
}

func WithKafkaGroupId(groupId string) OptFunc {
	if groupId == "" {
		return func(opts *Opts) {}
	} else {
		return func(opts *Opts) {
			opts.KafkaGroupId = groupId
		}
	}
}

func WithKafkaPartition(partition string) OptFunc {
	if partition == "" {
		return func(opts *Opts) {}
	} else {
		return func(opts *Opts) {
			kafkaPartition, err := strconv.Atoi(partition)
			if err != nil {
				return
			}
			opts.KafkaPartition = kafkaPartition
		}
	}
}

func WithStore(store string) OptFunc {
	if store == "" {
		return func(opts *Opts) {}
	} else {
		return func(opts *Opts) {
			opts.Store = store
		}
	}
}

func WithStoreAddr(storeAddr string) OptFunc {
	if storeAddr == "" {
		return func(opts *Opts) {}
	} else {
		return func(opts *Opts) {
			opts.StoreAddr = storeAddr
		}
	}
}

func WithStorePassword(storePassword string) OptFunc {
	if storePassword == "" {
		return func(opts *Opts) {}
	} else {
		return func(opts *Opts) {
			opts.StorePassword = storePassword
		}
	}
}

func WithStoreDb(storeDb string) OptFunc {
	if storeDb == "" {
		return func(opts *Opts) {}
	} else {
		return func(opts *Opts) {
			db, err := strconv.Atoi(storeDb)
			if err != nil {
				return
			}
			opts.StoreDb = db
		}
	}
}
