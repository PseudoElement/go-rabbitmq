package rabbit

type ReceiveHandler func(msg []byte) error
