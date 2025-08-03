package rabbitmq

import (
	"context"
	"github.com/MaxKudIT/messkudi/internal/domain/messages"
	"github.com/MaxKudIT/messkudi/internal/transport/web/dto/chat_message_dto"
	"github.com/google/uuid"
	"github.com/streadway/amqp"
	"log/slog"
)

type setT struct {
	pub *amqp.Channel
	sub *amqp.Channel
}

type setNav struct {
	exchangeName string
	routingKey   string
}

type rabbitmq struct {
	connection *amqp.Connection
	set        *setT
	setNav     *setNav
	l          *slog.Logger
	cms        ChatMessageStorage
}
type ChatMessageStorage interface {
	MessageById(ctx context.Context, id uuid.UUID) (chat_message_dto.ChatMessageDTODetailsServer, error)
	SaveMessage(ctx context.Context, message messages.ChatMessage) error
	DeleteUser(ctx context.Context, id uuid.UUID) error
}

func NewRabbitConnection(connectionstr string) (*amqp.Connection, error) {
	conn, err := amqp.Dial(connectionstr)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
func NewRabbit(connection *amqp.Connection, set *setT, setNav *setNav, cms ChatMessageStorage, l *slog.Logger) *rabbitmq {
	return &rabbitmq{connection: connection, set: set, setNav: setNav, cms: cms, l: l}
}

func NewSet(pub *amqp.Channel, sub *amqp.Channel, queue amqp.Queue) *setT {
	return &setT{pub: pub, sub: sub}
}
