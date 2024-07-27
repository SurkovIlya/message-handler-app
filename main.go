package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/SurkovIlya/message-handler-app/internal/messeger"
	"github.com/SurkovIlya/message-handler-app/internal/model"
	"github.com/SurkovIlya/message-handler-app/internal/msgprocesser"
	"github.com/SurkovIlya/message-handler-app/internal/server"
	"github.com/SurkovIlya/message-handler-app/internal/sources/kafka/producer"
	"github.com/SurkovIlya/message-handler-app/internal/sources/kafka/sub"
	"github.com/SurkovIlya/message-handler-app/internal/statistics"
	"github.com/SurkovIlya/message-handler-app/internal/storage/pg"
	st "github.com/SurkovIlya/message-handler-app/pkg/postgres"
)

const serverPort = "8080"

func main() {
	pgParams := st.DBParams{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Database: os.Getenv("POSTGRES_DB"),
	}

	kafkaParams := model.KafkaParams{
		Host:     os.Getenv("KAFKA_HOST"),
		Topic:    "cool-topic",
		MaxBytes: 10e6,
	}

	conn, err := st.Connect(pgParams)
	if err != nil {
		panic(err)
	}

	storage := pg.New(st.New(conn))

	kafkaSub := sub.New(kafkaParams)
	msgProcesser := msgprocesser.New(storage, kafkaSub)
	go msgProcesser.Start()

	kafkaProd := producer.New()
	msgr := messeger.New(storage, kafkaProd)
	sc := statistics.New(storage)
	srv := server.New(serverPort, msgr, sc)

	go func() {
		if err := srv.Run(); err != nil {
			log.Panicf("error occured while running http server: %s", err.Error())
		}
	}()

	log.Println("app Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Println("app Shutting Down")

	msgProcesser.Stop()

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Panicf("error occured on server shutting down: %s", err.Error())
	}
}
