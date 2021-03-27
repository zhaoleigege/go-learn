package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func timer() {
	input := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			input <- i
		}
	}()

	t1 := time.NewTimer(time.Second * 10)
	t2 := time.NewTimer(time.Second * 20)

	for {
		select {
		case msg := <-input:
			fmt.Println(msg)
		case <-t1.C:
			fmt.Println("10秒定时任务")
			t1.Reset(time.Second * 10)
		case <-t2.C:
			fmt.Println("20秒定时任务")
			t2.Reset(time.Second * 1)
		}
	}
}

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)

	go timer()

	for {
		s := <-c
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT: // 程序退出时执行的指令
			fmt.Println("程序退出")
			return
		case syscall.SIGHUP:
		// TODO reload
		default:
			return
		}
	}
}
