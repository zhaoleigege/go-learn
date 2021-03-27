package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name string
	Age  int
}

type Teacher struct {
	Name string
	Age  int
}

func main() {
	s1 := &Student{
		Name: "t1",
		Age:  1,
	}

	s2 := &Student{
		Name: "t2",
		Age:  2,
	}

	d1, _ := json.Marshal(s1)
	d2, _ := json.Marshal(s2)

	dataArr := make([]json.RawMessage, 0)
	dataArr = append(dataArr, d1)
	dataArr = append(dataArr, d2)

	data, _ := json.Marshal(dataArr)
	fmt.Println(string(data))
	sArr := make([]*Teacher, 0)
	if err := json.Unmarshal(data, &sArr); err != nil {
		panic(err)
	}

	for _, s := range sArr {
		fmt.Printf("%+v", s)
	}
}
