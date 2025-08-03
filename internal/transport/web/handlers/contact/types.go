package contact

import (
	"context"
	"github.com/MaxKudIT/messkudi/internal/domain"
	"github.com/MaxKudIT/messkudi/internal/domain/contacts"
	"github.com/MaxKudIT/messkudi/internal/transport/web/dto"
	"github.com/google/uuid"
	"log/slog"
)

type contactService interface {
	AllContacts(ctx context.Context, userid uuid.UUID) ([]contacts.ContactPreview, error)
	IsMyContact(ctx context.Context, userid uuid.UUID, contactid uuid.UUID) (bool, error)
	AddContact(ctx context.Context, userid uuid.UUID, contact uuid.UUID) error
	DeleteContact(ctx context.Context, userid uuid.UUID, contactid uuid.UUID) error
}

type userService interface {
	CreateUser(ctx context.Context, userp domain.User) (domain.User, error)
	UserById(ctx context.Context, id uuid.UUID) (dto.UserDTO, error)
	UserByPhoneNumber(ctx context.Context, phoneNumber string) (dto.UserDTO, error)
	UserIdByPhoneNumber(ctx context.Context, phoneNumber string) (uuid.UUID, error)
	UserIsExistsByPhoneNumber(ctx context.Context, phoneNumber string) (bool, error)
	DeleteUser(ctx context.Context, id uuid.UUID) error
}

type contactHandler struct {
	csv contactService
	us  userService
	l   *slog.Logger
}

func New(csv contactService, us userService, l *slog.Logger) *contactHandler {
	return &contactHandler{csv: csv, us: us, l: l}
}
