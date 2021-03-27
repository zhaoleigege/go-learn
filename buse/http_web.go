package main

import (
	"fmt"
	"log"
	"net/http"
)

func hanlder(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `{"error": {"code": "", "message": ""}, "result": {}, "success": true}`)
}

func mapper(h func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("请求前")
		h(w, r)
		fmt.Println("请求后")
	})
}

func logMapper(h http.Handler) http.Handler {
	// 把匿名函数func(w http.ResponseWriter, r *http.Request) {...}转换为http.HandlerFunc类型
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("log请求前")
		h.ServeHTTP(w, r)
		fmt.Println("log请求后")
	})
}

func main() {
	http.Handle("/", logMapper(mapper(hanlder)))

	log.Panic(http.ListenAndServe(":8000", nil))
}
