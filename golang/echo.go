package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	s, sep := "", "" // 短变量声明
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}

	fmt.Println(s)

	fmt.Println(strings.Join(os.Args[1:], " "))
}
