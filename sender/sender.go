package main

import (
	"fmt"
	"time"

	"github.com/pseudoelement/go-rabbitmq/rabbit"
)

type Msg struct {
	Sender string `json:"sender"`
	Text   string
}

func main() {
	time.Sleep(4 * time.Second)
	rmq := rabbit.NewRabbitMQ()

	rmq.CreateQueue("greeter")

	for i := 0; ; i++ {
		time.Sleep(2 * time.Second)
		msg := Msg{
			Sender: "main.go",
			Text:   fmt.Sprintf("%v Message from main service.", i),
		}
		rmq.Send("greeter", msg)

	}
}
