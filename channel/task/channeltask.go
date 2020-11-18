// 1 主协程启动两个协程，协程1负责发送数据给协程2，协程2负责接收并累加获得的数据。
// 2 主协程等待两个子协程退出，当主协程意外退出时通知两个子协程退出。
// 3 当发送协程崩溃和主动退出时通知接收协程也要退出，然后主协程退出
// 4 当接收协程崩溃或主动退出时通知发送协程退出，然后主协程退出。
// 5 无论三个协程主动退出还是panic，都要保证所有资源手动回收。
package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func produceData() int {
	return rand.Int() % 100
}

func Producer(ctx context.Context, send chan<- int, stopChan chan<- struct{}) {
	defer func() {
		if r := recover(); r != nil {
			select {
			case <-ctx.Done():
				return
			case stopChan <- struct{}{}:
				fmt.Println("生产者panic处理")
				return
			}
		}

		fmt.Println("生产者结束")
	}()

	for i := 0; i < 10; i++ {
		time.Sleep(500 * time.Millisecond)

		select {
		case <-ctx.Done():
			return
		case send <- produceData():
		}
	}

}

func handleData(old, sum int) int {
	panic("消费者错误")
	return old + sum
}

func Consumer(ctx context.Context, receive <-chan int, stopChan chan<- struct{}) int {
	defer func() {
		if r := recover(); r != nil {
			select {
			case <-ctx.Done():
				return
			case stopChan <- struct{}{}:
				fmt.Println("消费者panic处理")
				return
			}
		}

		fmt.Println("消费者结束")
	}()

	sum := 0
	for {
		select {
		case <-ctx.Done():
			return sum
		case data, ok := <-receive:
			if !ok {
				return sum
			}

			sum = handleData(data, sum)
		}
	}
}

func main() {
	rand.Seed(time.Now().Unix())

	ctx, cancel := context.WithCancel(context.Background())
	dataChan := make(chan int)
	stopChan := make(chan struct{}, 1)
	endChan := make(chan struct{})

	go func() {
		Producer(ctx, dataChan, stopChan)
		close(dataChan)
	}()

	go func() {
		result := Consumer(ctx, dataChan, stopChan)
		fmt.Printf("结果: %d\n", result)

		close(endChan)
	}()

	select {
	case <-stopChan:
	case <-endChan:
	}

	cancel()
	fmt.Println("主流程结束")
}
