package main

import (
	"fmt"
	"time"
)

func getResult(result chan string) {
	time.Sleep(time.Duration(2) * time.Second)
	result <- "结束"
}

func main() {
	result := make(chan string, 1)
	go func() {
		time.Sleep(time.Duration(2) * time.Second)
		result <- "结束"
	}()

	// go getResult(result)
	fmt.Println("开始输出")
	fmt.Println(<-result)

}
