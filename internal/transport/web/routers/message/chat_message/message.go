package chat_message

import (
	"context"
	"github.com/MaxKudIT/messkudi/internal/transport/web/middlewares"
	"github.com/gin-gonic/gin"
)

func (cmr *chatMessageRouter) ChatMessageRegRouters(ctx context.Context, router *gin.RouterGroup) {
	ChatMessage := router.Group("/cm")
	{
		ChatMessage.GET("/:id", middlewares.ValidateTokenAuthorization, func(c *gin.Context) { cmr.cmh.MessageById(c.Request.Context(), c) }) //middle
		ChatMessage.GET("/all/:id", middlewares.ValidateTokenAuthorization, func(c *gin.Context) { cmr.cmh.AllMessages(c.Request.Context(), c) })
		ChatMessage.POST("/create", middlewares.ValidateTokenAuthorization, func(c *gin.Context) { cmr.cmh.CreateMessage(c.Request.Context(), c) })
		ChatMessage.POST("/updater", middlewares.ValidateTokenAuthorization, func(c *gin.Context) { cmr.cmh.UpdateReadAtMessage(c.Request.Context(), c) })
		ChatMessage.DELETE("/delete", middlewares.ValidateTokenAuthorization, func(c *gin.Context) { cmr.cmh.DeleteMessage(c.Request.Context(), c) })
	}
}
