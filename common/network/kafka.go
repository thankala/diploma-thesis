package network

import (
	"context"
	"encoding/json"
	"github.com/anthdm/hollywood/actor"
	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

var (
	defaultServerAddresses = []string{"localhost:9094"}
	defaultTopic           = "AT1"
	defaultGroupId         = "AT1"
	defaultPartition       = 0
)

type KafkaOptFunc func(opts *kafkaOpts)

type kafkaOpts struct {
	listenAddresses []string
	topic           string
	groupId         string
	partition       int
}

func defaultOps() *kafkaOpts {
	return &kafkaOpts{
		defaultServerAddresses,
		defaultTopic,
		defaultGroupId,
		defaultPartition,
	}
}

type KafkaServer struct {
	opts   *kafkaOpts
	reader *kafka.Reader
	writer *kafka.Writer
}

func NewKafkaServer(opts ...KafkaOptFunc) *KafkaServer {
	options := defaultOps()
	for _, opt := range opts {
		opt(options)
	}
	return &KafkaServer{
		opts: options,
	}
}

func WithServerAddresses(serverAddresses string) KafkaOptFunc {
	if serverAddresses == "" {
		return func(opts *kafkaOpts) {}
	} else {
		return func(opts *kafkaOpts) {
			opts.listenAddresses = strings.Split(serverAddresses, ",")
		}
	}
}

func WithTopic(topic string) KafkaOptFunc {
	if topic == "" {
		return func(opts *kafkaOpts) {}
	} else {
		return func(opts *kafkaOpts) {
			opts.topic = topic
		}
	}
}

func WithGroupId(groupId string) KafkaOptFunc {
	if groupId == "" {
		return func(opts *kafkaOpts) {}
	} else {
		return func(opts *kafkaOpts) {
			opts.groupId = groupId
		}
	}
}

func WithPartition(partition string) KafkaOptFunc {
	if partition == "" {
		return func(opts *kafkaOpts) {}
	} else {
		return func(opts *kafkaOpts) {
			kafkaPartition, err := strconv.Atoi(partition)
			if err != nil {
				return
			}
			opts.partition = kafkaPartition
		}
	}
}

func (k *KafkaServer) Initialise(ctx *actor.Context) {
	logrus.WithFields(map[string]interface{}{
		"Task": ctx.PID().String(),
	}).Info("[SERVER] Initialising Kafka Reader")
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   k.opts.listenAddresses,
		Topic:     k.opts.topic,
		GroupID:   k.opts.groupId,
		Partition: k.opts.partition,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
	})
	k.reader = reader
	logrus.WithFields(map[string]interface{}{
		"Task": ctx.PID().String(),
	}).Info("[SERVER] Kafka Reader Initialized")
}

func (k *KafkaServer) Accept(ctx *actor.Context) {
	for {
		m, err := k.reader.ReadMessage(context.Background())
		if err != nil {
			break
		}
		var target Message
		err = json.Unmarshal(m.Value, &target)
		if err != nil {
			logrus.WithFields(map[string]interface{}{
				"Task":  ctx.PID().String(),
				"Topic": k.opts.topic,
				"Error": err,
			}).Error("Unable to deserialize message")
			return
		}
		switch target.Type {
		case CREATE_ARRANGEMENT:
			var message CreateArrangement
			err = json.Unmarshal(target.Data, &message)
			if err != nil {
				logrus.WithFields(map[string]interface{}{
					"Task":  ctx.PID().String(),
					"Topic": k.opts.topic,
					"Error": err,
				}).Error("Unable to deserialize message")
				return
			}
			ctx.Send(ctx.Parent(), &message)
		case START:
			actor.WithMiddleware()
			var message Start
			err = json.Unmarshal(target.Data, &message)
			if err != nil {
				logrus.WithFields(map[string]interface{}{
					"Task":  ctx.PID().String(),
					"Topic": k.opts.topic,
					"Error": err,
				}).Error("Unable to deserialize message")
				return
			}
			ctx.Send(ctx.Parent(), message.WithTraceId(uuid.New()))
		case STOP:
			var message Stop
			err = json.Unmarshal(target.Data, &message)
			if err != nil {
				logrus.WithFields(map[string]interface{}{
					"Task":  ctx.PID().String(),
					"Topic": k.opts.topic,
					"Error": err,
				}).Error("Unable to deserialize message")
				return
			}
			ctx.Send(ctx.Parent(), &message)
		default:
			logrus.WithFields(map[string]interface{}{
				"Task":  ctx.PID().String(),
				"Topic": k.opts.topic,
				"Error": err,
			}).Error("Unknown Message Type")
		}
	}
}

func (k *KafkaServer) Send(ctx *actor.Context, message Message) {
	w := &kafka.Writer{
		Addr:     kafka.TCP(k.opts.listenAddresses...),
		Topic:    message.Task.String(),
		Balancer: &kafka.LeastBytes{},
	}

	m, err := json.Marshal(&message)
	if err != nil {
		logrus.WithFields(map[string]interface{}{
			"Task":        ctx.PID().String(),
			"Sender":      ctx.Sender().String(),
			"MessageType": message.Type,
			"MessageTask": message.Task.String(),
			"Message":     message.Data,
			"Error":       err,
		}).Error("Unable to marshal message")
	}

	err = w.WriteMessages(context.Background(),
		// NOTE: Each Message has Topic defined, otherwise an error is returned.
		kafka.Message{
			Topic: message.Task.String(),
			Key:   []byte(ctx.Sender().String()),
			Value: m,
		})
	if err != nil {

	}
}
