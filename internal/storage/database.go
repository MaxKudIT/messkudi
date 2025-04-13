package storage

import (
	"database/sql"
)

type database struct {
	connection string
}

func NewDatabase(connection string) *database {
	return &database{connection: connection}
}

func (dbd database) ConnectionDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", dbd.connection)
	if err != nil {
		return &sql.DB{}, err
	}

	return db, nil
}
