package main

import (
	"log"
	"os"
	"time"

	"github.com/pseudoelement/go-rabbitmq/rabbit"
)

func main() {
	time.Sleep(4 * time.Second)
	rmq := rabbit.NewRabbitMQ()

	queueName := "queue_2"
	exchangeKind := os.Getenv("EXCHANGE_KIND")

	rmq.CreateExchange(exchangeKind, "test-logs")
	rmq.CreateQueue(queueName)
	rmq.BindQueue(rabbit.RMQ_QueueParams{
		QueueName:    queueName,
		ExchangeKind: exchangeKind,
		ExchangeName: "test-logs",
	})

	rmq.Listen(queueName, func(msgBytes []byte) error {
		log.Printf("Consumer-2 received a message: %s", msgBytes)
		return nil
	},
	)

	select {}
}
