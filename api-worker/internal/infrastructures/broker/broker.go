package broker

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type Broker interface {
	GetChannel() *amqp.Channel
}

type broker struct {
	channel *amqp.Channel
}

func (b *broker) GetChannel() *amqp.Channel {
	return b.channel
}

func GetBroker(brokerURI, exchangeName, queueName, routingKey string) (Broker, error) {

	conn, err := amqp.Dial(brokerURI)
	if err != nil {
		return nil, queueName, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, queueName, err
	}

	err = ch.ExchangeDeclare(
		exchangeName, // name
		"topic",      // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		return nil, queueName, err
	}

	q, err := ch.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)

	if err != nil {
		return nil, queueName, err
	}

	err = ch.QueueBind(
		q.Name,       // queue name
		routingKey,   // routing key
		exchangeName, // exchange
		false,
		nil)
	if err != nil {
		return nil, queueName, err
	}

	return &broker{
		channel: ch,
	}, err
}
