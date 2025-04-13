package server

import (
	"context"
	"github.com/gin-gonic/gin"
)

type userRouter interface {
	UserRegrouters(ctx context.Context, gr *gin.RouterGroup)
}
type authRouter interface {
	AuthRegrouters(ctx context.Context, gr *gin.RouterGroup)
}

type server struct {
	ur userRouter
	ar authRouter
}

func New(ur userRouter, ar authRouter) *server {
	return &server{ur: ur, ar: ar}
}
