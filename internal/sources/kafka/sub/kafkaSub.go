package sub

import (
	"context"
	"log"
	"strconv"

	"github.com/SurkovIlya/message-handler-app/internal/model"
	"github.com/segmentio/kafka-go"
)

type KafkaSub struct {
	consumer *kafka.Reader
	msgCh    chan model.Message
	exitCh   chan struct{}
}

func New(params model.KafkaParams) *KafkaSub {
	config := kafka.ReaderConfig{
		Brokers:  []string{params.Host},
		Topic:    params.Topic,
		MaxBytes: params.MaxBytes,
	}

	consumer := kafka.NewReader(config)

	return &KafkaSub{
		consumer: consumer,
		msgCh:    make(chan model.Message),
		exitCh:   make(chan struct{}),
	}
}

func (kafka *KafkaSub) Stop() {
	kafka.exitCh <- struct{}{}
}

func (kafka *KafkaSub) StartReading(ctx context.Context) {
	for {
		select {
		case _, ok := <-kafka.exitCh:
			if ok {
				close(kafka.msgCh)
				close(kafka.exitCh)

				log.Println("kafka reading stopped by message from exit channel")

				return
			}

		default:
			msg, err := kafka.consumer.ReadMessage(ctx)
			if err != nil {
				log.Printf("error kafka readMessage: %s", err)

				continue
			}

			// err = kafka.consumer.CommitMessages(ctx, msg)
			// if err != nil {
			// 	log.Printf("error commit message: %s", err)
			// }

			i, err := strconv.Atoi(string(msg.Key))
			if err != nil {
				log.Printf("error kafka Atoi: %s", err)
			}

			kafka.msgCh <- model.Message{
				ID:    uint32(i),
				Value: string(msg.Value),
			}
		}
	}
}

func (kafka *KafkaSub) GetMsgCh() chan model.Message {
	return kafka.msgCh
}
