package contact

import (
	"context"
	"github.com/MaxKudIT/messkudi/internal/transport/web/middlewares"
	"github.com/gin-gonic/gin"
)

func (cr *contactrouter) ContactRegRouters(ctx context.Context, router *gin.RouterGroup) {
	Contact := router.Group("/contacts")
	{

		Contact.GET("/all", middlewares.ValidateTokenAuthorization, func(c *gin.Context) { cr.ch.AllContacts(c.Request.Context(), c) }) //middle
		Contact.POST("/my", middlewares.ValidateTokenAuthorization, func(c *gin.Context) { cr.ch.IsMyContact(c.Request.Context(), c) }) //middle
		Contact.POST("/add", middlewares.ValidateTokenAuthorization, func(c *gin.Context) { cr.ch.AddContact(c.Request.Context(), c) })
		Contact.DELETE("/:id", middlewares.ValidateTokenAuthorization, func(c *gin.Context) { cr.ch.DeleteContact(c.Request.Context(), c) })
	}
}
