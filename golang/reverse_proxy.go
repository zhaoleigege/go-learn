package main

// 参考
// https://imququ.com/post/web-proxy.html
// https://blog.csdn.net/mengxinghuiku/article/details/65448600
// https://segmentfault.com/a/1190000003735562
// http://legendtkl.com/2016/09/06/go-pool/
// https://ninokop.github.io/2018/03/20/%E8%AE%B0fastHTTP%E5%8D%8F%E7%A8%8B%E6%B1%A0%E7%9A%84%E5%AE%9E%E7%8E%B0/
// https://github.com/raysonxin/gokit-article-demo/blob/master/arithmetic_consul_demo/gateway/main.go

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	proxy := &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			url, err := url.Parse("http://landcareweb.com/questions/24753/dui-shi-yong-gorilla-mux-urlcan-shu-de-han-shu-jin-xing-dan-yuan-ce-shi")
			if err != nil {
				fmt.Println(err)
			} else {
				req.URL = url
			}
		}}

	http.ListenAndServe(":8001", proxy)
}
