// 多个生产者一个接收者的情况 如何关闭channel
// 这里是接收者发送关闭的消息
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Producer 模拟生产者
// chan<- int 代表只能够写入
func Producer(wg *sync.WaitGroup, send chan<- int, stopChan <-chan struct{}) {
	defer wg.Done()

	rand.Seed(time.Now().Unix())

	for {
		// 这里先判断是否需要停止，是为了尽可能早的退出goroutine
		select {
		case <-stopChan:
			return
		default:

		}

		count := rand.Int() % 100

		// 这里会出现一个情况，在进行上面的业务逻辑时，可能stopChan已经关闭了，但是send这个channel还可以往里面写数据
		// 因为select是随机选择一个可以执行的chan
		select {
		case <-stopChan:
			return
		case send <- count:
		}

		time.Sleep(100 * time.Millisecond)
	}
}

// Consumer 模拟消费者
// <-chan int 代表只能接收消息
func Consumer(receive <-chan int, stopChan chan<- struct{}) {
	for data := range receive {
		if data < 5 {
			// 发送关闭信号
			close(stopChan)
			return
		}

		fmt.Println(data)
	}
}

// 适用于有缓冲或者无缓冲通道
// 这里的data channel没有手动关闭，但是也不会造成内存泄露，因为golang会自动回收不再使用的channel
func main() {
	// 生产者和消费者之间传递数据的channel
	data := make(chan int, 10)
	// 关闭消息的channel
	stop := make(chan struct{})

	// 创建生产者
	wg := &sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go Producer(wg, data, stop)
	}

	// 创建消费者
	go Consumer(data, stop)

	wg.Wait()
}
