module github.com/pseudoelement/go-rabbitmq/consumer

go 1.22.3

require github.com/pseudoelement/go-rabbitmq/rabbit v0.0.0-00010101000000-000000000000

require github.com/rabbitmq/amqp091-go v1.10.0 // indirect

replace github.com/pseudoelement/go-rabbitmq/rabbit => ../rabbit
