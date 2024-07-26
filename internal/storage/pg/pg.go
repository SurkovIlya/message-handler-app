package pg

import (
	"fmt"

	"github.com/SurkovIlya/message-handler-app/internal/server"
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

func (pgs *PostgresStorage) InsertMessage(message server.Message) error {
	query := `INSERT INTO messages (key_id, body_message) VALUES ($1, $2)`

	_, err := pgs.storage.Conn.Exec(query, message.ID, message.Value)
	if err != nil {
		return fmt.Errorf("error insert messages: %s", err)
	}

	return nil
}

func (pgs *PostgresStorage) UpdMesageProcessed(value int, key string) error {
	query := `UPDATE messages SET processed = $1 WHERE key_id = $2`

	_, err := pgs.storage.Conn.Exec(query, value, key)
	if err != nil {
		return fmt.Errorf("error upd messages: %s", err)
	}

	return nil
}

func (pgs *PostgresStorage) UpdMesageRead(value int, key string) error {
	query := `UPDATE messages SET is_read = $1 WHERE key_id = $2`

	_, err := pgs.storage.Conn.Exec(query, value, key)
	if err != nil {
		return fmt.Errorf("error upd messages: %s", err)
	}

	return nil
}
