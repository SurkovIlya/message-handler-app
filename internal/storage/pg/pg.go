package pg

import "github.com/SurkovIlya/message-handler-app/pkg/postgres"

type PostgresStorage struct {
	storage *postgres.Database
}

func New(storage *postgres.Database) *PostgresStorage {
	return &PostgresStorage{
		storage: storage,
	}
}
