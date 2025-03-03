package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/pseudoelement/go-rabbitmq/rabbit"
)

type Msg struct {
	Sender string `json:"sender"`
	Text   string
}

func getConsumersNames() []string {
	names := os.Getenv("CONSUMERS_NAMES")
	if names == "" {
		panic("env CONSUMERS_NAMES is empty")
	}

	splited := strings.Split(names, "___")

	return splited
}

func main() {
	time.Sleep(3 * time.Second)
	rmq := rabbit.NewRabbitMQ()

	consumerNames := getConsumersNames()
	for _, name := range consumerNames {
		rmq.CreateQueue(name)
	}

	for i := 0; ; i++ {
		time.Sleep(2 * time.Second)
		msg := Msg{
			Sender: "main.go",
			Text:   fmt.Sprintf("%v Message from main service.", i),
		}
		for _, name := range consumerNames {
			rmq.Send(name, msg)
			fmt.Println(i, " Msg sent.")
		}
	}
}
