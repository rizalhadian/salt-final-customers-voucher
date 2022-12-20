package infrastructure_transaction

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	infrastructure_transaction_http_response "salt-final-voucher/internal/infrastructure/transaction/http_response"
	infrastructure_transaction_interface "salt-final-voucher/internal/infrastructure/transaction/interface"
	"strconv"
	"strings"
	"time"
)

type InfrastructureTransaction struct {
	http_client   *http.Client
	base_endpoint string
}

func NewInfrastructureTransaction(http_client_value http.Client, base_endpoint_value string) infrastructure_transaction_interface.InterfaceInfrastructureTransaction {
	return &InfrastructureTransaction{
		// base_endpoint: "http://localhost:8080/api/customer/{customer_id}/transaction-count",
		http_client:   &http_client_value,
		base_endpoint: base_endpoint_value,
	}
}

func (ic InfrastructureTransaction) GetCustomersTransactionCount(ctx context.Context, customer_id int64) (customers_transaction_count *infrastructure_transaction_http_response.CustomersTransactionResponseSuccess, err error) {
	_, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	endpoint := strings.Replace(ic.base_endpoint, "{customer_id}", strconv.Itoa(int(customer_id)), -1)
	fmt.Println(endpoint)
	request, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, errors.New("500")
	}

	response, err := ic.http_client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode == 404 {
		return nil, errors.New("404")
	}

	customers_transaction_resp_success := infrastructure_transaction_http_response.CustomersTransactionResponseSuccess{}
	fmt.Println(response.Body)
	err = json.NewDecoder(response.Body).Decode(&customers_transaction_resp_success)
	if err != nil {
		return nil, err
	}

	return &customers_transaction_resp_success, nil
}
