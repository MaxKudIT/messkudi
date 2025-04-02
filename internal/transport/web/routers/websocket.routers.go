package routers

import (
	"github.com/MaxKudIT/messkudi/internal/transport/web/handlers"
	"github.com/gin-gonic/gin"
)

func webSocketRegistration(engine *gin.Engine) {
	WebSocket := engine.Group("/")
	{
		WebSocket.GET("/ws", handlers.WebsocketHandler)
	}
}
