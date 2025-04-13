package user

import (
	"context"
	"github.com/MaxKudIT/messkudi/internal/domain"
	"github.com/MaxKudIT/messkudi/internal/transport/web/dto"
	"github.com/google/uuid"
	"log/slog"
)

type userservice interface {
	CreateUser(ctx context.Context, userCr domain.User) (domain.User, error)
	UserById(ctx context.Context, id uuid.UUID) (dto.UserDTO, error)
	UserByPhoneNumber(ctx context.Context, phoneNumber string) (dto.UserDTO, error)
	UserIsExistsByPhoneNumber(ctx context.Context, phoneNumber string) (bool, error)
	//UpdateUser(c *gin.Context)
	DeleteUser(ctx context.Context, id uuid.UUID) error
}

type userhandler struct {
	us userservice
	l  *slog.Logger
}

func New(us userservice, l *slog.Logger) *userhandler {
	return &userhandler{us: us, l: l}
}
