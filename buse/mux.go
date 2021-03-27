package main

import (
	"fmt"
	"net/http"
	"net/http/pprof"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/debug/pprof/", pprof.Index)
	r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	r.HandleFunc("/debug/pprof/profile", pprof.Profile)
	r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	r.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
	r.Handle("/debug/pprof/heap", pprof.Handler("heap"))
	r.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
	r.Handle("/debug/pprof/block", pprof.Handler("block"))

	r.HandleFunc("/test", AccessHttp)
	http.Handle("/", r)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func AccessHttp(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("接受请求: %s\n", r.URL)
	w.Write([]byte("成功响应"))
}
