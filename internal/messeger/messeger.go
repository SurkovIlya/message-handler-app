package messeger

import (
	"fmt"

	"github.com/SurkovIlya/message-handler-app/internal/server"
)

type PostgresStorage interface {
	InsertMessage(message server.Message) error
	UpdMesageProcessed(value int, key string) error
	UpdMesageRead(value int, key string) error
}

type KafkaProd interface {
	Send(message server.Message) error
}

type MassagerManager struct {
	QueryStorage PostgresStorage
	KafkaProd    KafkaProd
}

func New(query PostgresStorage, producer KafkaProd) *MassagerManager {
	return &MassagerManager{
		QueryStorage: query,
		KafkaProd:    producer,
	}
}

func (mm *MassagerManager) Receiving(message server.Message) error {
	err := mm.QueryStorage.InsertMessage(message)
	if err != nil {
		return fmt.Errorf("error insert message: %s", err)
	}

	err = mm.KafkaProd.Send(message)
	if err != nil {
		return fmt.Errorf("error kafka send: %s", err)
	}

	err = mm.QueryStorage.UpdMesageProcessed(1, message.ID)
	if err != nil {
		return fmt.Errorf("error upd (processed) message: %s", err)
	}

	return nil
}

func (mm *MassagerManager) ReadMsg(colum string, isRead bool, key string) error {
	err := mm.QueryStorage.UpdMesageRead(1, key)
	if err != nil {
		return fmt.Errorf("error upd (is_read) message: %s", err)
	}

	return nil
}
