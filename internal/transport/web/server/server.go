package server

import (
	"context"
	"github.com/gin-gonic/gin"
)

func (s *server) Create() *gin.Engine {
	router := gin.Default()
	maingr := router.Group("") //теперь принадлежит основному router
	{
		s.ur.UserRegrouters(context.TODO(), maingr) //группа maingr обновляется в основном router
		s.ar.AuthRegrouters(context.TODO(), maingr)
	}

	return router
}
