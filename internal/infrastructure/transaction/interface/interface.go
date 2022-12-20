package infrastructure_transaction_interface

import (
	"context"
	infrastructure_transaction_http_response "salt-final-voucher/internal/infrastructure/transaction/http_response"
)

// type InterfaceInfrastructureVoucher interface {
// 	GetById(ctx context.Context, id int64) (customer *infrastructure_customer_http_response.Customer, http_response_code int, err error)
// }

type InterfaceInfrastructureTransaction interface {
	GetCustomersTransactionCount(ctx context.Context, customer_id int64) (customers_transaction_count *infrastructure_transaction_http_response.CustomersTransactionResponseSuccess, err error)
}
