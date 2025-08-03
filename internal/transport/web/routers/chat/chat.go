package chat

import (
	"context"
	"github.com/MaxKudIT/messkudi/internal/transport/web/middlewares"
	"github.com/gin-gonic/gin"
)

func (cr *chatRouter) ChatRegRouters(ctx context.Context, router *gin.RouterGroup) {
	Chat := router.Group("/chat")
	{
		Chat.GET("/:id", middlewares.ValidateTokenAuthorization, func(c *gin.Context) { cr.ch.ChatById(c.Request.Context(), c) }) //middle
		Chat.GET("/all", middlewares.ValidateTokenAuthorization, func(c *gin.Context) { cr.ch.AllChatsPreview(c.Request.Context(), c) })
		Chat.POST("/find", middlewares.ValidateTokenAuthorization, func(c *gin.Context) { cr.ch.ChatDataByUsersId(c.Request.Context(), c) })
		Chat.POST("/create", middlewares.ValidateTokenAuthorization, func(c *gin.Context) { cr.ch.CreateChat(c.Request.Context(), c) })
		Chat.DELETE("/delete", middlewares.ValidateTokenAuthorization, func(c *gin.Context) { cr.ch.DeleteChat(c.Request.Context(), c) })
	}
}
