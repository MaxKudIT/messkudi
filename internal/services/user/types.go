package user

import (
	"github.com/MaxKudIT/messkudi/internal/domain/user"
	"github.com/google/uuid"
	"log/slog"
)

type userStorage interface {
	GetUserById(id uuid.UUID) (user.User, error)
	SaveUser(userp user.User) error
	DeleteUser(id uuid.UUID) error
}

type userService struct {
	us userStorage
	l  *slog.Logger
}

func New(us userStorage, l *slog.Logger) *userService {
	return &userService{us: us, l: l}
}
