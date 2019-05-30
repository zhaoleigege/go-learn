package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"strconv"
)

var (
	port int
)

type ResEntity struct {
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func init() {
	flag.IntVar(&port, "port", 8000, "端口号")
}

func router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/get", simpleGet).Methods(http.MethodGet)
	r.HandleFunc("/post", simplePost).Methods(http.MethodPost)
	r.HandleFunc("/form", simpleForm).Methods(http.MethodPost)
	r.HandleFunc("/person/{name}", simpleVarRouter).Methods(http.MethodGet)

	dir, err := os.UserHomeDir()
	if err != nil {
		dir = fmt.Sprintf("%s/mux", "/var")
	} else {
		dir = fmt.Sprintf("%s/Desktop/mux", dir)
	}
	log.Printf("文件存储目录为：%s\n", dir)

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))

	rSub := r.PathPrefix("/sub").Subrouter()
	{
		rSub.HandleFunc("/first", subFirst).Methods(http.MethodGet)
		rSub.HandleFunc("/second", subSecond).Methods(http.MethodGet)
	}
	rSub.Use(logMiddleWare) // 注册中间件，这里只给子路由设置了中间件
	return r
}

func main() {
	flag.Parse()
	address := fmt.Sprintf(":%d", port)
	log.Printf("服务器启动%s\n", address)

	if err := http.ListenAndServe(address, router()); err != nil {
		log.Fatal(err, "服务器启动失败")
	}

}

// subFirst是第一个子路由
func subFirst(w http.ResponseWriter, r *http.Request) {
	log.Printf("访问第一个子路由\n")
	responseMsg(w, 0, "", struct{ Message string `json:"message"` }{Message: "第一个子路由"})
}

// subSecond是第二个子路由
func subSecond(w http.ResponseWriter, r *http.Request) {
	log.Printf("访问第二个子路由\n")
	responseMsg(w, 0, "", struct{ Message string `json:"message"` }{Message: "第二个子路由"})
}

// simpleVarRouter展示了url动态参数的解析
func simpleVarRouter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // 只能够取到url中的动态值

	name, ok := vars["name"]
	if !ok {
		log.Printf("参数name不存在\n")
		responseMsg(w, -1, "缺少参数name", nil)
		return
	}

	if err := r.ParseForm(); err != nil {
		responseMsg(w, -1, err.Error(), nil)
		return
	}

	var age int
	param, ok := r.Form["age"]

	if ok {
		if p, err := strconv.Atoi(param[0]); err != nil {
			log.Printf("age: %s输入错误\n", param)
			responseMsg(w, -1, "参数age错误", nil)
			return
		} else {
			age = p
		}
	}

	person := &Person{
		Name: name,
		Age:  age,
	}

	log.Printf("获取数据%+v\n", person)
	responseMsg(w, 0, "", &person)
}

// simpleForm演示了如何上传form数据
func simpleForm(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		responseMsg(w, -1, err.Error(), nil)
		return
	}

	for k, v := range r.Form {
		log.Printf("key: %s, value: %q\n", k, v)
	}

	responseMsg(w, 0, "", struct{ Message string `json:"message"` }{Message: "信息接收成功"})
}

// simplePost演示了如何上传json数据
func simplePost(w http.ResponseWriter, r *http.Request) {
	var person Person

	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		responseMsg(w, -1, err.Error(), nil)
		return
	}

	type Message struct {
	}

	log.Printf("获取数据%+v\n", person)

	responseMsg(w, 0, "", struct{ Message string `json:"message"` }{Message: "信息接收成功"})
}

// simpleGet演示了最普通的get请求，并解析请求参数
func simpleGet(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()

	param, ok := values["age"]
	if !ok {
		log.Printf("参数age不存在\n")
		responseMsg(w, -1, "缺少参数age", nil)
		return
	}

	age, err := strconv.Atoi(param[0])
	if err != nil {
		str := fmt.Sprintf("参数age: %d 转换错误", age)
		log.Println(str)
		responseMsg(w, -1, str, nil)
		return
	}

	var name string
	param, ok = values["name"]
	if !ok {
		log.Printf("参数name不存在\n")
	} else {
		name = param[0]
	}

	p := &Person{
		Name: name,
		Age:  age,
	}

	log.Printf("获取数据%+v\n", p)
	responseMsg(w, 0, "", &p)
}

func responseMsg(w http.ResponseWriter, code int, msg string, data interface{}) {
	resEnt := &ResEntity{}

	resEnt.Code = code
	resEnt.Message = msg

	if data != nil {
		resEnt.Data = data
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(resEnt)

	if err != nil {
		fmt.Println("数据响应错误：" + err.Error())
	}
}

// logMiddleWare实现了一个中间，对所有的请求加上访问日志
func logMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			log.Printf("访问了 %s\n", r.URL)
			next.ServeHTTP(w, r)
		})
}
