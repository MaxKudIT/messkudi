package middlewares

import (
	"fmt"
	"github.com/MaxKudIT/messkudi/internal/utils"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func ValidateTokenAuthorization(c *gin.Context) {

	fulltoken, err := c.Cookie("access_token")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
		return
	}
	if fulltoken == "" {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "no token found"})
		return
	}
	token := strings.TrimPrefix(fulltoken, "Bearer ")
	claims, err := utils.ValidateToken(token, os.Getenv("SECRET_KEY"))
	if err != nil {
		fmt.Println(1231313)
		if err.Error() == "token is expired" {
			fmt.Println(381283)
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "token is expired"})
			return
		}
		fmt.Println(69)
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "token is expired or not valid format"})
		return
	}
	user_id := claims["user_id"]
	c.Set("user_id", user_id)
	c.Next()
}
