package producer

import (
	"context"
	"fmt"
	"os"

	"github.com/SurkovIlya/message-handler-app/internal/server"
	"github.com/segmentio/kafka-go"
)

type KafkaProd struct {
	Writer *kafka.Writer
}

func New() *KafkaProd {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{os.Getenv("KAFKA_HOST")},
		Topic:    "example-topic",
		Balancer: &kafka.LeastBytes{},
	})

	return &KafkaProd{
		Writer: writer,
	}
}

func (kkprod *KafkaProd) Send(message server.Message) error {
	err := kkprod.Writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(message.ID),
			Value: []byte(message.Value),
		},
	)
	if err != nil {
		return fmt.Errorf("failed to write messages: %v", err)
	}
	// defer kkprod.Writer.Close()

	return nil
}
