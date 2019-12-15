package main

import (
	"fmt"
	"io"
)

type Student struct {
	Name string
}

func (s Student) String() string {
	return s.Name
}

func (stu Student) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			_, _ = io.WriteString(s, stu.Name + "：uuu")
			return
		}
		fallthrough
	case 's':
		_, _ = io.WriteString(s, stu.Name + "iiii")
	case 'q':
		_, _ = fmt.Fprintf(s, "%q", stu.Name + "aaa")
	}
}

func main() {
	s := &Student{"托尔斯泰"}
	fmt.Println(s)
	fmt.Printf("%+v\n", s)
}
