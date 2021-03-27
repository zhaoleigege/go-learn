package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type TimeStr struct {
	Name string
}

type TimeStruct struct {
	Time time.Time `json:"time"`
}

func main() {
	fmt.Println("打印当地时间、UTC时间和RFC3339表示的时间")
	t := time.Now()
	fmt.Println(t)
	fmt.Println(t.UTC())
	fmt.Println(t.Format(time.RFC3339))

	fmt.Println("打印格式化的时间")
	format := "2006-01-02 15:04:05 -0700 Mon"
	fmt.Println(t.Format(format))

	fmt.Println("打印格式化的时间")
	format = "2006-01-02 03:04:05PM MST Mon"
	fmt.Println(t.Format(format))

	fmt.Println("解析传过来的不带时区的字符串时间为CST时间")
	// 获得系统的时区，在中国为CST即+0800
	loc, err := time.LoadLocation("Local")
	if err != nil {
		panic(err)
	}
	format = "2006-1-02 15:04:05"
	t, err = time.ParseInLocation(format, "2019-12-10 23:34:59", loc)
	if err != nil {
		panic(err)
	}
	fmt.Println(t)
	fmt.Println(t.UTC())

	fmt.Println("解析传过来的带时区的字符串时间")
	format = "2006-01-02T15:04:05-07:00"
	t, err = time.Parse(format, "2019-09-21T16:59:50.646848+08:00")
	if err != nil {
		panic(err)
	}
	fmt.Println(t)
	fmt.Println(t.UTC())

	fmt.Println("小练习：")
	str := "2016-12-04 15:39:06 +0800 CST"
	format = "2006-01-02 15:04:05 -0700 MST"
	t, err = time.Parse(format, str)
	fmt.Println(t)
	fmt.Println(t.UTC())

	format = "2006-01-02T15:04:05-07:00"
	t1, err := time.Parse(format, "2019-09-21T16:59:50.646848+08:00")
	if err != nil {
		panic(err)
	}
	fmt.Printf("t1 = %s\n", t1)
	t2, err := time.Parse(format, "2019-09-21T16:59:50+08:00")
	if err != nil {
		panic(err)
	}
	fmt.Printf("t2 = %s\n", t2)

	fmt.Println(t1 == t2)
	fmt.Println(t1.Format(format) == t2.Format(format))

	timestruct := &TimeStruct{
		Time: time.Now(),
	}

	result, err := json.Marshal(timestruct)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(result))

	// orderDate := time.Date(2017, 11, 1, 18, 30, 0, 0, time.FixedZone("CST", 8*60*60))
	// fmt.Println(orderDate.Format("2006-01-02"))

	// test := &TimeStr{Name: "test"}
	// fmt.Printf("%v", test)
}
