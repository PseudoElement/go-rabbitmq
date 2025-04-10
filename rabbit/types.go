package rabbit

type ReceiveHandler func(msg []byte) error

type RMQ_QueueParams struct {
	QueueName string
	/* direct, fanout, topic, headers */
	ExchangeKind string
	ExchangeName string
	/* Required for `topic` exchanges */
	RoutingKey string
}

type RMQ_SendParams struct {
	RMQ_QueueParams
	Msg any
}
