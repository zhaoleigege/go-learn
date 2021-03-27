package _go

import (
	"fmt"
	"github.com/elazarl/goproxy"
	"log"
	"net/http"
	"regexp"
)

func main() {
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = false
	proxy.OnRequest(goproxy.ReqHostMatches(regexp.MustCompile("^.*$"))).
		HandleConnect(goproxy.AlwaysMitm)

	proxy.OnRequest(goproxy.DstHostIs("www.reddit.com")).DoFunc(
		func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
			return r, goproxy.NewResponse(r,
				goproxy.ContentTypeText, http.StatusForbidden,
				"Don't waste your time!")
		})

	proxy.OnRequest().
		DoFunc(func(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
			fmt.Println(req.Header)

			return req, nil
		})

	log.Fatal(http.ListenAndServe(":8080", proxy))
}
