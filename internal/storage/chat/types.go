package chat

import (
	"database/sql"
	"log/slog"
)

type chatStorage struct {
	db *sql.DB
	l  *slog.Logger
}

func New(db *sql.DB, l *slog.Logger) *chatStorage {
	return &chatStorage{db: db, l: l}
}
