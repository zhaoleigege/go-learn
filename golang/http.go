package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"./klook"

	"github.com/gorilla/mux"
)

func getUser(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, _ := strconv.Atoi(vars["id"])

	klook.GetUser(id)

	var user User
	user.Id = id
	user.Name = "test"
	user.Password = "test"
	user.Status = 0

	resStr, _ := json.Marshal(user)

	response.Header().Set("Content-Type", "application/json")
	response.Write(resStr)
}

func main() {
	r := mux.NewRouter()
	// 根据id号获取用户信息
	r.HandleFunc("/user/{id}", getUser).Methods("GET")

	// 获取所有用户信息
	r.HandleFunc("/users", func(http.ResponseWriter, *http.Request) {

	}).Methods("GET")

	// 添加一个新的用户
	r.HandleFunc("/user", func(http.ResponseWriter, *http.Request) {

	}).Methods("POST")

	// 修改一个用户的信息
	r.HandleFunc("/user", func(http.ResponseWriter, *http.Request) {

	}).Methods("PUT")

	// 删除一个用户
	r.HandleFunc("/user/{id}", func(http.ResponseWriter, *http.Request) {

	}).Methods("DELETE")

	http.ListenAndServe(":80", r)
}
