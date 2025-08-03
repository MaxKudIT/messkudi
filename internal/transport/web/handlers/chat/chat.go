package chat

import (
	"context"
	"fmt"
	"github.com/MaxKudIT/messkudi/internal/transport/web/dto"
	"github.com/MaxKudIT/messkudi/internal/transport/web/dto/chat_message_dto"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func (ch *chatHandler) ChatById(ctx context.Context, c *gin.Context) {
	ctxnew, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	id := c.Param("id")
	idu, err := uuid.Parse(id)
	if err != nil {
		ch.l.Info("Error parsing id", "error", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	chat, err := ch.csv.ChatById(ctxnew, idu)
	if err != nil {
		ch.l.Info("Error getting chat", "error", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ch.l.Info("Successfully got chat", "chat", chat)
	c.JSON(http.StatusOK, gin.H{"chat": chat})
}

func (ch *chatHandler) ChatDataByUsersId(ctx context.Context, c *gin.Context) {
	ctxnew, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	var object struct {
		//idone uuid.UUID
		Idtwo uuid.UUID
	}

	if err := c.ShouldBindJSON(&object); err != nil {
		ch.l.Error("Error parsing chatDTO", "error", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(object.Idtwo)
	id, exists := c.Get("user_id")
	if !exists {
		ch.l.Error("error not found")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "No user id found in context"})
		return
	}
	idone, err := uuid.Parse(id.(string))
	if err != nil {
		ch.l.Error("Error parsing uuid idone", "error", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	headerdata, err := ch.us.UserDataForChatHeader(ctxnew, object.Idtwo)
	if err != nil {
		ch.l.Info("Error getting header data", "error", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	idch, err := ch.csv.ChatByUsersId(ctx, idone, object.Idtwo)
	if err != nil {
		ch.l.Error("Error getting chat", "error", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err == nil && idch == uuid.Nil {
		c.JSON(http.StatusOK, gin.H{"headerdata": headerdata, "messages": []chat_message_dto.ChatMessageDTODetailsServer{}, "chat": nil})
		return
	}
	messages, err := ch.cms.AllMessages(ctxnew, idch)
	c.JSON(http.StatusOK, gin.H{"headerdata": headerdata, "messages": messages, "chat": idch})

}

func (ch *chatHandler) AllChatsPreview(ctx context.Context, c *gin.Context) {
	ctxnew, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	fmt.Println(1231313)

	userId, exists := c.Get("user_id")
	if !exists {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "user id not found"})
		return
	}
	userIdParse, err := uuid.Parse(userId.(string))
	if err != nil {
		ch.l.Error("Error parsing userId", "error", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	previews, err := ch.csv.AllChatsPreview(ctxnew, userIdParse)
	if err != nil {
		ch.l.Error("Error getting all previews", "error", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for _, preview := range previews {
		fmt.Println(len(preview.MessageMeta.UnReadMessages))
	}
	c.JSON(http.StatusOK, gin.H{"previews": previews})
}

func (ch *chatHandler) CreateChat(ctx context.Context, c *gin.Context) {
	ctxnew, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	creatorid, exists := c.Get("user_id")
	if !exists {
		ch.l.Info("Error getting creator id")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "No user_id found"})
		return
	}

	creatoridu, err := uuid.Parse(creatorid.(string))
	if err != nil {
		ch.l.Error("Error parsing creator id", "error", err)
		return
	}
	var chatDTO dto.ChatDTOClient

	if err := c.ShouldBindJSON(&chatDTO); err != nil {
		ch.l.Error("Error parsing chatDTO", "error", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	chatP, err := chatDTO.Parse()
	if err != nil {
		ch.l.Error("Error parsing chatDTO", "error", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	chatid := uuid.New()

	chat := dto.ToDomainChat(chatid, creatoridu, time.Now(), time.Now(), chatP)

	if err := ch.csv.CreateChat(ctxnew, chat); err != nil {
		ch.l.Error("Error creating chat", "error", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": chat.Id})
}

func (ch *chatHandler) DeleteChat(ctx context.Context, c *gin.Context) {
	ctxnew, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	id := c.Param("id")

	idu, err := uuid.Parse(id)
	if err != nil {
		ch.l.Error("Error parsing id", "error", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := ch.csv.DeleteChat(ctxnew, idu); err != nil {
		ch.l.Error("Error deleting chat", "error", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ch.l.Info("Successfully deleted chat", "chat", id)
	c.JSON(http.StatusOK, gin.H{"id": id})

}
