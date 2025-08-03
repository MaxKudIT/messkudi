package rabbitmq

func (rmq *rabbitmq) Setup(queueName string) error {
	if _, err := rmq.newQueue(queueName); err != nil {
		return err
	}
	if err := rmq.newExchange(); err != nil {
		return err
	}
	if err := rmq.binding(queueName); err != nil {
		return err
	}
	return nil
}
