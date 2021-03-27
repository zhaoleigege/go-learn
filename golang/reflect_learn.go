package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func sprintf(x interface{}) string {
	type stringer interface {
		String() string
	}

	switch x := x.(type) {
	case stringer:
		return x.String()
	case string:
		return x
	case int:
		return strconv.Itoa(x)
	case bool:
		if x {
			return "true"
		} else {
			return "false"
		}
	default:
		return "???"
	}
}

func main() {
	age := 20

	ageType := reflect.TypeOf(age)
	fmt.Println(ageType.String())

	ageValue := reflect.ValueOf(age)
	fmt.Println(ageValue.Kind() == reflect.Uint)
}
