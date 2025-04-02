package user

import (
	"github.com/MaxKudIT/messkudi/internal/transport/web/middlewares"
	"github.com/gin-gonic/gin"
)

func (ur *userrouter) UserRegrouters(engine *gin.Engine) {
	Users := engine.Group("/users")
	{
		Users.GET("/:id", middlewares.ValidateTokenAuthorization, ur.uh.GetUser)
		Users.POST("/registration", ur.uh.CreateUser)
		//Users.PATCH("/update", middlewares.ValidateTokenAuthorization, (*ur.uh).UpdateUser)
		Users.DELETE("/:id", middlewares.ValidateTokenAuthorization, ur.uh.DeleteUser)
	}
}
