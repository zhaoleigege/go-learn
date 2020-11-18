// 多个生产者多个消费者的情况
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Producer 模拟生产者
// cancel 通过该channel判断是不是停止生产数据
// data 往该channel写数据
// stop 当要停止的时候往该channel发一条消息
func Producer(wg *sync.WaitGroup, cancel <-chan struct{}, data chan<- int, stop chan<- struct{}) {
	defer wg.Done()
	rand.Seed(time.Now().Unix())

	for {
		select {
		case <-cancel:
			return
		default:
		}

		count := rand.Int() % 100

		// 退出goroutine
		if count < 5 {
			select {
			case stop <- struct{}{}:
				fmt.Println("生产者停止")
			default:
			}
			fmt.Println("生产者退出")
			return
		}

		// 往goroutine写数据
		fmt.Printf("生产: %d\n", count)

		select {
		case <-cancel:
			return
		default:
		}

		select {
		case <-cancel:
			return
		case data <- count:
		}
	}
}

// Consumer 模拟消费者
// cancel 通过该channel判断是不是停止消费数据
// data 读该channel的数据
// stop 当要停止的时候往该channel发一条消息
func Consumer(wg *sync.WaitGroup, cancel <-chan struct{}, data <-chan int, stop chan<- struct{}) {
	defer wg.Done()

	for {
		// 判断是不是停止消费
		select {
		case <-cancel:
			return
		default:

		}

		// 读取数据
		var count int
		select {
		case <-cancel:
			return
		case count = <-data:
		}

		// 发出停止消费的消息
		if count > 95 {
			select {
			case stop <- struct{}{}:
				fmt.Println("消费者停止")
			default:
			}
			fmt.Println("消费者退出")
			return
		}

		// 正常的消费
		fmt.Printf("消费: %d\n", count)
	}
}

// StopWatcher 监听是不是需要取消生产和消费的过程
func StopWatcher(cancel chan<- struct{}, stop <-chan struct{}) {
	<-stop
	close(cancel)
}
func main() {
	data := make(chan int, 10)
	stop := make(chan struct{}, 1) // stop的缓冲区得设置为1，不然后在通过select方法向stop写数据的时候，由于stop的接收者可能没有启动起来，导致一次stop信号丢失
	cancel := make(chan struct{})

	wg := &sync.WaitGroup{}

	// 生产者
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go Producer(wg, cancel, data, stop)
	}

	// 消费者
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go Consumer(wg, cancel, data, stop)
	}

	// 先休息两秒
	//time.Sleep(2 * time.Second)
	//fmt.Println("流程走到这里")
	// 监听是否取消
	go StopWatcher(cancel, stop)

	wg.Wait()
}
