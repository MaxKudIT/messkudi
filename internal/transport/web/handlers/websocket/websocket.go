package websocket

import (
	"context"
	"fmt"
	"github.com/MaxKudIT/messkudi/internal/domain/clients"
	websocket2 "github.com/MaxKudIT/messkudi/internal/services/websocket"
	"github.com/MaxKudIT/messkudi/internal/transport/web/dto/chat_message_dto"

	"github.com/MaxKudIT/messkudi/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"os"
)

func (wsh *websockethandler) WSHandler(c *gin.Context) {

	ctxnew, cancel := context.WithCancel(context.Background())
	defer cancel()
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		wsh.l.Error("Error upgrading to WebSocket", "error", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "WebSocket upgrade failed"})
		return
	}
	fmt.Println(conn.LocalAddr().String())
	token, err := c.Cookie("access_token")
	if err != nil {
		wsh.l.Error("Error getting token", "error", err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Error getting token"})
		return
	}
	claims, err := utils.ValidateToken(token, os.Getenv("SECRET_KEY"))
	if err != nil {
		wsh.l.Error("Error validating token by first message", "error", err)
		return
	}

	wsh.l.Info("Current connection: ", conn.RemoteAddr().String())
	if err != nil {
		wsh.l.Error("Error while connection ws")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	wsh.l.Info("Success connection ws")

	//if err := wsh.rmq.Setup("queue_first"); err != nil {
	//	wsh.l.Error("Error setting up rabbitmq", "error", err)
	//}

	uuidClient := uuid.New()

	userIdstr := claims["user_id"]
	userId, err := uuid.Parse(userIdstr.(string))
	if err != nil {
		wsh.l.Error("Error parsing user id", "error", err)
		return
	}
	client := clients.Client{Conn: conn, ClientId: uuidClient}
	clients.Session.AddClient(userId, &client)

	contactsid, err := wsh.csv.AllContacts(ctxnew, userId)
	if contactsid == nil {

		wsh.l.Info("user dont have contacts!")
	}
	//go func() {
	//	ticker := time.NewTicker(10 * time.Second)
	//	defer ticker.Stop()
	//
	//	for {
	//		select {
	//		case <-ticker.C:
	//			if err := conn.WriteMessage(websocket.PingMessage, nil); err != nil {
	//				cancel()
	//				return
	//			}
	//		case <-ctxnew.Done():
	//			return
	//		}
	//	}
	//}()
	//
	//conn.SetPongHandler(func(string) error {
	//	wsh.l.Info("pong received")
	//	return nil
	//})

	//for {
	//	_, message, err := conn.ReadMessage()
	//	if err != nil {
	//		wsh.l.Error("Error while reading message pong")
	//		break
	//	}
	//	wsh.l.Info(string(message))
	//}

	wss := websocket2.New(userId, make(chan chat_message_dto.ChatMessageDTOClientParsing, 40), wsh.cms, wsh.cs, wsh.rmq, wsh.l)
	go func() {
		if err := wss.Read(); err != nil {
			wsh.l.Error("Error while reading", "error", err)
			return
		}
		wsh.l.Info("Success reading end")
		return
	}()
	go func() {
		if err := wss.Write(); err != nil {
			wsh.l.Error("Error while writing", "error", err)
			return
		}
		wsh.l.Info("Success writing end")
		return
	}()

}
