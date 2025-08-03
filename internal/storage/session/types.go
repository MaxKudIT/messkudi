package session

import (
	"database/sql"
	"log/slog"
)

type sessionStorage struct {
	db *sql.DB
	l  *slog.Logger
}

func New(db *sql.DB, l *slog.Logger) *sessionStorage {
	return &sessionStorage{db: db, l: l}
}
