package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "pp!le5Pl"

	for i := 0; i < len(str); i++ {
		println(strings.ToLower(string(str[i])))
	}
	fmt.Println()
	fmt.Println("[{\"order_id\":200114225,\"ticket_id\":200163943,\"sku_id\":800008000501,\"price_id\":169593,\"price_cost_db\":100,\"price_cost_db_currency\":\"EUR\",\"price_value_db\":3000000,\"price_value_db_currency\":\"HKD\",\"price_market_db\":9000,\"price_market_db_currency\":\"HKD\"}]")

	fmt.Println("test")

	fmt.Println(2 & (1 << 2))
}
