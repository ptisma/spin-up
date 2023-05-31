package worker

import (
	"api-worker/internal/infrastructures/broker"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Worker interface {
	StartDockerImageBuild() error
	ReadFromQueue() (<-chan amqp.Delivery, error)
}

type worker struct {
	broker broker.Broker
}

func (w *worker) StartDockerImageBuild() error {

}

func (w *worker) ReadFromQueue() (<-chan amqp.Delivery, error) {

	msgs, err := w.broker.GetChannel().Consume(
		p.QueueName, // queue
		"",          // consumer
		true,        // auto ack
		false,       // exclusive
		false,       // no local
		false,       // no wait
		nil,         // args
	)

	return msgs, err

}

func GetWorker(broker broker.Broker) (Worker, error) {
	var err error
	return &worker{broker}, err

}
