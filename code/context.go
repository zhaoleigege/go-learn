package main

import "context"

import "fmt"

import "time"

func watch(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("err: %s\n", ctx.Err().Error())
			return
		}
	}
}

func main() {
	ctx, _ := context.WithCancel(context.Background())
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	go watch(ctx)

	time.Sleep(2 * time.Second)
	cancel()
	fmt.Println("结束")
}
