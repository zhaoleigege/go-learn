package memory

import "github.com/zhaoleigege/mq"

func NewMemoryBroker() mq.Broker {
	return &meBroker{SubMap: make(map[string][]chan<- mq.ReceiveEnvelop)}
}

type meBroker struct {
	// SubMap topic对应的消费者有哪些
	SubMap map[string][]chan<- mq.ReceiveEnvelop
}

func (b *meBroker) Publish(topic string, sender <-chan mq.SendEnvelop) {
	if sender == nil {
		return
	}

	go func() {
		var data mq.SendEnvelop
		select {
		case data = <-sender:
			for _, sub := range b.SubMap[topic] {
				sub <- &mq.Receiver{Data: data.Body()}
			}
		}

		data.Error() <- nil
	}()
}

func (b *meBroker) Subscribe(topic string) <-chan mq.ReceiveEnvelop {
	subs := b.SubMap[topic]
	if len(subs) <= 0 {
		subs = make([]chan<- mq.ReceiveEnvelop, 0)
	}

	receiveChan := make(chan mq.ReceiveEnvelop)

	subs = append(subs, receiveChan)
	b.SubMap[topic] = subs

	return receiveChan
}
