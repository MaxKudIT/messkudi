package websocket

import (
	"context"
	"github.com/gin-gonic/gin"
)

func (ur *wsrouters) WebsocketRegrouters(ctx context.Context, gr *gin.RouterGroup) {
	Websocket := gr.Group("/")
	{
		Websocket.GET("/ws", ur.wsh.WSHandler)

	}
}
