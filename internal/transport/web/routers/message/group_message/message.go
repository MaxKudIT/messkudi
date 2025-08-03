package chat_message

import (
	"context"
	"github.com/MaxKudIT/messkudi/internal/transport/web/middlewares"
	"github.com/gin-gonic/gin"
)

func (gmr *groupMessageRouter) GroupMessageRegRouters(ctx context.Context, router *gin.RouterGroup) {
	GroupMessage := router.Group("/gm")
	{
		GroupMessage.GET("/:id", middlewares.ValidateTokenAuthorization, func(c *gin.Context) { gmr.gmh.MessageById(c.Request.Context(), c) }) //middle
		GroupMessage.POST("/create", middlewares.ValidateTokenAuthorization, func(c *gin.Context) { gmr.gmh.CreateMessage(c.Request.Context(), c) })
		GroupMessage.DELETE("/delete", middlewares.ValidateTokenAuthorization, func(c *gin.Context) { gmr.gmh.DeleteMessage(c.Request.Context(), c) })
	}
}
