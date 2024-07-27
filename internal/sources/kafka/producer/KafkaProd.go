package producer

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/SurkovIlya/message-handler-app/internal/model"
	"github.com/segmentio/kafka-go"
)

type KafkaProd struct {
	Writer *kafka.Writer
}

func New() *KafkaProd {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{os.Getenv("KAFKA_HOST")},
		Topic:    "cool-topic",
		Balancer: &kafka.LeastBytes{},
	})

	return &KafkaProd{
		Writer: writer,
	}
}

func (kkprod *KafkaProd) Send(message model.Message) error {
	err := kkprod.Writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(strconv.FormatUint(uint64(message.ID), 10)),
			Value: []byte(message.Value),
		},
	)
	if err != nil {
		return fmt.Errorf("failed to write messages: %v", err)
	}

	return nil
}
