package main

import (
	"fmt"
	"strings"
)

type Student struct {
	Name string
}

func getStudent(s **Student) {
	fmt.Printf("%p\n", s)
}

func main() {
	s := &Student{
		Name: "test",
	}

	fmt.Printf("%p\n", s)
	getStudent(&s)

	str := ""
	strArr := strings.Split(str, ",")
	for i, s := range strArr {
		fmt.Printf("%d: %s\n", i, s)
	}
}
