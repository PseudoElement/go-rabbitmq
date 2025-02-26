package main

import (
	"time"

	"github.com/pseudoelement/go-rabbitmq/rabbit"
)

type Msg struct {
	Sender string `json:"sender"`
	Text   string
}

func main() {
	time.Sleep(2 * time.Second)

	rmq := rabbit.NewRabbitMQ()
	rmq.CreateQueue("checker")

	msg := Msg{
		Sender: "main.go",
		Text:   "Message from main service.",
	}

	err := rmq.Listen("checker", nil)
	if err != nil {
		panic(err)
	}

	rmq.Send("checker", msg)
	rmq.Send("checker", msg)
	rmq.Send("checker", msg)

	select {}
}
