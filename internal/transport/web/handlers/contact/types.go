package contact

import (
	"context"
	"github.com/google/uuid"
	"log/slog"
)

type contactService interface {
	AllContacts(ctx context.Context, userid uuid.UUID) ([]uuid.UUID, error)
	AddContact(ctx context.Context, userid uuid.UUID, contact uuid.UUID) error
}

type contactHandler struct {
	csv contactService
	l   *slog.Logger
}

func New(csv contactService, l *slog.Logger) *contactHandler {
	return &contactHandler{csv: csv, l: l}
}
