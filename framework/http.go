package main

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/dghubble/sling"
)

func main() {
	req, err := sling.New().Get("http://www.baidu.com").Request()
	if err != nil {
		panic(err)
	}

	httpClient := &http.Client{}

	resp, err := httpClient.Do(req)
	if err != nil {
		panic(err)
	}
	headers := resp.Header

	for k, v := range headers {
		fmt.Println(fmt.Sprintf("%s - %s", k, v))
	}
	fmt.Printf("length-> %d \n", resp.ContentLength)

	resBody := resp.Body
	buf := new(bytes.Buffer)
	buf.ReadFrom(resBody)
	fmt.Println(buf.String())

}
