package websocket

import (
	"github.com/gin-gonic/gin"
)

type wshandlers interface {
	WSHandler(c *gin.Context)
}

type wsrouters struct {
	wsh wshandlers
}

func New(wsh wshandlers) *wsrouters {
	return &wsrouters{wsh: wsh}
}
