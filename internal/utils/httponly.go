package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func SetAccessTokenHTTPOnlyCookie(c *gin.Context, accessToken string) {
	c.SetCookie(
		"access_token",
		accessToken,
		60*60*24*30,
		"/",
		"",
		false,
		true,
	)

}

func SetTokensAndIdHTTPOnlyCookie(c *gin.Context, accessToken, refreshToken string, id uuid.UUID) {
	c.SetCookie(
		"access_token",
		accessToken,
		60*60*24*30,
		"/",
		"",
		false,
		true,
	)

	c.SetCookie(
		"refresh_token",
		refreshToken,
		60*60*24*30,
		"/",
		"",
		false,
		true,
	)

	c.SetCookie(
		"user_id",
		id.String(),
		60*60*24*30,
		"/",
		"",
		false,
		true,
	)

}
func ClearAllCookies(c *gin.Context) {
	cookies := c.Request.Cookies()

	for _, cookie := range cookies {
		c.SetCookie(
			cookie.Name,
			"",
			-1,
			"/",
			cookie.Domain,
			cookie.Secure,
			cookie.HttpOnly,
		)
	}
}
