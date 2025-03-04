package rabbit

type ReceiveHandler func(msg []byte) error

type RMQ_QueueParams struct {
	QueueName string
	/* direct, fanout, topic, headers */
	ExchangeKind string
	ExchangeName string
}

type RMQ_SendParams struct {
	RMQ_QueueParams
	Msg any
}
