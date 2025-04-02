package user

import "database/sql"

type userStorage struct {
	db *sql.DB
}

func New(db *sql.DB) *userStorage {
	return &userStorage{db: db}
}
