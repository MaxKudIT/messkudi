package web

import "github.com/gin-gonic/gin"

type userRouters interface {
	UserRegrouters(engine *gin.Engine)
}
type server interface {
	StartServer(port string)
}
type routers struct {
	ur []userRouters
}

func New(ur []userRouters) server {
	return &routers{ur: ur}
}
