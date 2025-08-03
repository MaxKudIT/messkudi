package auth

import (
	"context"
	"github.com/gin-gonic/gin"
)

type authHandler interface {
	UserAuthData(ctx context.Context, c *gin.Context)
	AccessTokenUpdate(ctx context.Context, c *gin.Context)
	Logout(ctx context.Context, c *gin.Context)
	IsAuth(ctx context.Context, c *gin.Context)
}

type authRouter struct {
	ah authHandler
}

func New(ah authHandler) *authRouter {
	return &authRouter{ah: ah}
}
