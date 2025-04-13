package middlewares

import (
	"github.com/MaxKudIT/messkudi/internal/utils"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func ValidateTokenAuthorization(c *gin.Context) {

	fulltoken := c.GetHeader("Authorization")
	if fulltoken == "" {

		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "no token found"})
		return
	}
	token := strings.TrimPrefix(fulltoken, "Bearer ")
	claims, err := utils.ValidateToken(token, os.Getenv("SECRET_KEY"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "token is expired or not valid format"})
		return
	}
	user_id := claims["user_id"]
	c.Set("user_id", user_id)
	c.Next()
}
