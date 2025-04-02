package routers

import (
	"github.com/MaxKudIT/messkudi/internal/transport/web/handlers"
	"github.com/gin-gonic/gin"
)

type AuthRouter struct {
	authhand handlers.IAuthHandler
}

func NewAuthRouters(authhand handlers.AuthHandler) AuthRouter {
	return AuthRouter{authhand: authhand}
}

func (ar AuthRouter) AuthRegistration(engine *gin.Engine) {
	Auth := engine.Group("/")
	{
		Auth.POST("/login", ar.authhand.Authorization)
	}
}
