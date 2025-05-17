package chat_message

import (
	"database/sql"
	"log/slog"
)

type chatMessageStorage struct {
	db *sql.DB
	l  *slog.Logger
}

func New(db *sql.DB, l *slog.Logger) *chatMessageStorage {
	return &chatMessageStorage{db: db, l: l}
}
