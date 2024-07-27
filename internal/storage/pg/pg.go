package pg

import (
	"fmt"

	"github.com/SurkovIlya/message-handler-app/internal/model"
	"github.com/SurkovIlya/message-handler-app/pkg/postgres"
)

type PostgresStorage struct {
	storage *postgres.Database
}

func New(storage *postgres.Database) *PostgresStorage {
	return &PostgresStorage{
		storage: storage,
	}
}

func (pgs *PostgresStorage) InsertMessage(message model.Message) error {
	query := `INSERT INTO messages (id, body_message) VALUES ($1, $2)`

	_, err := pgs.storage.Conn.Exec(query, message.ID, message.Value)
	if err != nil {
		return fmt.Errorf("error insert messages: %s", err)
	}

	return nil
}

func (pgs *PostgresStorage) UpdMesageProcessed(id uint32) error {
	query := `UPDATE messages SET processed = $1 WHERE id = $2`

	_, err := pgs.storage.Conn.Exec(query, true, id)
	if err != nil {
		return fmt.Errorf("error upd messages: %s", err)
	}

	return nil
}
