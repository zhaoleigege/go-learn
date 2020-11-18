package memory

import (
	"fmt"
	"github.com/zhaoleigege/mq"
	"sync"
	"testing"
)

func TestBroker(t *testing.T) {
	broker := NewMemoryBroker()
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go producer(wg, broker)

	wg.Add(1)
	go consumer(wg, broker)

	wg.Wait()
	fmt.Println("主流程结束")
}

func producer(wg *sync.WaitGroup, broker mq.Broker) {
	defer wg.Done()

	errChan := make(chan error)
	sender := make(chan mq.SendEnvelop)

	broker.Publish("test", sender)

	sender <- &mq.Sender{
		Data:    []byte("生产者测试"),
		ErrChan: errChan,
	}

	if err := <-errChan; err != nil{
		panic("生产者错误")
	}

	fmt.Println("生产者结束")
}

func consumer(wg *sync.WaitGroup, broker mq.Broker) {
	defer wg.Done()

	receiveChan := broker.Subscribe("test")
	data := <-receiveChan

	fmt.Printf("消费者收到:%s\n", string(data.Body()))
}
