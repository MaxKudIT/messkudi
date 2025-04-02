package middlewares

import (
	"github.com/MaxKudIT/messkudi/internal/utils"
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func ValidateTokenAuthorization(c *gin.Context) {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatal(err.Error())
		return
	}
	fulltoken := c.GetHeader("Authorization")
	if fulltoken == "" {

		c.String(403, "Нет прав доступав!")
		return
	}
	token := strings.TrimPrefix(fulltoken, "Bearer ")
	_, err := utils.ValidateToken(token, os.Getenv("JWT_SECRET"))
	if err != nil {

		c.String(403, err.Error())
		return
	}

	c.Next()
}
