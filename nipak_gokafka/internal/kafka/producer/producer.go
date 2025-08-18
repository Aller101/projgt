package producer

import (
	"errors"
	"fmt"
	"strings"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var (
	errUnknownType = errors.New("unknown event type")
)

const (
	flushTimeout = 5000 //ms
)

type Producer struct {
	producer *kafka.Producer
}

func NewProducer(addr []string) (*Producer, error) {

	config := &kafka.ConfigMap{
		// "bootstrap.servers": addr[0],
		"bootstrap.servers":       strings.Join(addr, ","),
		"client.id":               "go-producer",
		"acks":                    "all",
		"retries":                 5,
		"socket.timeout.ms":       10000,
		"security.protocol":       "plaintext",
		"api.version.request":     "true",
		"broker.version.fallback": "3.7.0",
	}

	p, err := kafka.NewProducer(config)
	if err != nil {
		return nil, fmt.Errorf("error with new prodicer: %w", err)
	}

	return &Producer{producer: p}, nil
}

func (p *Producer) Produce(message string, topic string) error {
	kafkaMsg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value: []byte(message),
	}
	kafkaChan := make(chan kafka.Event)
	if err := p.producer.Produce(kafkaMsg, kafkaChan); err != nil {
		return fmt.Errorf("error with produce msg: %w", err)
	}

	e := <-kafkaChan

	switch ev := e.(type) {
	case *kafka.Message:
		return nil
	case kafka.Error:
		return ev
	default:
		return errUnknownType
	}
}

func (p *Producer) Close() {
	p.producer.Flush(flushTimeout)
	p.producer.Close()
}
