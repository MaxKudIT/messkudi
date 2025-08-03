package contact

import (
	"database/sql"
	"log/slog"
)

type contactStorage struct {
	db *sql.DB
	l  *slog.Logger
}

func New(db *sql.DB, l *slog.Logger) *contactStorage {
	return &contactStorage{db: db, l: l}
}
