package messeger

import "github.com/SurkovIlya/message-handler-app/internal/server"

type PostgresStorage interface {
}

type MassagerManager struct {
	QueryStorage PostgresStorage
}

func New(query PostgresStorage) *MassagerManager {
	return &MassagerManager{
		QueryStorage: query,
	}
}

func (mm *MassagerManager) Receiving(message server.ReceivingBody) error {
	return nil
}
