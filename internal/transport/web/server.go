package web

import "github.com/gin-gonic/gin"

func (r *routers) StartServer(port string) {
	router := gin.Default()
	router.Use(gin.Recovery())
	for _, ur := range r.ur {
		ur.UserRegrouters(router)
	}
	router.Run(port)
}
