package main

import (
	"fmt"
	"net/url"
	"strconv"
)

func main() {
	var param = url.Values{}
	param.Add("account_id", strconv.Itoa(123))
	param.Add("account_name", "Rail China")

	reqUrl := "http://t45.test.klook.io" + "/v1/refundserv/refund_ticket/do_refund_pure?" + param.Encode()

	fmt.Println(reqUrl)
}
