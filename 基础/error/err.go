package main

import (
	// "errors"
	"fmt"
	"runtime"
	"bytes"
)

type stack []uintptr
type errString struct {
	s string
	*stack
}

func (err errString) Error() string {
	var buf bytes.Buffer
	for _, s := range *err.stack{
		f := runtime.FuncForPC(s)
		file, line := f.FileLine(s)
	
		buf.WriteString(fmt.Sprintf("%s:%d %s\n", file, line, f.Name()))
	}
	return buf.String()
}

func New(text string) error {
	return &errString{
		s:     text,
		stack: callers(),
	}
}

func callers() *stack {
	const depth = 32
	var pcs [depth]uintptr

	n := runtime.Callers(3, pcs[:])

	var st stack = pcs[0:n]
	return &st
}

func errT() error{
	return New("测试错误")
}

func main() {
	// err := errors.New("测试错误")
	// fmt.Println(err)

	fmt.Println(errT())
}
