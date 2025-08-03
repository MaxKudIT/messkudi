package chat_message

import (
	"context"
	"github.com/gin-gonic/gin"
)

type groupMessageHandler interface {
	MessageById(ctx context.Context, c *gin.Context)
	CreateMessage(ctx context.Context, c *gin.Context)
	DeleteMessage(ctx context.Context, c *gin.Context)
}

type groupMessageRouter struct {
	gmh groupMessageHandler
}

func New(gmh groupMessageHandler) *groupMessageRouter {
	return &groupMessageRouter{gmh: gmh}
}
