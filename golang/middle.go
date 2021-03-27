package main

import (
	"fmt"
	"log"
	"net/http"
)

// middleware 为接口类型
type middleware interface {
	MiddleWare(handler http.Handler) http.Handler
}

// MiddlewareFunc 为函数类型
type MiddlewareFunc func(handler http.Handler) http.Handler

// Middleware 是函数类型MiddlewareFunc实现了middleware接口定义的方法，所以MiddlewareFunc就是middleware接口类型
func (mw MiddlewareFunc) Middleware(handler http.Handler) http.Handler {
	return mw(handler)
}

// HandlerFunc 类型其实是http.Handler，这里则是为了添加新的方法
type HandlerFunc http.HandlerFunc

func (hf HandlerFunc) Middlewares(mws ...MiddlewareFunc) HandlerFunc {
	return HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, mw := range mws {
			mw.Middleware(http.Handler.(hf))
		}
	})
}

// Login 登陆
func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("登陆")
}

// Logging 中间件
func Logging(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("loggind")
		handler.ServeHTTP(w, r)
	})
}

func main() {
	http.HandleFunc("/", HandlerFunc(Login).Middlewares(Logging))

	log.Panic(http.ListenAndServe(":8000", nil))
}
