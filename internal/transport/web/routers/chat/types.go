package chat

import (
	"context"
	"github.com/gin-gonic/gin"
)

type chatHandler interface {
	ChatById(ctx context.Context, c *gin.Context)
	ChatDataByUsersId(ctx context.Context, c *gin.Context)
	AllChatsPreview(ctx context.Context, c *gin.Context)
	CreateChat(ctx context.Context, c *gin.Context)
	DeleteChat(ctx context.Context, c *gin.Context)
}

type chatRouter struct {
	ch chatHandler
}

func New(ch chatHandler) *chatRouter {
	return &chatRouter{ch: ch}
}
