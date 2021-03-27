package main

import (
	"fmt"
	"log"
	"net/http"
)

// HTTPHandler 是http.Handler类型
type HTTPHandler http.HandlerFunc

// Middleware 是中间件
type Middleware func(http.ResponseWriter, *http.Request)

func (handler HTTPHandler) add(middleware Middleware) HTTPHandler {
	return func(writer http.ResponseWriter, request *http.Request) {
		middleware(writer, request)
		handler(writer, request)
	}
}

// Logging 打印日志
func Logging(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("打印日志")
}

// Filter 过滤不合法的请求
func Filter(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("打印日志")
}

// WriteHandler 执行的地方
func WriteHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("执行")
}

func main() {
	handler := HTTPHandler(WriteHandler)

	http.HandleFunc("/", handler.add(Logging))

	log.Panic(http.ListenAndServe(":8000", nil))
}
