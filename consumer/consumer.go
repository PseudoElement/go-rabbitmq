package main

import (
	"log"
	"time"

	"github.com/pseudoelement/go-rabbitmq/rabbit"
)

func main() {
	time.Sleep(4 * time.Second)
	rmq := rabbit.NewRabbitMQ()

	rmq.CreateQueue("consumer")
	rmq.Listen("consumer", func(msgBytes []byte) error {
		log.Printf("Consumer-1 received a message: %s", msgBytes)
		return nil
	},
	)

	select {}
}
