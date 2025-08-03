package rabbitmq

import (
	"context"
	"encoding/json"
	"github.com/MaxKudIT/messkudi/internal/transport/web/dto/chat_message_dto"
	"github.com/google/uuid"
	"log"
	"time"
)

func (rmq *rabbitmq) ConsumeMessages(queueName string) ([]chat_message_dto.ChatMessageDTOClient, error) {

	var rdata []chat_message_dto.ChatMessageDTOClient
	data, err := rmq.set.sub.Consume(queueName, "", true, false, false, false, nil)
	if err != nil {
		return nil, err
	}
	for msg := range data {
		var message chat_message_dto.ChatMessageDTOClient
		if err := json.Unmarshal(msg.Body, &message); err != nil {
			rmq.l.Error("Error unmarshaling")
			msg.Nack(false, true)
			return nil, err
		}
		rdata = append(rdata, message)
		msg.Ack(false)

	}
	return rdata, nil
}

func (rmq *rabbitmq) GetMessagesWithBroker(queueName string, ctx context.Context) {
	for {
		var messageDTO chat_message_dto.ChatMessageDTOClient
		msg, ok, err := rmq.set.pub.Get(queueName, true)
		if err != nil {
			log.Println("Ошибка при получении сообщения:", err)
			break
		}
		if !ok {
			break
		}
		if err := json.Unmarshal(msg.Body, &messageDTO); err != nil {
			rmq.l.Error("Error unmarshaling")
			msg.Nack(false, true)
		}
		messageDTOP, err := messageDTO.UuidParse()
		if err != nil {
			rmq.l.Error(err.Error())
		}
		messageDomain := chat_message_dto.ToDomain(uuid.New(), time.Now(), time.Now(), nil, messageDTOP)
		rmq.cms.SaveMessage(ctx, messageDomain)
	}
}
