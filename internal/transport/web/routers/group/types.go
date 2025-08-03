package group

import (
	"context"
	"github.com/gin-gonic/gin"
)

type groupHandler interface {
	GroupById(ctx context.Context, c *gin.Context)
	CreateGroup(ctx context.Context, c *gin.Context)
	JoinGroup(ctx context.Context, c *gin.Context)
	DeleteGroup(ctx context.Context, c *gin.Context)
}

type groupRouter struct {
	gh groupHandler
}

func New(gh groupHandler) *groupRouter {
	return &groupRouter{gh: gh}
}
