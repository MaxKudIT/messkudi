package user

import (
	"context"
	"github.com/MaxKudIT/messkudi/internal/domain"
	"github.com/MaxKudIT/messkudi/internal/transport/web/dto"
	"github.com/google/uuid"
	"log/slog"
)

type userStorage interface {
	UserById(ctx context.Context, id uuid.UUID) (dto.UserDTO, error)
	SaveUser(ctx context.Context, userp domain.User) error
	UserByPhoneNumber(ctx context.Context, phoneNumber string) (dto.UserDTO, error)
	UserIsExistsByPhoneNumber(ctx context.Context, phoneNumber string) (bool, error)
	UserIdByPhoneNumber(ctx context.Context, phoneNumber string) (uuid.UUID, error)
	DeleteUser(ctx context.Context, id uuid.UUID) error
	UserDataForChatHeader(ctx context.Context, id uuid.UUID) (dto.ChatHeader, error)
}

type userService struct {
	us userStorage
	l  *slog.Logger
}

func New(us userStorage, l *slog.Logger) *userService {
	return &userService{us: us, l: l}
}
