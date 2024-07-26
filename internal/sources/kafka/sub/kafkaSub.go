package sub

import (
	"context"
	"log"
	"os"

	"github.com/segmentio/kafka-go"
)

type MessagerManager interface {
	ReadMsg(colum string, isRead bool, key string) error
}

type KafkaSub struct {
	Reader          *kafka.Reader
	MessagerManager MessagerManager
}

func New(mm MessagerManager) *KafkaSub {
	config := kafka.ReaderConfig{
		Brokers:  []string{os.Getenv("KAFKA_HOST")},
		Topic:    "example-topic",
		MaxBytes: 10,
	}

	reader := kafka.NewReader(config)

	return &KafkaSub{
		Reader:          reader,
		MessagerManager: mm,
	}
}

func (kafka *KafkaSub) Start() {
	for {
		msg, err := kafka.Reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("error readMessage: %s", err)

			continue
		}

		go func(key string) {
			err = kafka.MessagerManager.ReadMsg("is_read", true, key)
			if err != nil {
				log.Printf("error ReadMsg: %s", err)
			}
		}(string(msg.Key))
	}
}
