package main

import (
	"fmt"
	"os"
	"time"

	"github.com/pseudoelement/go-rabbitmq/rabbit"
)

type Msg struct {
	Sender string `json:"sender"`
	Text   string
}

func getQueueName() string {
	name := os.Getenv("QUEUE_NAME")
	if name == "" {
		panic("env QUEUE_NAME is empty")
	}

	return name
}

func main() {
	time.Sleep(3 * time.Second)
	rmq := rabbit.NewRabbitMQ()

	exchangeKind := os.Getenv("EXCHANGE_KIND")

	// exchange is MAIN entity to send/read messages - queueName in fanout mechanizme doesn't mattes
	rmq.CreateExchange(exchangeKind, "test-logs")
	queueName := getQueueName()
	rmq.CreateQueue(queueName)

	for i := 0; ; i++ {
		time.Sleep(2 * time.Second)
		err := rmq.Send(rabbit.RMQ_SendParam{
			RMQ_QueueParam: rabbit.RMQ_QueueParam{
				QueueName:    queueName,
				ExchangeName: "test-logs",
				ExchangeKind: exchangeKind,
			},
			Msg: Msg{
				Sender: "main.go",
				Text:   fmt.Sprintf("%v Message from main service.", i),
			},
		})
		fmt.Println("Error ===> ", err)
		fmt.Println(i, " Msg sent.")
	}
}
