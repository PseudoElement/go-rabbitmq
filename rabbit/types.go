package rabbit

type ReceiveHandler func(msg []byte) error

type RMQ_QueueParam struct {
	QueueName string
	/* direct, fanout, topic, headers*/
	ExchangeKind string
	ExchangeName string
}

type RMQ_SendParam struct {
	RMQ_QueueParam
	Msg any
}
