package handlers

import (
	"github.com/MaxKudIT/messkudi/internal/domain/auth"
	"github.com/MaxKudIT/messkudi/internal/services/user"
	"github.com/MaxKudIT/messkudi/internal/utils"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type IAuthHandler interface {
	Authorization(c *gin.Context)
}

type AuthHandler struct {
	userserv user.UserServices
}

func NewAuthHandler(userserv user.UserServices) AuthHandler {
	return AuthHandler{userserv: userserv}
}
func (ah AuthHandler) Authorization(c *gin.Context) {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Println(err.Error())
		return
	}

	var data auth.UserCredentials
	if err := c.ShouldBindJSON(&data); err != nil {
		log.Println(err.Error())
		return
	}
	id := ah.userserv.GetUserIdService(data)
	accessToken, err := utils.CreateAccessToken(os.Getenv("JWT_SECRET"), id)
	if err != nil {
		log.Println(err.Error())
		return
	}
	c.JSON(200, gin.H{"Data": data, "AccessToken": accessToken, "RefreshToken": utils.CreateRefreshToken()})
}
