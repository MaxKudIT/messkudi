package group

import (
	"context"
	"github.com/MaxKudIT/messkudi/internal/transport/web/middlewares"
	"github.com/gin-gonic/gin"
)

func (gr *groupRouter) GroupRegRouters(ctx context.Context, router *gin.RouterGroup) {
	Group := router.Group("/group")
	{
		Group.GET("/:id", middlewares.ValidateTokenAuthorization, func(c *gin.Context) { gr.gh.GroupById(c.Request.Context(), c) }) //middle
		Group.POST("/create", middlewares.ValidateTokenAuthorization, func(c *gin.Context) { gr.gh.CreateGroup(c.Request.Context(), c) })
		Group.POST("/join", middlewares.ValidateTokenAuthorization, func(c *gin.Context) { gr.gh.JoinGroup(c.Request.Context(), c) })
		Group.DELETE("/delete", middlewares.ValidateTokenAuthorization, func(c *gin.Context) { gr.gh.DeleteGroup(c.Request.Context(), c) })
	}
}
