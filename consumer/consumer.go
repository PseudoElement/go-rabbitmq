package main

import (
	"time"

	"github.com/pseudoelement/go-rabbitmq/rabbit"
)

func main() {
	time.Sleep(4 * time.Second)
	rmq := rabbit.NewRabbitMQ()

	rmq.CreateQueue("greeter")
	rmq.Listen("greeter", nil)

	select {}
}
