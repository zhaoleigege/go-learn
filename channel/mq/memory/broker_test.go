package memory

import (
	"fmt"
	"github.com/zhaoleigege/mq"
	"sync"
	"testing"
	"time"
)

// go test -cpu=1,9,55,99 -race -count=1000 -failfast
// 使用竞态测试
// cpu指定GOMAXPROCS数量为1 9 55 99
// -race指定每次跑测试时都使用不同的cpu数量
// count指定测试的次数
// failfast指定一旦出错就立马停止
func TestBroker(t *testing.T) {
	broker := NewMemoryBroker()
	wg := &sync.WaitGroup{}

	// 这里可能存在生产者已经生产完消息但是消费者还没有创建的情况
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

	if err := <-errChan; err != nil {
		panic("生产者错误")
	}

	fmt.Println("生产者结束")
}

func consumer(wg *sync.WaitGroup, broker mq.Broker) {
	defer wg.Done()

	receiveChan := broker.Subscribe("test")
	var data mq.ReceiveEnvelop

	timer := time.NewTimer(1 * time.Second)
	select {
	case <-timer.C:
	case data = <-receiveChan:
	}

	if data == nil {
		fmt.Printf("消费者没有收到数据\n")
	} else {
		fmt.Printf("消费者收到:%s\n", string(data.Body()))

	}
}
