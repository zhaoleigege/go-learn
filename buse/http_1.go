package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetUser(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		responseMsg(response, 300, err.Error(), nil)
		return
	}

	user := User{
		Id: id,
	}

	err = user.selectUser()

	if err != nil {
		responseMsg(response, 300, err.Error(), nil)
	} else {
		responseMsg(response, 200, "", user)
	}
}

func GetAllUsers(response http.ResponseWriter, request *http.Request) {
	users, err := selectUsers()

	if err != nil {
		responseMsg(response, 300, err.Error(), nil)
	} else {
		responseMsg(response, 200, "", users)
	}
}

func AddUser(response http.ResponseWriter, request *http.Request) {
	var user User

	err := json.NewDecoder(request.Body).Decode(&user)

	if err != nil {
		responseMsg(response, 300, err.Error(), nil)
		return
	}

	err = insertUser(user)

	if err != nil {
		responseMsg(response, 300, err.Error(), nil)
	} else {
		responseMsg(response, 200, "用户添加成功", nil)
	}
}

func UpdateUser(response http.ResponseWriter, request *http.Request) {
	var user User

	err := json.NewDecoder(request.Body).Decode(&user)

	if err != nil {
		responseMsg(response, 300, err.Error(), nil)
		return
	}

	err = updateUser(user)

	if err != nil {
		responseMsg(response, 300, err.Error(), nil)
	} else {
		responseMsg(response, 200, "用户更新成功", nil)
	}
}

func DeleteUser(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		responseMsg(response, 300, err.Error(), nil)
		return
	}

	err = deleteUser(id)

	if err != nil {
		responseMsg(response, 300, err.Error(), nil)
	} else {
		responseMsg(response, 200, "用户删除成功", nil)
	}
}

func responseMsg(response http.ResponseWriter, code int, msg string, data interface{}) {
	resEnt := ResEntity{}

	resEnt.Code = code
	resEnt.Message = msg

	if data != nil {
		resEnt.Data = data
	}

	response.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(response).Encode(resEnt)

	if err != nil {
		fmt.Println("数据响应错误：" + err.Error())
	}
}
