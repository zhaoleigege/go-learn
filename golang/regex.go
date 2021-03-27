package main

import (
	"fmt"
	"regexp"
)

func main() {
	text := `[6ba687f] 价格比对并插入数据库失败, ticketId: 37084, error:
	[6ba687f] 价格比对并插入数据库失败, ticketId: 36929, error:
[ service/orderserv/projectkdb/fix/unit_fix.go:559 priceCompareAndUpdate:价格比较不相等]
[ service/orderserv/projectkdb/fix/unit_fix.go:649 PriceFix:]
[ service/orderserv/apis/order/datafix/price_fix.go:156 handlerFixPriceInfo:]
[ service/orderserv/apis/order/datafix/price_fix.go:121 fixPriceData:]
[ service/orderserv/apis/order/datafix/price_fix.go:63 HandlePriceFixRequest.func1:]
[ klook.libs/utils/func_util.go:23 HandlePanic.func1:]
[6ba687f] 价格比对并插入数据库失败, ticketId: 36983, error:
	`

	reg := regexp.MustCompile(`(价格比对并插入数据库失败, ticketId: \d+, error:)`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
}
