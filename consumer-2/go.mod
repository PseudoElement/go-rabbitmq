module github.com/pseudoelement/go-rabbitmq/consumer-2

go 1.22.3

require github.com/pseudoelement/go-rabbitmq/rabbit v0.0.0-20250226224007-f3215295ff01

replace github.com/pseudoelement/go-rabbitmq/rabbit => ../rabbit

require github.com/rabbitmq/amqp091-go v1.10.0 // indirect
