// 监听者模式的实现
package main

import (
	"bufio"
	"fmt"
	"os"
)

type Dog struct {
	Name    string
	Sitters map[string][]chan string
}

func (d *Dog) AddSitter(name string, ch chan string) {
	if d.Sitters == nil {
		d.Sitters = make(map[string][]chan string)
	}

	d.Sitters[name] = append(d.Sitters[name], ch)
}

func (d *Dog) Emit(name string, response string) {
	for _, ch := range d.Sitters[name] {
		go func(handler chan string) {
			fmt.Println("通道开始: ")
			select {
			case handler <- response:
				fmt.Println("通道消耗了数据: ", response)
			default:
				fmt.Println("没有向通道内写数据")
			}
		}(ch)
	}
}

func main() {
	d := &Dog{}
	ch1 := make(chan string)
	d.AddSitter("test", ch1)

	ch2 := make(chan string)
	d.AddSitter("test", ch2)

	go func() {
		for {
			msg := <-ch1
			fmt.Println("ch1接收到: ", msg)
		}
	}()

	go func() {
		for {
			msg := <-ch2
			fmt.Println("ch2接收到: ", msg)
		}
	}()

	fmt.Println("协程数据1")
	d.Emit("test", "协程数据1")

	fmt.Println("协程数据2")
	d.Emit("test", "协程数据2")

	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		d.Emit("test", input.Text())
	}
}
