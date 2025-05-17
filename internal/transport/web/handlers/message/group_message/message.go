package chat_message

import (
	"context"
	"github.com/MaxKudIT/messkudi/internal/domain/messages"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func (cmh *chatMessageHandler) MessageById(ctx context.Context, c *gin.Context) {
	ctxnew, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	id := c.Param("id")
	idu, err := uuid.Parse(id)
	if err != nil {
		cmh.l.Error("Error parsing id", "error", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	message, err := cmh.cmsv.MessageById(ctxnew, idu)
	if err != nil {
		cmh.l.Error("Error getting chat message", "error", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	cmh.l.Info("Successfully getting chat message")
	c.JSON(http.StatusOK, gin.H{"message": message})
}

func (cmh *chatMessageHandler) CreateMessage(ctx context.Context, c *gin.Context) {
	ctxnew, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var message messages.ChatMessage

	if err := c.ShouldBindJSON(&message); err != nil {
		cmh.l.Error("Error parsing message", "error", err)
		return
	}
	if err := cmh.cmsv.CreateMessage(ctxnew, message); err != nil {
		cmh.l.Error("Error creating chat message", "error", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error})
	}
	cmh.l.Info("Successfully created chat message")
	c.JSON(http.StatusOK, gin.H{"message": message})
}
func (cmh *chatMessageHandler) DeleteMessage(ctx context.Context, c *gin.Context) {
	ctxnew, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	id := c.Param("id")
	idu, err := uuid.Parse(id)
	if err != nil {
		cmh.l.Error("Error parsing id", "error", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := cmh.cmsv.DeleteMessage(ctxnew, idu); err != nil {
		cmh.l.Error("Error deleting chat message", "error", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	cmh.l.Info("Successfully deleting chat message")
	c.JSON(http.StatusOK, gin.H{"message deleting: ": idu})
}
