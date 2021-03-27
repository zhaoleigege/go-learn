package main

import "fmt"

type Student struct {
	Name string
}

func main() {
	sMap := make(map[int]*Student)
	sMap[10] = nil

	if s, ok := sMap[10]; ok {
		fmt.Println(s.Name)
	}

}
