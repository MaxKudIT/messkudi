package message

import (
	"database/sql"
	"log/slog"
)

type messageStorage struct {
	db *sql.DB
	l  *slog.Logger
}

func New(db *sql.DB, l *slog.Logger) *messageStorage {
	return &messageStorage{db: db, l: l}
}
