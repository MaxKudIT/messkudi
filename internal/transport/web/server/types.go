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
type wsRouter interface {
	WebsocketRegrouters(ctx context.Context, gr *gin.RouterGroup)
}
type contactRouter interface {
	ContactRegRouters(ctx context.Context, router *gin.RouterGroup)
}

type chatRouter interface {
	ChatRegRouters(ctx context.Context, router *gin.RouterGroup)
}

type groupRouter interface {
	GroupRegRouters(ctx context.Context, router *gin.RouterGroup)
}

type chatMessageRouter interface {
	ChatMessageRegRouters(ctx context.Context, router *gin.RouterGroup)
}
type groupMessageRouter interface {
	GroupMessageRegRouters(ctx context.Context, router *gin.RouterGroup)
}

type server struct {
	ur  userRouter
	ar  authRouter
	wr  wsRouter
	cr  contactRouter
	cmr chatMessageRouter
	gmr groupMessageRouter
	chr chatRouter
	gr  groupRouter
}

func New(ur userRouter, ar authRouter, wr wsRouter, cr contactRouter, cmr chatMessageRouter, gmr groupMessageRouter, chr chatRouter, gr groupRouter) *server {
	return &server{ur: ur, ar: ar, wr: wr, cr: cr, cmr: cmr, gmr: gmr, chr: chr, gr: gr}
}
