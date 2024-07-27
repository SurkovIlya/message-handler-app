package messeger

import (
	"fmt"
	"hash/crc32"
	"time"

	"github.com/SurkovIlya/message-handler-app/internal/model"
)

type PostgresStorage interface {
	InsertMessage(message model.Message) error
}

type KafkaProd interface {
	Send(message model.Message) error
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

func (mm *MassagerManager) Receiving(message model.Message) error {
	message.ID = calculateCRC32(message.Value)

	err := mm.QueryStorage.InsertMessage(message)
	if err != nil {
		return fmt.Errorf("error insert message: %s", err)
	}

	err = mm.KafkaProd.Send(message)
	if err != nil {
		return fmt.Errorf("error kafka send: %s", err)
	}

	return nil
}

func calculateCRC32(data string) uint32 {
	nano := fmt.Sprint(time.Now().UnixNano())
	return crc32.ChecksumIEEE([]byte(data + nano))
}
