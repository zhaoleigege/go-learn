package middleware

// 参考 https://doordash.engineering/2019/07/22/writing-delightful-http-middlewares-in-go/
//     https://gist.github.com/x0a1b/f7c64c92cfc6eae7ba1c2931bbabb475

import "net/http"

// Interceptor 中间件拦截器
type Interceptor func(http.ResponseWriter, *http.Request, http.HandlerFunc)

// HandlerFunc 执行具体操作的函数
type HandlerFunc http.HandlerFunc

// Register 是拦截器注册函数
func (handler HandlerFunc) Register(interceptor Interceptor) HandlerFunc {
	return func(write http.ResponseWriter, request *http.Request) {
		interceptor(write, request, http.HandlerFunc(handler))
	}
}

// HandlerFuncChain 是拦截器链
type HandlerFuncChain []Interceptor

// Handler 方法实现了拦截器链的注册
func (chain HandlerFuncChain) Handler(handler http.HandlerFunc) http.HandlerFunc {
	curHandler := HandlerFunc(handler)

	for i := range chain {
		curHandler = curHandler.Register(chain[len(chain)-i-1])
	}

	return http.HandlerFunc(curHandler)
}
