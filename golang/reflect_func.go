package main

import (
	"fmt"
	"reflect"
	"runtime"
	"time"

	"./kreflect"
)

func test() {
	fmt.Println(GetFuncName(kreflect.Inner))
}

func GetFuncName(fc interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(fc).Pointer()).Name()
}

func main() {
	go func() {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println("协程结束")
	}()

	// time.Sleep(1000 * time.Millisecond)

	test()

}
