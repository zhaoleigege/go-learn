package main

import (
	"fmt"
	"runtime"
)

func CallerName() string {
	if pc, file, line, ok := runtime.Caller(1); ok {
		fmt.Printf("file: %s, line: %d", file, line)
		return runtime.FuncForPC(pc).Name()
	}

	return ""
}

func CallersName() []string {
	result := make([]string, 0)

	pc := make([]uintptr, 10)
	n := runtime.Callers(3, pc)
	for i := 0; i < n; i++ {
		f := runtime.FuncForPC(pc[i])
		file, line := f.FileLine(pc[i])
		result = append(result, fmt.Sprintf("%s:%d %s\n", file, line, f.Name()))
	}

	return result
}

func t1() []string {
	return CallersName()
}

func t2() []string {
	return t1()
}

func T() []string {
	result := CallersName()
	fmt.Println(result)
	return result
}

func main() {
	//CallerName()

	//T()

	fmt.Println(t2())
}
