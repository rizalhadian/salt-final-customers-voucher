package infrastructure_transaction_http_response

import "time"

type CustomersTransactionResponseSuccess struct {
	Success bool                      `json:"success"`
	Message string                    `json:"message"`
	Data    CustomersTransactionCount `json:"data"`
}

type CustomersTransactionCount struct {
	Id                         int64     `json:"id"`
	Customer_id                int64     `json:"customer_id"`
	Total_transaction_spend    float64   `json:"total_transaction_spend"`
	Transaction_count          int32     `json:"transaction_count"`
	First_transaction_datetime time.Time `json:"first_transaction_datetime"`
	Last_transaction_datetime  time.Time `json:"last_transaction_datetime"`
}
