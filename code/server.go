package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// RefundResp 返回的结构体
type RefundResp struct {
	TicketGUID   string `json:"ticket_guid"`
	Currency     string `json:"currency"`
	RefundAmount string `json:"refund_amount"`
	RefundFee    string `json:"refund_fee"`
	Refundable   bool   `json:"refundable"`
	Reason       string `json:"reason"`
	ApiName      string `json:"api_name"`
}

type Resp struct {
	Result  json.RawMessage `json:"result,omitempty"`
	Success bool            `json:"success"`
}

func main() {
	http.HandleFunc("/v1/apibusserv/order/query", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("访问")

		// resp := &RefundResp{
		// 	TicketGUID:   "61cdb9fd-524a-4905-4693-09de3fecab72",
		// 	Currency:     "TWD",
		// 	RefundAmount: "6214",
		// 	RefundFee:    "0",
		// 	Refundable:   true,
		// 	Reason:       "",
		// 	ApiName:      "car-rental",
		// }
		// bytes, err := json.Marshal()
		// if err != nil {
		// 	panic(err)
		// }

		result := &Resp{
			Success: true,
			Result:  []byte(`{"error":{"code":"","message":""},"result":{"ticket_id":212739300,"apiconn_name":"JTR","status":"","cancelable":false,"partial_cancelable":false,"reason":"API JTR 不支持退款","refund_available":false,"refund_amount":0,"refund_fee":0,"currency":""},"success":true}`),
		}

		bytesResult, err := json.Marshal(result)
		if err != nil {
			panic(err)
		}

		time.Sleep(5 * time.Second)
		// w.WriteHeader(http.StatusInternalServerError)
		w.Write(bytesResult)
	})

	http.ListenAndServe(":10094", nil)
}
