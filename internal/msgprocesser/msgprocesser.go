package msgprocesser

import (
	"context"
	"log"
	"time"

	"github.com/SurkovIlya/message-handler-app/internal/model"
)

const maxGoroutinesCnt = 20

type Storage interface {
	UpdMesageProcessed(id uint32) error
}

type Broker interface {
	StartReading(ctx context.Context)
	Stop()
	GetMsgCh() chan model.Message
}

type MsgProcesser struct {
	db      Storage
	broker  Broker
	limiter chan struct{}
}

func New(st Storage, br Broker) *MsgProcesser {
	return &MsgProcesser{
		db:      st,
		broker:  br,
		limiter: make(chan struct{}, maxGoroutinesCnt),
	}
}

func (mp *MsgProcesser) Start() {
	go mp.broker.StartReading(context.Background())

	msgCh := mp.broker.GetMsgCh()

	for val := range msgCh {
		mp.limiter <- struct{}{}
		go func(v model.Message) {
			defer func() {
				<-mp.limiter
			}()

			err := handleMsg(v)
			if err != nil {
				log.Printf("error handleMsg: %s", err)

				return
			}

			err = mp.db.UpdMesageProcessed(v.ID)
			if err != nil {
				log.Printf("error upd (processed) message: %s", err)
			}
		}(val)

	}
}

func (mp *MsgProcesser) Stop() {
	mp.broker.Stop()
}

func handleMsg(msg model.Message) error {
	time.Sleep(100 * time.Millisecond)

	log.Printf("HANDLE: %v", msg)

	return nil
}
