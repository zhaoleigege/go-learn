package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func router() *mux.Router {
	r := mux.NewRouter()
	// 根据id号获取用户信息
	r.HandleFunc("/user/{id}", GetUser).Methods("GET")

	// 获取所有用户信息
	r.HandleFunc("/users", GetAllUsers).Methods("GET")

	// 添加一个新的用户
	r.HandleFunc("/user", AddUser).Methods("POST")

	// 修改一个用户的信息
	r.HandleFunc("/user", UpdateUser).Methods("PUT")

	// 删除一个用户
	r.HandleFunc("/user/{id}", DeleteUser).Methods("DELETE")

	return r
}

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)

	// 数据库连接初始化
	if err := DbInit(); err != nil {
		log.Fatal(err)
	}

	//defer DbDestroy()

	// 设置http服务器的相关信息
	server := http.Server{Addr: ":80", Handler: router()}

	go func() {
		_ = server.ListenAndServe() // 启动http服务器
	}()

	for {
		s := <-c
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT: // 程序退出时执行的指令
			DbDestroy()
			_ = server.Shutdown(nil)
			fmt.Println("程序退出")
			return
		case syscall.SIGHUP:
		// TODO reload
		default:
			return
		}
	}
}
