package chat_message

import (
	"database/sql"
	"log/slog"
)

type groupMessageStorage struct {
	db *sql.DB
	l  *slog.Logger
}

func New(db *sql.DB, l *slog.Logger) *groupMessageStorage {
	return &groupMessageStorage{db: db, l: l}
}
