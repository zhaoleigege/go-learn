package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

// 数据库连接初始化
func init() {
	if err := DbInit(); err != nil {
		log.Fatal(err)
	}
}

func TestDeleteUser(t *testing.T) {
	res, err := http.NewRequest(http.MethodDelete, "/user/11", nil)
	if err != nil {
		t.Error(err)
	}

	resRec := httptest.NewRecorder()
	router().ServeHTTP(resRec, res)

	fmt.Println(resRec.Body)
}

func TestAddUser(t *testing.T) {
	user := User{Id: 11, Name: "test11", Password: "test11", Status: 0}

	response, _ := json.Marshal(user)

	res, err := http.NewRequest(http.MethodPost, "/user", bytes.NewBuffer(response))
	if err != nil {
		t.Error(err)
	}

	resRec := httptest.NewRecorder()
	router().ServeHTTP(resRec, res)

	fmt.Println(resRec.Body)
}

func TestGetUser(t *testing.T) {
	res, err := http.NewRequest(http.MethodGet, "/user/11", nil)
	if err != nil {
		t.Error(err)
	}

	resRec := httptest.NewRecorder()
	router().ServeHTTP(resRec, res)

	fmt.Println(resRec.Body)
}

func TestUpdateUser(t *testing.T) {
	user := User{Id: 11, Name: "test11", Password: "test", Status: 0}

	response, _ := json.Marshal(user)

	res, err := http.NewRequest(http.MethodPut, "/user", bytes.NewBuffer(response))
	if err != nil {
		t.Error(err)
	}

	resRec := httptest.NewRecorder()
	router().ServeHTTP(resRec, res)

	fmt.Println(resRec.Body)
}

func TestGetAllUsers(t *testing.T) {
	res, err := http.NewRequest(http.MethodGet, "/users", nil)
	if err != nil {
		t.Error(err)
	}

	resRec := httptest.NewRecorder()
	router().ServeHTTP(resRec, res)

	fmt.Println(resRec.Body)
}
