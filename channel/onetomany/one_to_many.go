// 一个生产者多个接收者的情况 如何关闭channel
// 这里关闭的消息应该由生产者发出
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Producer 模拟生产者
// chan<- int 代表只能够写入
func Producer(send chan<- int) {
	rand.Seed(time.Now().Unix())

	for {
		count := rand.Int() % 100
		if count < 5 {
			close(send)
			return
		}

		send <- count
		time.Sleep(100 * time.Millisecond)
	}
}

// Consumer 模拟消费者
// <-chan int 代表只能接收消息
func Consumer(wg *sync.WaitGroup, receive <-chan int) {
	defer wg.Done()

	for data := range receive {
		fmt.Println(data)
	}
}

func main() {
	// 适用于有缓冲或者无缓冲通道
	data := make(chan int)

	// 创建消费者
	wg := &sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go Consumer(wg, data)
	}

	// 创建生产者
	go Producer(data)

	wg.Wait()
}
