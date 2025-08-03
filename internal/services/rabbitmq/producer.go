package rabbitmq

import (
	"encoding/json"
	"github.com/MaxKudIT/messkudi/internal/transport/web/dto/chat_message_dto"
	"github.com/streadway/amqp"
	"time"
)

func (rmq *rabbitmq) ProduceMessage(data chat_message_dto.ChatMessageDTOClient) error {
	datajson, err := json.Marshal(data)
	if err != nil {
		rmq.l.Error("Error marshalling message", data, err)
		return err
	}

	if err := rmq.set.pub.Publish(rmq.setNav.exchangeName, rmq.setNav.routingKey, false, false, amqp.Publishing{
		ContentType:  "application/json",
		Body:         datajson,
		DeliveryMode: amqp.Persistent,
		Timestamp:    time.Now().UTC(),
	}); err != nil {
		rmq.l.Error("Error publishing message", data, err)
		return err
	}
	return nil
}
