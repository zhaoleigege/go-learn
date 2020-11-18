// 一个生产者多个接收者的情况 如何关闭channel
package channel

import (
	"math/rand"
	"time"
)

// Producer 模拟生产者
// chan<- int 代表只能够写入
func Producer(send chan<- int) {
	rand.Seed(time.Now().Unix())

	for {
		count := rand.Int() % 100
		if count < 5{
			close(send)
			return
		}

		send <- count
		time.Sleep(100 * time.Millisecond)
	}
}

func Consumer(receive <-chan int) {
	
}