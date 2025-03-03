package rabbit

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	ch     *amqp.Channel
	queues map[string]amqp.Queue
}

func NewRabbitMQ() *RabbitMQ {
	rmq := &RabbitMQ{queues: make(map[string]amqp.Queue)}
	err := rmq.run()
	failOnError(err, "RabbitMQ run error: ")

	return rmq
}

func (this *RabbitMQ) run() error {
	conn, err := amqp.Dial(getRabbitMqUrl())
	failOnError(err, "Failed to connect to RabbitMQ: ")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel: ")

	this.ch = ch

	return nil
}

func (this *RabbitMQ) CreateQueue(name string) error {
	queue, err := this.ch.QueueDeclare(
		name,  // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	this.queues[name] = queue

	return err
}

func (this *RabbitMQ) BindQueue(p RMQ_QueueParam) error {
	var routingKey string
	if p.ExchangeKind == "fanout" {
		routingKey = ""
	} else {
		routingKey = p.QueueName
	}

	err := this.ch.QueueBind(
		p.QueueName,    // queue name
		routingKey,     // routing key
		p.ExchangeName, // exchange
		false,
		nil,
	)
	failOnError(err, "Failed to bind a queue")

	return nil
}

/*
* @param kind - direct(1 msg to 1 queue), fanout(1 msg to all queue with same name)
 */
func (this *RabbitMQ) CreateExchange(kind string, name string) error {
	err := this.ch.ExchangeDeclare(
		name,  // name
		kind,  // type
		true,  // durable
		false, // auto-deleted
		false, // internal
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to ExchangeDeclare: ")

	return nil
}

func (this *RabbitMQ) Send(p RMQ_SendParam) error {
	_, ok := this.queues[p.QueueName]
	if !ok {
		return fmt.Errorf("Invalid queue name: %s.", p.QueueName)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	byteMsg, err := json.Marshal(p.Msg)
	failOnError(err, "Marshal error: ")

	// routing key is "" if `fanout` exchange used
	// routing key is queueName if `direct` exchange used
	var routingKey string
	if p.ExchangeKind == "fanout" {
		routingKey = ""
	} else {
		routingKey = p.QueueName
	}

	err = this.ch.PublishWithContext(ctx,
		p.ExchangeName,
		routingKey,
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        byteMsg,
		})

	return err
}

func (this *RabbitMQ) Listen(queueName string, onMsg ReceiveHandler) error {
	_, ok := this.queues[queueName]
	if !ok {
		return fmt.Errorf("Invalid queue name: %s.", queueName)
	}

	msgs, err := this.ch.Consume(
		queueName, // queue
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	failOnError(err, "Failed to register a consumer")

	go func() {
		for d := range msgs {
			if onMsg != nil {
				err := onMsg(d.Body)
				failOnError(err, "RabbitMQ listen error: ")
			}
		}
	}()

	return nil
}
