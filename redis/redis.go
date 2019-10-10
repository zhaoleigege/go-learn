package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		panic(err)
	}
	defer c.Close()

	_, err = c.Do("set", "name", "test")
	if err != nil {
		panic(err)
	}

	value, err := redis.String(c.Do("get", "name"))
	if err != nil {
		panic(err)
	}

	fmt.Println(value)
}
