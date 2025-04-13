package auth

import (
	"context"
	"github.com/MaxKudIT/messkudi/internal/transport/web/middlewares"
	"github.com/gin-gonic/gin"
)

func (ar *authRouter) AuthRegrouters(ctx context.Context, router *gin.RouterGroup) {
	Auth := router.Group("/")
	{
		Auth.POST("/auth", func(c *gin.Context) { ar.ah.UserAuthData(c.Request.Context(), c) })
		Auth.POST("/accesstoken", middlewares.ValidateTokenAuthorization, func(c *gin.Context) { ar.ah.AccessTokenUpdate(c.Request.Context(), c) })
	}
}
