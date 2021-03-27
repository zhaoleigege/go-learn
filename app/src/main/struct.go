package main

import "fmt"

type Student struct {
	Id   int
	Name string
}

func (s *Student) makePStudent(id int, name string) {
	s.Id = id
	s.Name = name
}

func (s Student) makeTStudent(id int, name string) {
	s.Id = id
	s.Name = name
	fmt.Println("=================")
	fmt.Println(s)
	fmt.Println("=================")

}

func main() {
	var student *Student

	student = new(Student)

	fmt.Println(student)

	student.makePStudent(20, "赵磊") // 值改变了

	fmt.Println(student)

	student.makeTStudent(21, "test") // 值没有改变

	fmt.Println(student)
}
