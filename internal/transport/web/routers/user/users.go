package user

import (
	"context"
	"github.com/MaxKudIT/messkudi/internal/transport/web/middlewares"
	"github.com/gin-gonic/gin"
)

func (ur *userrouter) UserRegrouters(ctx context.Context, gr *gin.RouterGroup) {

	Users := gr.Group("/users")
	{

		Users.POST("/registration", func(c *gin.Context) { ur.uh.CreateUser(c.Request.Context(), c) })
		Users.GET("/isExists/:phonenumber", func(c *gin.Context) { ur.uh.UserIsExistsByPhoneNumber(c.Request.Context(), c) })
		UsersAuth := Users.Group("")
		UsersAuth.Use(middlewares.ValidateTokenAuthorization)
		{
			UsersAuth.GET("/:id", func(c *gin.Context) { ur.uh.UserById(c.Request.Context(), c) })
			UsersAuth.DELETE("/:id", func(c *gin.Context) { ur.uh.DeleteUser(c.Request.Context(), c) })
		}
	}

}
