package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestSimpleGet(t *testing.T) {
	res, err := http.NewRequest(http.MethodGet, "/get?name=test&age=21", nil)
	if err != nil {
		t.Error(err)
	}

	resRec := httptest.NewRecorder()
	router().ServeHTTP(resRec, res)

	t.Log(resRec.Body)
}
func TestSimplePost(t *testing.T) {
	person := Person{Name: "test", Age: 21}
	req, err := json.Marshal(person)
	if err != nil {
		log.Fatal("json格式化错误:", err)
	}

	res, err := http.NewRequest(http.MethodPost, "/post", bytes.NewBuffer(req))
	if err != nil {
		t.Error(err)
	}

	resRec := httptest.NewRecorder()
	router().ServeHTTP(resRec, res)

	t.Log(resRec.Body)
}

func TestSimpleForm(t *testing.T) {
	forms := url.Values{}
	forms.Add("name", "test")
	forms.Add("age", "21")

	res, err := http.NewRequest(http.MethodPost, "/form", strings.NewReader(forms.Encode()))
	if err != nil {
		t.Error(err)
	}
	res.Header.Add("Content-Type", "application/x-www-form-urlencoded") // 记住设置请求头

	resRec := httptest.NewRecorder()
	router().ServeHTTP(resRec, res)

	t.Log(resRec.Body)
}

func TestSimpleVarRouter(t *testing.T) {
	res, err := http.NewRequest(http.MethodGet, "/person/test?age=21", nil)
	if err != nil {
		t.Error(err)
	}

	resRec := httptest.NewRecorder()
	router().ServeHTTP(resRec, res)

	t.Log(resRec.Body)
}

func TestSubRouter(t *testing.T) {
	urls := [2]string{"/sub/first", "/sub/second"}

	for _, v := range urls {
		res, err := http.NewRequest(http.MethodGet, v, nil)
		if err != nil {
			t.Error(err)
		}

		resRec := httptest.NewRecorder()
		router().ServeHTTP(resRec, res)

		t.Log(resRec.Body)
	}
}

