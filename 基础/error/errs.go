package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func errT() error {
	return errors.New("测试错误")
}

func fn2() error {
	return errors.New("error")
}

func fn1() error {
	return errors.WithMessage(fn2(), "fn1")
}

func fn0() error {
	return errors.WithMessage(fn1(), "fn")
}

func fn() error {
	return fn0()
}

func main() {
	// fmt.Printf("打印详细错误堆栈:\n%+v", errT())
	// fmt.Printf("打印错误字符串:%v", errT())

	fmt.Printf("打印详细错误堆栈:\n%+v", fn())
	//fmt.Printf("\n打印错误字符串:%v", fn())
}
