package main

import "fmt"

type Action interface {
	Say() string
}

type Student struct {
	Name string
	Age  int
}

func (s *Student) Say() string {
	return fmt.Sprintf("%s - %d", s.Name, s.Age)
}

type ExStudent struct {
	Student
}

func (ex *ExStudent) Say() string {
	doAction(&ex.Student)
	return fmt.Sprintf("%s", ex.Name)
}

func doAction(action Action) {
	fmt.Println(action.Say())
}

func main() {
	s := &ExStudent{
		Student: Student{
			Name: "test",
			Age:  21,
		},
	}

	doAction(s)
}
