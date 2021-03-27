package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func chanErr() (chan map[int]string, chan error) {
	resMap := make(chan map[int]string, 1)
	resErr := make(chan error, 1)

	resErr <- errors.Errorf("错误")
	resMap <- nil
	return resMap, resErr
}

func main() {
	rMap, err := chanErr()

	if <-err != nil {
		fmt.Println(<-rMap)
	}

}
