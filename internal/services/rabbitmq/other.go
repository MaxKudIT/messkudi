package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
)

func (rmq *rabbitmq) newQueue(name string) (amqp.Queue, error) {
	if _, err := rmq.set.pub.QueueDeclarePassive(
		name,
		true,
		false,
		false,
		false,
		nil,
	); err == nil {
		return amqp.Queue{Name: name}, nil
	}

	q, err := rmq.set.pub.QueueDeclare(
		name,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		rmq.l.Error("Queue creation error", "error", err, "queue", name)
		return amqp.Queue{}, fmt.Errorf("queue declare failed: %w", err)
	}
	return q, nil
}

func (rmq *rabbitmq) newExchange() error {
	if err := rmq.set.pub.ExchangeDeclarePassive(
		rmq.setNav.exchangeName,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	); err == nil {
		return nil
	}

	// Создаем exchange
	err := rmq.set.pub.ExchangeDeclare(
		rmq.setNav.exchangeName,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		rmq.l.Error("Exchange creation error", "error", err, "exchange", rmq.setNav.exchangeName)
		return fmt.Errorf("exchange declare failed: %w", err)
	}
	return nil
}

func (rmq *rabbitmq) binding(queueName string) error {
	_, err := rmq.set.pub.QueueInspect(queueName)
	if err == nil {
		return nil
	}

	err = rmq.set.pub.QueueBind(
		queueName,
		rmq.setNav.routingKey,
		rmq.setNav.exchangeName,
		false,
		nil,
	)
	if err != nil {
		rmq.l.Error("Binding error",
			"error", err,
			"queue", queueName,
			"exchange", rmq.setNav.exchangeName,
			"routingKey", rmq.setNav.routingKey)
		return fmt.Errorf("queue bind failed: %w", err)
	}
	return nil
}
