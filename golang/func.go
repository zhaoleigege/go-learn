package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func action(f func()) {
	f()
}

func main() {

	name := "test"

	wg.Add(1)
	go action(func() {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(name)
		wg.Done()
	})

	fmt.Println("开始阻塞")

	wg.Wait()

	fmt.Println("执行结束")
}
