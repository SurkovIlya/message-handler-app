package msgprocesser

import (
	"log"
	"time"

	"github.com/SurkovIlya/message-handler-app/internal/model"
)

type Storage interface {
	UpdMesageProcessed(id uint32) error
}

type Broker interface {
	StartReading()
	Stop()
	GetMsgCh() chan model.Message
}

type MsgProcesser struct {
	DB     Storage
	Broker Broker
}

func New(st Storage, br Broker) *MsgProcesser {
	return &MsgProcesser{
		DB:     st,
		Broker: br,
	}
}

func (mp *MsgProcesser) Start() {
	go mp.Broker.StartReading()

	msgCh := mp.Broker.GetMsgCh()

	for val := range msgCh {
		err := handleMsg(val)
		if err != nil {
			log.Printf("error handleMsg: %s", err)

			continue
		}

		err = mp.DB.UpdMesageProcessed(val.ID)
		if err != nil {
			log.Printf("error upd (processed) message: %s", err)
		}
	}
}

func (mp *MsgProcesser) Stop() {
	mp.Broker.Stop()
}

func handleMsg(msg model.Message) error {
	time.Sleep(100 * time.Millisecond)

	log.Printf("HANDLE: %v", msg)

	return nil
}
