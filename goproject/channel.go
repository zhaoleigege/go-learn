package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	stopChan := make(chan struct{})
	resultChan := make(chan string, 1)

	go func() {
		for i := 0; i < 5; i++ {
			(func(index int) {
				wg.Add(1)
				go consumer(index, resultChan, wg)
			})(i)
		}
		wg.Wait()
		close(stopChan)
		fmt.Println("执行完成")
	}()

	result := make([]string, 0)
	times := 0
Loop:
	for {
		select {
		case <-stopChan:
			fmt.Println("退出循环")
			break Loop
		case str, ok := <-resultChan:
			if !ok {
				break
			}

			times++

			fmt.Printf("times-> %d, str: -> %s\n", times, str)
			if times == 5 {
				time.Sleep(10 * time.Second)
			}
			time.Sleep(2 * time.Second)
			result = append(result, str)
		}
	}

	fmt.Println("关闭channel")
	close(resultChan)
	fmt.Println(result)
}

func consumer(index int, rc chan<- string, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
	}()

	fmt.Printf("index-> %d, 执行\n", index)
	rc <- fmt.Sprintf("consumer-%d", index)
}
