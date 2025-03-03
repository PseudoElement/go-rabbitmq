package main

import (
	"log"
	"time"

	"github.com/pseudoelement/go-rabbitmq/rabbit"
)

func main() {
	time.Sleep(4 * time.Second)
	rmq := rabbit.NewRabbitMQ()

	rmq.CreateQueue("consumer-2")
	rmq.Listen("consumer-2", func(msgBytes []byte) error {
		log.Printf("Consumer-2 received a message: %s", msgBytes)
		return nil
	},
	)

	select {}
}
