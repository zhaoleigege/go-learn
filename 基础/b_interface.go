package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("姓名：%s, 年龄：%d", p.Name, p.Age)
}

func main() {
	p := &Person{
		Name: "test",
		Age:  21,
	}

	fmt.Printf("%+v\n", p)
}
