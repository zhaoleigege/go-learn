package main

import "fmt"

import "time"

func getFunc(f func()) {
	time.Sleep(2000)
	f()
}

type Func func(id int)

func main() {
	f := func(id int) {
		fmt.Println(id)
	}

	funcs := make([]Func, 0)

	for i := 0; i < 10; i++ {
		funcs = append(funcs, f)
	}

	for i := 0; i < 10; i++ {
		go funcs[i](i)
	}

	time.Sleep(1000000000000)
}
