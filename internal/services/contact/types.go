package contact

import (
	"context"
	"github.com/MaxKudIT/messkudi/internal/domain/contacts"
	"github.com/google/uuid"
	"log/slog"
)

type contactStorage interface {
	AddContact(ctx context.Context, userid uuid.UUID, contactid uuid.UUID) error
	IsMyContact(ctx context.Context, id uuid.UUID, contactId uuid.UUID) (bool, error)
	AllContacts(ctx context.Context, userid uuid.UUID) ([]contacts.ContactPreview, error)
	DeleteContact(ctx context.Context, id uuid.UUID, contactId uuid.UUID) error
}

type contactService struct {
	cs contactStorage
	l  *slog.Logger
}

func New(cs contactStorage, l *slog.Logger) *contactService {
	return &contactService{cs: cs, l: l}
}
