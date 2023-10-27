package rabbitmq

import (
	"fmt"

	"github.com/streadway/amqp"
	"kimpton.io/url-checker/internal/config"
	"kimpton.io/url-checker/internal/domain"
)

var _ domain.MessageQueue = &rabbitMQ{}

type rabbitMQ struct {
	conn *amqp.Connection
}

func Connect(conf config.RabbitMQ) (domain.MessageQueue, error) {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/",
		conf.User,
		conf.Password,
		conf.Host,
		conf.Port))
	if err != nil {
		return nil, err
	}

	return &rabbitMQ{
		conn: conn,
	}, nil
}

func (r *rabbitMQ) Publish(name domain.QueueName, data []byte) error {

	queueName := string(name)

	ch, err := r.conn.Channel()
	defer ch.Close()
	if err != nil {
		return err
	}
	_, err = ch.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		return err
	}
	return ch.Publish("", queueName, false, false, amqp.Publishing{
		ContentType: "application/octet-stream",
		Body:        data,
	})
}

func (r *rabbitMQ) Consume(name domain.QueueName) (<-chan []byte, error) {

	queueName := string(name)

	ch, err := r.conn.Channel()
	if err != nil {
		return nil, err
	}

	// TODO: Limit the number of items consumed at a time
	msgs, err := ch.Consume(queueName, "", true, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	out := make(chan []byte)
	go func() {
		for msg := range msgs {
			out <- msg.Body
		}
		close(out)
	}()
	return out, nil
}
