package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	wg := &sync.WaitGroup{}
	stopChan := make(chan struct{})
	pChan := make(chan []string)
	rChan := make(chan int)

	// 生产者
	go func() {
		producer(ctx, -1, pChan)
	}()

	// 消费者
	go func() {
		for i := 0; i < 5; i++ {
			(func(index int) {
				wg.Add(1)
				go cConsumer(ctx, index, pChan, rChan, wg)
			})(i)
		}
		wg.Wait()
		close(stopChan)
		fmt.Println("执行完成")
	}()

	// 处理结果
	total := 0
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-stopChan:
				cancel()
				return
			case c, ok := <-rChan:
				if !ok {
					return
				}

				total += c
			}
		}
	}()

Loop:
	for {
		select {
		case <-ctx.Done():
			break Loop
		}
	}

	fmt.Println(total)
}

func producer(ctx context.Context, index int, rp chan<- []string) {
	maxId := 100
	gap := 10

	for start := 0; start < maxId; start += gap {
		strArr := make([]string, 0, gap)
		for i := 0; i < gap; i++ {
			strArr = append(strArr, fmt.Sprintf("%d producer-%d", index, start+i))
		}
		//time.Sleep(500 * time.Microsecond)

		select {
		case <-ctx.Done():
			return
		case rp <- strArr:
			break
		}
	}

	close(rp)
}

func cConsumer(ctx context.Context, index int, rc <-chan []string, resultChan chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for strArr := range rc {
		fmt.Printf("index: %d接收\n", index)

		time.Sleep(2 * time.Second)
		select {
		case <-ctx.Done():
			return
		case resultChan <- len(strArr):
			break
		}
	}
}
