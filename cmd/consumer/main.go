package main

import (
	"goeda/internal"
	"log"
)

func main() {

	conn, err := internal.ConnectRabbitMQ("jian", "123456", "localhost:5672", "customers")
	if err != nil {
		panic(err)
	}

	mqClient, err := internal.NewRabbitMQClient(conn)
	if err != nil {
		panic(err)
	}

	messageBus, err := mqClient.Consume("customers_created", "email-service", true)
	if err != nil {
		panic(err)
	}

	// blocking is used to block forever
	blocking := make(chan struct{})

	go func() {
		for message := range messageBus {
			// breakpoint here
			log.Printf("New Message: %v", message)
			// Multiple means that we acknowledge a batch of messages, leave false for now
			if err := message.Ack(false); err != nil {
				log.Printf("Acknowledged message failed: Retry ? Handle manually %s\n", message.MessageId)
				continue
			}
			log.Printf("Acknowledged message %s\n", message.MessageId)
		}
	}()

	log.Println("Consuming, to close the program press CTRL+C")
	// This will block forever
	<-blocking

}
