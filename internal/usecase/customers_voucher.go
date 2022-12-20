package usecase

import (
	"context"
	"errors"
	"math/rand"
	"salt-final-voucher/domain/entity"
	"salt-final-voucher/domain/interface_repo"
	"salt-final-voucher/domain/interface_usecase"
	interface_infrastructure_transaction "salt-final-voucher/internal/infrastructure/transaction/interface"
	"strconv"
	"time"
)

type UsecaseCustomerVoucher struct {
	infraTransaction     interface_infrastructure_transaction.InterfaceInfrastructureTransaction
	repoCustomersVoucher interface_repo.InterfaceRepoCustomersVoucher
}

func NewUsecaseCustomerVoucher(infraTransactionVal interface_infrastructure_transaction.InterfaceInfrastructureTransaction, repoCustomersVoucherVal interface_repo.InterfaceRepoCustomersVoucher) interface_usecase.InterfaceRepoCustomersVoucher {
	return &UsecaseCustomerVoucher{
		infraTransaction:     infraTransactionVal,
		repoCustomersVoucher: repoCustomersVoucherVal,
	}
}

func (ucv *UsecaseCustomerVoucher) Redeem(ctx context.Context, entity_transaction *entity.DTOTransaction) ([]*entity.CustomersVoucher, error) {
	return nil, nil
}

func (ucv *UsecaseCustomerVoucher) Generate(ctx context.Context, customer_id int64) ([]*entity.CustomersVoucher, error) {
	// if customer_id == 0 {
	// 	return nil, errors.New("Cannot Use customer_id 0")
	// }

	get_customers_transaction_count, get_customers_transaction_count_err := ucv.infraTransaction.GetCustomersTransactionCount(ctx, customer_id)
	if get_customers_transaction_count_err != nil {
		if get_customers_transaction_count_err.Error() == "404" {
			return nil, errors.New("404")
		} else {
			return nil, errors.New("500")
		}
	}

	var entity_customers_vouchers []*entity.CustomersVoucher
	current_time := time.Now()

	for _, voucher := range entity.Vouchers {
		if get_customers_transaction_count.Data.Total_transaction_spend > voucher.Condition_to_obtain_trx_min_total_spend {

			_, err_count_voucher := ucv.repoCustomersVoucher.GetByVoucherId(ctx, int32(voucher.Id))
			if err_count_voucher != nil {
				if err_count_voucher.Error() == "404" {
					// Genereate Voucher Here
					random_int_for_code := rangeIn(voucher.Code_random_total_digit)
					random_string_for_code := strconv.Itoa(random_int_for_code)

					dto_customers_voucher := &entity.DTOCustomersVoucher{
						Customer_id:  customer_id,
						Voucher_id:   voucher.Id,
						Voucher_name: voucher.Name,
						Code:         voucher.Prefix_code + random_string_for_code,
						Expired_at:   current_time.Add(voucher.Expired_add_day_from_generated),
						Status:       120,
					}

					entity_customers_voucher, entity_customers_voucher_err := entity.NewCustomersVoucher(dto_customers_voucher)
					if entity_customers_voucher_err != nil {
						return nil, entity_customers_voucher_err
					}
					store_err := ucv.repoCustomersVoucher.Store(ctx, entity_customers_voucher)
					if store_err != nil {
						return nil, errors.New("500")
					}

					entity_customers_vouchers = append(entity_customers_vouchers, entity_customers_voucher)
				}
			}

		}
	}
	return entity_customers_vouchers, nil
}

func rangeIn(digit int) int {
	var low int
	var low_string string
	var hi int
	var hi_string string

	for i := 0; i < digit; i++ {

		if i == 0 {
			low_string = "1"
			hi_string = "9"
		} else {
			low_string = low_string + "0"
			hi_string = hi_string + "9"
		}
	}

	low, _ = strconv.Atoi(low_string)
	hi, _ = strconv.Atoi(hi_string)

	return low + rand.Intn(hi-low)
}
