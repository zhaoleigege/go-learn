// golang原生http访问示例
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Get get请求
func Get() (string, error) {
	url := "http://www.baidu.com"

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

// IdItem id结构体
type IdItem struct {
	Id int64 `json:"id"`
}

// OtherInfo 结构体
type OtherInfo struct {
	Type           int64     `json:"type"`
	TicketItemList []*IdItem `json:"ticket_item_list"`
}

// Post post请求
func Post() (string, error) {
	url := "http://t58.test.klook.io/v1/orderserv/otherinfo"
	reqStruct := &OtherInfo{
		Type: 1,
		TicketItemList: []*IdItem{
			{Id: 1795950},
		},
	}

	reqBod, err := json.Marshal(reqStruct)
	if err != nil {
		return "", err
	}

	contentType := "application/json"

	resp, err := http.Post(url, contentType, bytes.NewBuffer(reqBod))
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

// Do 可以自定义的http请求
func Do() (string, error) {
	method := "POST"
	url := "http://t58.test.klook.io/v1/orderserv/otherinfo"
	contentType := "application/json"

	reqStruct := &OtherInfo{
		Type: 1,
		TicketItemList: []*IdItem{
			{Id: 1795950},
		},
	}

	reqBod, err := json.Marshal(reqStruct)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(reqBod))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", contentType)
	req.Header.Set("X-Klook-Request-Id", "client-do")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func main() {
	content, err := Do()
	if err != nil {
		panic(err)
	}

	fmt.Println(content)
}
