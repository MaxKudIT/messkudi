package websocket

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/MaxKudIT/messkudi/internal/domain/chats"
	"github.com/MaxKudIT/messkudi/internal/domain/clients"
	"github.com/MaxKudIT/messkudi/internal/transport/web/dto/chat_message_dto"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"time"
)

func (wbs *websocketService) Read() error {
	//var message messages.Message //chatid - на клиете, Content - очевидно, isCreator - из БД,

	ctx, cancel := context.WithCancel(context.Background())

	client := clients.Session.LoadClient(wbs.sid)
	clients.Session.All()
	if client == nil {
		wbs.l.Info("client not found")
	}
	defer clients.Session.RemoveClient(wbs.sid)
	defer cancel()
	defer func() {
		close(wbs.mailing)
		client.Conn.Close()
	}()

	for {

		var message chat_message_dto.ChatMessageDTOClient
		_, messaget, err := client.Conn.ReadMessage() // сервер читает сообщение
		if err != nil {
			wbs.l.Error("Error reading", "error", err)
			return err
		}

		if err := json.Unmarshal(messaget, &message); err != nil {
			wbs.l.Error("Error while unmarshalling message", "error", err)
			return err
		}
		messageP, err := message.UuidParse()
		fmt.Println(messageP.Id, messageP.ChatId, messageP.CorrespondenceType)
		if err != nil {
			wbs.l.Error("Error while parsing message", "error", err)
			return err
		}
		client := clients.Session.LoadClient(messageP.RecieverId)
		messageDomain := chat_message_dto.ToDomain(time.Now(), time.Now(), nil, messageP)
		switch {
		// 1 - клиент в сети, чат создан; 2 - клиент не в сети, чат не создан; 3 - клиент не в сети, чат создан; 4 - клиент в сети, чат не создан

		case client != nil && messageDomain.ChatId == uuid.Nil: //клиент в сети, чат не создан
			fmt.Print("Здесь")
			chatId := uuid.New()
			if err := wbs.cs.SaveChat(ctx, chats.Chat{Id: chatId, CreatorId: messageP.SenderId, Participant: messageP.RecieverId, CreatedAt: time.Now(), UpdatedAt: time.Now()}); err != nil {
				wbs.l.Error("Error saving chat", "error", err)
				return err
			}
			messageDomain.ChatId = chatId
			if err := wbs.cms.SaveMessage(ctx, messageDomain); err != nil {
				wbs.l.Error("Error while saving message", "error", err)
				return err
			}
			wbs.mailing <- messageP
		case client != nil && messageDomain.ChatId != uuid.Nil: //клиент в сети, чат создан
			if err := wbs.cms.SaveMessage(ctx, messageDomain); err != nil {
				wbs.l.Error("Error while saving message", "error", err)
				return err
			}
			wbs.mailing <- messageP
		case client == nil && messageDomain.ChatId == uuid.Nil: //клиент не в сети, чат не создан
			chatId := uuid.New()
			if err := wbs.cs.SaveChat(ctx, chats.Chat{Id: chatId, CreatorId: messageP.SenderId, Participant: messageP.RecieverId, CreatedAt: time.Now(), UpdatedAt: time.Now()}); err != nil {
				wbs.l.Error("Error saving chat", "error", err)
				return err
			}
			messageDomain.ChatId = chatId
			if err := wbs.cms.SaveMessage(ctx, messageDomain); err != nil {
				wbs.l.Error("Error while saving message", "error", err)
				return err
			}
		case client == nil && messageDomain.ChatId != uuid.Nil: //клиент не в сети, чат создан
			if err := wbs.cms.SaveMessage(ctx, messageDomain); err != nil {
				wbs.l.Error("Error while saving message", "error", err)
				return err
			}
		}

	}

}

func (wbs *websocketService) Write() error {

	for message := range wbs.mailing { //цикл будет работать, пока канал не будет закрыт (после чтения из канала, данные исчезают)
		client := clients.Session.LoadClient(message.RecieverId)
		messagejson, err := json.Marshal(message)
		if err != nil {
			wbs.l.Error("Error while marshalling message", "error", err)
			return err
		}
		err = client.Conn.WriteMessage(websocket.TextMessage, messagejson)
		if err != nil {
			wbs.l.Error("Error writing", "error", err)
			return err
		}
		wbs.l.Info("Success writing")
	}
	wbs.l.Info("Channel is closed.")
	return nil

}
