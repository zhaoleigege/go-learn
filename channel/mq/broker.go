package mq

// SendEnvelop 生产者的包装
type SendEnvelop interface {
	// 发送的消息
	Body() []byte
	// 发送遇到的错误
	Error() chan<- error
}

// ReceiveEnvelop 接收者的包装
type ReceiveEnvelop interface {
	// 可以接收的消息
	Body() []byte
}

// Broker 管理消息的生产和发送
type Broker interface {
	// Publish 注册一个消息发送者，需要传入一个消息生产者channel
	Publish(topic string, sender <-chan SendEnvelop)
	// Subscribe 注册一个消息消费者，返回一个消息接收者channel
	Subscribe(topic string) <-chan ReceiveEnvelop
}

type Sender struct {
	Data    []byte
	ErrChan chan error
}

func (s *Sender) Body() []byte {
	return s.Data
}

func (s *Sender) Error() chan<- error {
	return s.ErrChan
}

type Receiver struct {
	Data []byte
}

func (r *Receiver) Body() []byte {
	return r.Data
}
