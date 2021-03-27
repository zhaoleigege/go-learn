package main

import (
	"fmt"
)

func main() {
	arr := make([]int, 5)
	for i := range arr {
		arr[i] = i
	}

	fmt.Println(arr)

	change(arr)
	fmt.Println(arr)
}

func change(arr []int) {
	for i := range arr {
		arr[i] = 5
	}
}
