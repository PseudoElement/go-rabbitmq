package rabbit

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	ch     *amqp.Channel
	queues map[string]amqp.Queue
}

func NewRabbitMQ() *RabbitMQ {
	rmq := &RabbitMQ{}
	err := rmq.run()
	failOnError(err, "RabbitMQ run error: ")

	return rmq
}

func (this *RabbitMQ) run() error {
	conn, err := amqp.Dial(getRabbitMqUrl())
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	this.ch = ch

	return nil
}

func (this *RabbitMQ) CreateQueue(name string) error {
	queue, err := this.ch.QueueDeclare(
		name,  // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	this.queues[name] = queue

	return err
}

func (this *RabbitMQ) Send(queueName string, msg any) error {
	queue, ok := this.queues[queueName]
	if !ok {
		return fmt.Errorf("Invalid queue name: %s.", queueName)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	byteMsg, err := json.Marshal(msg)
	failOnError(err, "Marshal error: ")

	err = this.ch.PublishWithContext(ctx,
		"",         // exchange
		queue.Name, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        byteMsg,
		})

	return err
}

func (this *RabbitMQ) Listen(queueName string, onMsg ReceiveHandler) error {
	queue, ok := this.queues[queueName]
	if !ok {
		return fmt.Errorf("Invalid queue name: %s.", queueName)
	}

	msgs, err := this.ch.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	failOnError(err, "Failed to register a consumer")

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			if onMsg != nil {
				err := onMsg(d.Body)
				failOnError(err, "RabbitMQ listen error: ")
			}
		}
	}()

	return nil
}
