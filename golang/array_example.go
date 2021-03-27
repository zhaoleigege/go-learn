package main

import (
	"fmt"
	"log"
	"net/http"

	"./middleware"
)

// Interceptor1 第一个拦截器
func Interceptor1() middleware.Interceptor {
	return func(write http.ResponseWriter, request *http.Request, next http.HandlerFunc) {
		fmt.Println("1")
		next(write, request)
	}
}

// Interceptor2 第二个拦截器
func Interceptor2() middleware.Interceptor {
	return func(write http.ResponseWriter, request *http.Request, next http.HandlerFunc) {
		fmt.Println("2")
		next(write, request)
	}
}

type RefundResp struct {
	TicketGUID   string           `json:"ticket_guid"`
	Currency     string           `json:"currency"`
	RefundAmount string `json:"refund_amount"`
	RefundFee   string `json:"refund_fee"`
	Refundable   bool             `json:"refundable"`
	Reason       string           `json:"reason"`
}

// hanlder 实际执行的方法
func hanlder(w http.ResponseWriter, r *http.Request) {
	fmt.Println("访问")
	
	resp := &RefundResp{
		TicketGUID  : "61cdb9fd-524a-4905-4693-09de3fecab72",
Currency    : "CNY",
RefundAmount: "400",
RefundFee:   "86",
Refundable  : true,
Reason      :"",
	}

	bytes, err := json.Marshal(resp)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, bytes)
}

func main() {
	middlewareChain := middleware.HandlerFuncChain{
		Interceptor1(),
		Interceptor2(),
	}

	http.Handle("/", middlewareChain.Handler(hanlder))

	log.Panic(http.ListenAndServe(":11000", nil))
}
