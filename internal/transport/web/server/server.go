package server

import (
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func (s *server) Create() *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	maingr := router.Group("") //теперь принадлежит основному router
	{
		s.ur.UserRegrouters(context.TODO(), maingr) //группа maingr обновляется в основном router
		s.ar.AuthRegrouters(context.TODO(), maingr)
		s.wr.WebsocketRegrouters(context.TODO(), maingr)
		s.cr.ContactRegRouters(context.TODO(), maingr)
		s.cmr.ChatMessageRegRouters(context.TODO(), maingr)
		s.gmr.GroupMessageRegRouters(context.TODO(), maingr)
		s.chr.ChatRegRouters(context.TODO(), maingr)
		s.gr.GroupRegRouters(context.TODO(), maingr)
	}

	return router
}
