package main

import "fmt"

type Student struct {
	name string
	age  int
}

func main() {
	s1 := &Student{name: "test", age: 21}

	fmt.Println(*s1)

	// s2 := &Student{}
	s2 := new(Student)
	*s2 = *s1 // 相当于赋值
	s2.age = 22

	fmt.Println(*s1)
	fmt.Println(*s2)
}
