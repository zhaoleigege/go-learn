package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2) // 因为有两个动作，所以增加2个计数
	go func() {
		fmt.Println("Goroutine 1")
		wg.Done() // 操作完成，减少一个计数
	}()

	go func() {
		fmt.Println("Goroutine 2")
		wg.Done() // 操作完成，减少一个计数
	}()

	wg.Wait()
}
