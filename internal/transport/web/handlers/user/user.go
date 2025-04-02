package user

import (
	"github.com/MaxKudIT/messkudi/internal/dto"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func (uh *userhandler) GetUser(c *gin.Context) {
	id := c.Param("id")
	uuid, _ := uuid.Parse(id)
	user, err := uh.us.GetUser(uuid)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{"Data:": user})
}

func (uh *userhandler) CreateUser(c *gin.Context) {
	var userdt dto.UserDTO

	if err := c.ShouldBindJSON(&userdt); err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := uh.us.CreateUser(userdt)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request payload",
			"details": err.Error(),
		})
		return

	}
	c.JSON(201, gin.H{"Data:": user})
}

//	func (ush UserHandler) UpdateUser(c *gin.Context) {
//		var user user.UserUpdate
//
//		if err := c.ShouldBindJSON(&user); err != nil {
//			log.Print(exceptions.CreateUserExc().Error())
//		}
//		ush.userserv.UpdateUserService(user)
//	}

func (uh *userhandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	uuid, _ := uuid.Parse(id)
	if err := uh.us.DeleteUser(uuid); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request payload",
			"details": err.Error(),
		})
	}
}
