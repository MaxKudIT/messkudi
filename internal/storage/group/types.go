package group

import (
	"database/sql"
	"log/slog"
)

type groupStorage struct {
	db *sql.DB
	l  *slog.Logger
}

func New(db *sql.DB, l *slog.Logger) *groupStorage {
	return &groupStorage{db: db, l: l}
}
