package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name string
	Age  int
}

func main() {
	s1 := make([]*Student, 0)

	for i := 0; i < 5; i++ {
		s1 = append(s1, &Student{
			Name: fmt.Sprintf("%s-%d", "test", i),
			Age:  i*10 + 1,
		})
	}

	bytes, err := json.Marshal(s1)
	if err != nil {
		panic(err)
	}

	s2 := make([]*Student, 0)
	err = json.Unmarshal(bytes, &s2)
	if err != nil {
		panic(err)
	}

	for _, s := range s2 {
		fmt.Printf("%+v\n", s)
	}
}
