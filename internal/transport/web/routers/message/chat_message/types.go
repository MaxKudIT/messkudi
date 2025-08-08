package chat_message

import (
	"context"
	"github.com/gin-gonic/gin"
)

type chatMessageHandler interface {
	MessageById(ctx context.Context, c *gin.Context)
	AllMessages(ctx context.Context, c *gin.Context)
	CreateMessage(ctx context.Context, c *gin.Context)
	UpdateReadAtMessage(ctx context.Context, c *gin.Context)
	DeleteMessage(ctx context.Context, c *gin.Context)
}

type chatMessageRouter struct {
	cmh chatMessageHandler
}

func New(cmh chatMessageHandler) *chatMessageRouter {
	return &chatMessageRouter{cmh: cmh}
}
