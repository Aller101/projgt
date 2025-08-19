package main

import (
	"fmt"
	"log"
	"log/slog"
	"nipak_gokafka/internal/kafka/producer"
)

var (
	address = []string{"localhost:9094"}
)

const (
	topic = "my-topic"
)

func main() {
	p, err := producer.NewProducer(address)
	if err != nil {
		log.Fatal(err)
	}

	defer p.Close()
	for i := range 100 {
		msg := fmt.Sprintf("kafka msg %d", i)
		if err := p.Produce(msg, topic); err != nil {
			slog.Error("err send msg", "err", err)
		}
	}
}
