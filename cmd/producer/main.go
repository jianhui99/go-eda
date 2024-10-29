package main

import (
	"context"
	"goeda/error"
	"goeda/internal"
	"log"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := internal.ConnectRabbitMQ("jian", "123456", "localhost:5672", "customers")
	error.FailOnError(err, "fail to connect to RabbitMQ")
	defer conn.Close()

	client, err := internal.NewRabbitMQClient(conn)
	error.FailOnError(err, "fail to initializing RabbitMQ")
	defer client.Close()

	err = client.CreateQueue("customers_created", true, false)
	error.FailOnError(err, "fail to create queue")

	err = client.CreateBinding("customers_created", "customers.created.*", "customer_events")
	error.FailOnError(err, "fail to create queue binding")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = client.Send(ctx, "customer_events", "customers.created.us", amqp091.Publishing{
		ContentType:  "text.plain",
		DeliveryMode: amqp091.Persistent,
		Body:         []byte(`Hello World`),
	})
	error.FailOnError(err, "fail to send")

	log.Println(client)
}
