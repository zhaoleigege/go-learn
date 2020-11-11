package main

import (
	"fmt"
	"runtime"
)

func main() {
	testCaller()
}

func testCaller() {
	fmt.Println("调用testCaller方法")

	// 打印该方法(testCaller)的调用堆栈
	callersTrace()
}

// callerTrace 打印函数调用栈某一层的信息
func callerTrace() {
	// Caller参数为0时代表当前函数(runtime.Caller)被调用的信息
	// 为1是代表调用callerTrace()函数的信息
	pc, file, line, ok := runtime.Caller(1)
	fmt.Printf("pc: %v, file: %s, line: %d, ok: %t\n", pc, file, line, ok)

	// 通过runtime.FuncForPC传入pc可以获取函数自生相关的信息
	fc := runtime.FuncForPC(pc)
	fmt.Printf("funcName: %s\n", fc.Name())
}

// callersTrace 返回函数调用栈的程序计数器，并放入到一个uintptr数组中
// 这里可以一直返回到main调用的地方
func callersTrace() {
	pc := make([]uintptr, 10) // at least 1 entry needed
	n := runtime.Callers(2, pc)
	for i := 0; i < n; i++ {
		// 由于内联函数的原因，需要把程序计数器-1才能返回相关函数真正的信息
		f := runtime.FuncForPC(pc[i] - 1)
		file, line := f.FileLine(pc[i] - 1)
		fmt.Printf("%s:%d %s\n", file, line, f.Name())
	}
}

func callersFramesTrace() {
	pc := make([]uintptr, 10)
	n := runtime.Callers(0, pc)
	pc = pc[:n]

	frames := runtime.CallersFrames(pc)
	for {
		frame, again := frames.Next()
		if !again {
			break
		}

		fmt.Printf(" %s:%d %s\n", frame.File, frame.Line, frame.Function)
	}
}
