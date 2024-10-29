package internal

import (
	"context"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQClient struct {
	conn *amqp.Connection // connection used by the client (tcp connection)
	ch   *amqp.Channel    // channel is used to process / send message
}

func ConnectRabbitMQ(username, password, host, vhost string) (*amqp.Connection, error) {
	return amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s/%s", username, password, host, vhost))
}

func NewRabbitMQClient(conn *amqp.Connection) (RabbitMQClient, error) {
	ch, err := conn.Channel()
	if err != nil {
		return RabbitMQClient{}, err
	}

	return RabbitMQClient{
		conn: conn,
		ch:   ch,
	}, nil
}

func (r RabbitMQClient) Close() error {
	return r.ch.Close()
}

func (r RabbitMQClient) CreateQueue(name string, durable, autodelete bool) error {
	_, err := r.ch.QueueDeclare(name, durable, autodelete, false, false, nil)
	return err
}

func (r RabbitMQClient) CreateBinding(name, binding, exchange string) error {
	return r.ch.QueueBind(name, binding, exchange, false, nil)
}

func (r RabbitMQClient) Send(ctx context.Context, exchange, routingKey string, options amqp.Publishing) error {
	return r.ch.PublishWithContext(ctx, exchange, routingKey, true, false, options)
}

func (r RabbitMQClient) Consume(queue, consumer string, autoAck bool) (<-chan amqp.Delivery, error) {
	return r.ch.Consume(queue, consumer, autoAck, false, false, false, nil)
}
