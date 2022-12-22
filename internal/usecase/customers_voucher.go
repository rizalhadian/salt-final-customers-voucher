package usecase

import (
	"context"
	"errors"
	"fmt"
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

type TotalAmountPerCategory struct {
	items_type_id int64
	total_qty     int32
	total_amount  float64
}

func NewUsecaseCustomerVoucher(infraTransactionVal interface_infrastructure_transaction.InterfaceInfrastructureTransaction, repoCustomersVoucherVal interface_repo.InterfaceRepoCustomersVoucher) interface_usecase.InterfaceRepoCustomersVoucher {
	return &UsecaseCustomerVoucher{
		infraTransaction:     infraTransactionVal,
		repoCustomersVoucher: repoCustomersVoucherVal,
	}
}

func (ucv *UsecaseCustomerVoucher) Redeem(ctx context.Context, dto_transaction *entity.DTOTransaction, dto_transactions_items []*entity.DTOTransactionsItem, dto_transactions_vouchers_redeem []*entity.DTOCustomersVoucher) ([]*entity.CustomersVoucher, error) {
	var entity_customers_vouchers_redeemed []*entity.CustomersVoucher

	fmt.Println("Check And Get Voucher By Code")

	for _, dto_voucher := range dto_transactions_vouchers_redeem {
		entity_voucher, entity_voucher_err := ucv.repoCustomersVoucher.GetRedeemableVoucherByCode(ctx, dto_voucher.Code)
		if entity_voucher_err != nil {
			if entity_voucher_err.Error() == "404" {
				return nil, errors.New("404")
			}
			return nil, errors.New("500")
		}
		voucher_found := false
		for _, voucher := range entity.Vouchers {
			if voucher.Id == entity_voucher.GetVoucherId() {
				voucher_found = true
				entity_voucher.SetVoucher(voucher)
				entity_customers_vouchers_redeemed = append(entity_customers_vouchers_redeemed, entity_voucher)
			}
		}
		if voucher_found == false {
			return nil, errors.New("404")
		}
	}
	fmt.Println(entity_customers_vouchers_redeemed[0].GetId())

	fmt.Println("Count And Grouping Total Amount Per Category")
	var total_amounts_per_category []TotalAmountPerCategory
	for dto_transactions_item_index, dto_transactions_item := range dto_transactions_items {
		if dto_transactions_item_index == 0 {
			total_amounts_per_category = append(total_amounts_per_category, TotalAmountPerCategory{
				items_type_id: dto_transactions_item.Items_type_id,
				total_qty:     dto_transactions_item.Qty,
				total_amount:  dto_transactions_item.Total_price,
			})
		} else {
			found_item_type_id := false
			for total_amount_per_category_item_index, total_amount_per_category_item := range total_amounts_per_category {
				if total_amount_per_category_item.items_type_id == dto_transactions_item.Items_type_id {
					total_amount_per_category_item.total_amount += dto_transactions_item.Total_price
					total_amount_per_category_item.total_qty += dto_transactions_item.Qty
					// Ga bisa di set langsung, musti pake index
					total_amounts_per_category[total_amount_per_category_item_index].total_amount = total_amount_per_category_item.total_amount
					total_amounts_per_category[total_amount_per_category_item_index].total_qty = total_amount_per_category_item.total_qty
					found_item_type_id = true
				}
			}
			if found_item_type_id == false {
				total_amounts_per_category = append(total_amounts_per_category, TotalAmountPerCategory{
					items_type_id: dto_transactions_item.Items_type_id,
					total_qty:     dto_transactions_item.Qty,
					total_amount:  dto_transactions_item.Total_price,
				})
			}
		}
	}
	fmt.Println(total_amounts_per_category)

	fmt.Println("Count Voucher : ")
	fmt.Println(len(entity_customers_vouchers_redeemed))

	fmt.Println("Cek dan perhitungan voucher")

	for entity_customers_voucher_redeemed_index, entity_customers_voucher_redeemed := range entity_customers_vouchers_redeemed {
		fmt.Println("Cek dan perhitungan voucher Index : ")
		fmt.Println(entity_customers_voucher_redeemed_index)

		is_voucher_req_match := true
		total_qty_voucher_req_items_type_ids := 0
		total_amount_voucher_req_items_type_ids := 0.0

		fmt.Println("Cek dan perhitungan voucher Condition_to_redeem_trx_has_items_type_ids")
		for _, req_items_type_id := range entity_customers_voucher_redeemed.GetVoucher().Condition_to_redeem_trx_has_items_type_ids {
			fmt.Println("Cek dan perhitungan voucher Condition_to_redeem_trx_has_items_type_id")
			for _, total_amount_per_category := range total_amounts_per_category {
				fmt.Println("Cek dan perhitungan voucher Condition_to_redeem_trx_has_items_type_id, loop total amount")
				if total_amount_per_category.items_type_id == req_items_type_id {
					fmt.Println("Cek dan perhitungan voucher Condition_to_redeem_trx_has_items_type_id, total_amount_per_category.items_type_id == req_items_type_id")
					fmt.Println(total_amount_per_category.items_type_id)
					fmt.Println(req_items_type_id)
					total_qty_voucher_req_items_type_ids += int(total_amount_per_category.total_qty)
					total_amount_voucher_req_items_type_ids += total_amount_per_category.total_amount
				}
			}
		}
		fmt.Println("total_qty_voucher_req_items_type_ids")
		fmt.Println(total_qty_voucher_req_items_type_ids)
		fmt.Println("total_amount_voucher_req_items_type_ids")
		fmt.Println(total_amount_voucher_req_items_type_ids)

		if total_qty_voucher_req_items_type_ids == 0 {
			is_voucher_req_match = false
			fmt.Println("cond 1 not passed")
		}

		if entity_customers_voucher_redeemed.GetVoucher().Condition_to_redeem_trx_min_qty_with_same_items_type_ids > int32(total_qty_voucher_req_items_type_ids) {
			is_voucher_req_match = false
			fmt.Println("cond 2 not passed")
		}

		if entity_customers_voucher_redeemed.GetVoucher().Condition_to_redeem_trx_total_min_amount > dto_transaction.Total_amount {
			is_voucher_req_match = false
			fmt.Println("cond 3 not passed")

		}

		if entity_customers_voucher_redeemed.GetVoucher().Condition_to_redeem_trx_min_amount_with_same_items_type_ids > float64(total_amount_voucher_req_items_type_ids) {
			is_voucher_req_match = false

			fmt.Println("cond 4 not passed")

		}

		if is_voucher_req_match == false {
			return nil, errors.New("Transaction doesn't fullfil voucher's requirements")
		}

		voucher_amount_deduction := 0.00
		if entity_customers_voucher_redeemed.GetVoucher().Voucher_deduction_is_on_total_trx == true {
			fmt.Println("Deduction on total trx")
			if entity_customers_voucher_redeemed.GetVoucher().Voucher_deduction_is_constant_amount {

				if dto_transaction.Total_amount > entity_customers_voucher_redeemed.GetVoucher().Voucher_deduction_amount {
					voucher_amount_deduction = dto_transaction.Total_amount
				} else {
					voucher_amount_deduction = entity_customers_voucher_redeemed.GetVoucher().Voucher_deduction_amount
				}

			}

			if entity_customers_voucher_redeemed.GetVoucher().Voucher_deduction_is_on_percentage {
				voucher_amount_deduction = dto_transaction.Total_amount * (float64(entity_customers_voucher_redeemed.GetVoucher().Voucher_deduction_percentage / 100))
			}
		}

		if entity_customers_voucher_redeemed.GetVoucher().Voucher_deduction_is_on_total_specific_items_type_ids == true {
			fmt.Println("Deduction on total specific ids")
			if entity_customers_voucher_redeemed.GetVoucher().Voucher_deduction_is_constant_amount {
				fmt.Println("Deduction on total specific ids | constanst amount")
				if total_amount_voucher_req_items_type_ids > entity_customers_voucher_redeemed.GetVoucher().Voucher_deduction_amount {
					voucher_amount_deduction = total_amount_voucher_req_items_type_ids
				} else {
					voucher_amount_deduction = entity_customers_voucher_redeemed.GetVoucher().Voucher_deduction_amount
				}
			}

			if entity_customers_voucher_redeemed.GetVoucher().Voucher_deduction_is_on_percentage {
				fmt.Println("Deduction on total specific ids | percentage amount")

				voucher_amount_deduction = float64(total_amount_voucher_req_items_type_ids) * (float64(entity_customers_voucher_redeemed.GetVoucher().Voucher_deduction_percentage) / 100.00)
				fmt.Println(voucher_amount_deduction)
			}
		}

		entity_customers_vouchers_redeemed[entity_customers_voucher_redeemed_index].SetTotalAmount(float64(total_amount_voucher_req_items_type_ids))
		entity_customers_vouchers_redeemed[entity_customers_voucher_redeemed_index].SetTotalDiscountAmount(voucher_amount_deduction)
		entity_customers_vouchers_redeemed[entity_customers_voucher_redeemed_index].SetFinalTotalAmount(entity_customers_voucher_redeemed.GetTotalAmount() - entity_customers_voucher_redeemed.GetTotalDiscountAmount())
		entity_customers_vouchers_redeemed[entity_customers_voucher_redeemed_index].SetStatus(121)
		entity_customers_vouchers_redeemed[entity_customers_voucher_redeemed_index].SetTransactionId(dto_transaction.Id)

		update_err := ucv.repoCustomersVoucher.Update(ctx, entity_customers_vouchers_redeemed[entity_customers_voucher_redeemed_index])
		if update_err != nil {
			return nil, errors.New("500")
		}
	}

	return entity_customers_vouchers_redeemed, nil
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
