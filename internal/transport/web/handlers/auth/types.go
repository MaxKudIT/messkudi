package auth

import (
	"context"
	"github.com/MaxKudIT/messkudi/internal/domain"
	"github.com/MaxKudIT/messkudi/internal/transport/web/dto"
	"github.com/gin-gonic/gin"
	"log/slog"
)

type authService interface {
	UserAuthData(ctx context.Context, usercr dto.UserCredentials) (domain.AuthData, error)
	AccessTokenUpdate(ctx context.Context, c *gin.Context, rt *dto.RefreshTokenDTO) (domain.AccessTokenUpdateData, error)
}

type authHandler struct {
	as authService
	l  *slog.Logger
}

func New(as authService, l *slog.Logger) *authHandler {
	return &authHandler{as: as, l: l}
}
