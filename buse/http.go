package main

import (
	"fmt"
	"net/http"

	// pprof性能监控
	_ "net/http/pprof"
)

func main() {
	http.HandleFunc("/test", AccessHttp)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func AccessHttp(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("接受请求: %s\n", r.URL)
	w.Write([]byte("成功响应"))
}
