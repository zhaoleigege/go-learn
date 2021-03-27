package main

import (
	"fmt"
	"time"

	"github.com/Terry-Mao/goconf"
)

type TestConfig struct {
	Id     int            `goconf:"core:id"`
	Col    string         `goconf:"core:col"`
	Ignore int            `goconf:"-"`
	Arr    []string       `goconf:"core:arr:,"`
	Test   time.Duration  `goconf:"core:t_1:time"`
	Buf    int            `goconf:"core:buf:memory"`
	M      map[int]string `goconf:"core:m:,"`
}

func main() {
	conf := goconf.New()
	if err := conf.Parse("./test_conf.txt"); err != nil {
		panic(err)
	}

	core := conf.Get("core")
	if core == nil {
		panic("该元素不存在")
	}

	id, err := core.Int("id")
	if err != nil {
		panic(err)
	}

	fmt.Println(id)

	tf := &TestConfig{}

	if err := conf.Unmarshal(tf); err != nil {
		panic(err)
	}

	fmt.Println(tf)
}
