package main

import (
	"fmt"
	"io/ioutil"
)

type Page struct {
	Title string
	Body  []byte
}

func (page *Page) save() error {
	filename := page.Title + ".txt"
	return ioutil.WriteFile(filename, page.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	return &Page{Title: filename, Body: body}, nil
}

func main() {
	p1 := &Page{Title: "test", Body: []byte("测试页面")}
	p1.save()

	p2, err := loadPage("test")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(p2.Body))
}
