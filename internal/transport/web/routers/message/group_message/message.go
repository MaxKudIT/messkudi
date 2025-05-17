package chat_message

import (
	"context"
	"github.com/gin-gonic/gin"
)

func (cmr *chatMessageRouter) ChatMessageRegRouters(ctx context.Context, router *gin.RouterGroup) {
	Contact := router.Group("/chatmessages")
	{
		Contact.GET("/:id", func(c *gin.Context) { cmr.cmh.MessageById(c.Request.Context(), c) }) //middle
		Contact.POST("/create", func(c *gin.Context) { cmr.cmh.CreateMessage(c.Request.Context(), c) })
		Contact.DELETE("/delete", func(c *gin.Context) { cmr.cmh.DeleteMessage(c.Request.Context(), c) })
	}
}
