package main

import (
	"fmt"
	"time"
)

func main() {
	tm := time.Unix(time.Now().Unix(), 0).In(time.UTC).Format("2006-01-02 15:04:05")
	fmt.Println(tm)

	// 解析的时间字符串为给定的时区
	loc, _ := time.LoadLocation("Asia/Shanghai")
	t, err := time.ParseInLocation("2006-01-02 15:04:05", "2018-09-10 00:00:00", loc)
	if err != nil {
		panic(err)
	}
	fmt.Println(t)

	// 可以通过fixZone自定义Location
	eightZone := int((8 * time.Hour).Seconds())
	eightLoc := time.FixedZone("eight", eightZone)
	t, err = time.ParseInLocation("2006-01-02 15:04:05", "2018-09-10 00:00:00", eightLoc)
	if err != nil {
		panic(err)
	}
	fmt.Println(t)
	fmt.Println(t.In(time.UTC).Sub(t))

	createTime, err := time.Parse("2006-01-02 15:04:05", "2019-12-24 14:20:55")
	if err != nil {
		panic(err)
	}
	fmt.Println(createTime)
	fmt.Println(createTime.Add(-8 * time.Hour))

	expireTimeGmt, err := time.ParseInLocation("2006-01-02 15:04:05", "2019-12-24 15:25:24", time.UTC)
	if err != nil {
		panic(err)
	}
	fmt.Println(expireTimeGmt)
}
