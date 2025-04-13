package user

import (
	"context"
	"github.com/gin-gonic/gin"
)

type userhandlers interface {
	UserById(ctx context.Context, c *gin.Context)
	CreateUser(ctx context.Context, c *gin.Context)
	UserByPhoneNumber(ctx context.Context, c *gin.Context)
	UserIsExistsByPhoneNumber(ctx context.Context, c *gin.Context)
	DeleteUser(ctx context.Context, c *gin.Context)
}

type userrouter struct {
	uh userhandlers
}

func New(uh userhandlers) *userrouter {
	return &userrouter{uh: uh}
}
