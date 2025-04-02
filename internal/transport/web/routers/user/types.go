package user

import "github.com/gin-gonic/gin"

type userhandlers interface {
	GetUser(c *gin.Context)
	CreateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type userrouter struct {
	uh userhandlers
}

func New(uh userhandlers) *userrouter {
	return &userrouter{uh: uh}
}
