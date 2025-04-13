package auth

import (
	"database/sql"
	"log/slog"
)

type authStorage struct {
	db *sql.DB
	l  *slog.Logger
}

func New(db *sql.DB, l *slog.Logger) *authStorage {
	return &authStorage{db: db, l: l}
}
