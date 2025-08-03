package auth

import (
	"context"
	"github.com/gin-gonic/gin"
)

func (ar *authRouter) AuthRegrouters(ctx context.Context, router *gin.RouterGroup) {
	Auth := router.Group("/")
	{
		Auth.POST("/auth", func(c *gin.Context) { ar.ah.UserAuthData(c.Request.Context(), c) })
		Auth.POST("/accesstoken", func(c *gin.Context) { ar.ah.AccessTokenUpdate(c.Request.Context(), c) })
		Auth.GET("/isAuth", func(c *gin.Context) { ar.ah.IsAuth(c.Request.Context(), c) })
		Auth.GET("/logout", func(c *gin.Context) { ar.ah.Logout(c.Request.Context(), c) })
	}
}
