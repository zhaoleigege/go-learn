package main

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.TODO())
	timer := time.AfterFunc(5*time.Second, func() {
		cancel()
	})
	req, err := http.NewRequest("GET", "http://httpbin.org/range/2048?duration=8&chunk_size=256", nil)
	if err != nil {
		log.Fatal(err)
	}
	req = req.WithContext(ctx)
	log.Println("Sending request...")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		if ctx.Err() != nil {
			fmt.Println("超时错误")
			fmt.Println(ctx.Err())
			return
		}

		fmt.Println("其他错误")
		log.Fatal(err)
	}
	defer resp.Body.Close()

	log.Println("Reading body...")
	for {
		timer.Reset(2 * time.Second)
		// Try instead: timer.Reset(50 * time.Millisecond)
		_, err = io.CopyN(ioutil.Discard, resp.Body, 256)
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
	}
}
