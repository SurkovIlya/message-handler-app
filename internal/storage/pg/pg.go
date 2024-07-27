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

func (pgs *PostgresStorage) GetStat() (model.Statistic, error) {
	var statistic model.Statistic
	query := `SELECT 
		COUNT(CASE WHEN processed = TRUE THEN 1 END) AS count_processed_true,
		COUNT(CASE WHEN processed = FALSE THEN 1 END) AS count_processed_false
		FROM messages;`

	row := pgs.storage.Conn.QueryRow(query)

	err := row.Scan(&statistic.Handled, &statistic.InProcess)
	if err != nil {
		return statistic, fmt.Errorf("error scan: %s", err)
	}

	return statistic, nil
}
