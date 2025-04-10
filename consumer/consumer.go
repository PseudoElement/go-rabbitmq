package main

import (
	"log"
	"os"
	"time"

	"github.com/pseudoelement/go-rabbitmq/rabbit"
)

// RabbitMQ `topic` exchange docs https://www.rabbitmq.com/tutorials/tutorial-five-go
func main() {
	time.Sleep(4 * time.Second)
	rmq := rabbit.NewRabbitMQ()

	queueName := "queue_1"
	exchangeKind := os.Getenv("EXCHANGE_KIND")

	rmq.CreateExchange(exchangeKind, "test-logs")
	rmq.CreateQueue(queueName)
	rmq.BindQueue(rabbit.RMQ_QueueParams{
		QueueName:    queueName,
		ExchangeKind: exchangeKind,
		ExchangeName: "test-logs",
		RoutingKey:   "payment.card.*",
	})

	rmq.Listen(queueName, func(msgBytes []byte) error {
		log.Printf("Consumer-1 received a message: %s", msgBytes)
		return nil
	},
	)

	select {}
}
