package entity

import "time"

const (
	Voucher_Status_Id_Service_And_Sparepart = 1
	Voucher_Status_Id_Console               = 2
	Voucher_Status_Id_Accesories            = 3
	Voucher_Status_Id_Game                  = 4
)

type Voucher struct {
	Id                                                            int64
	Name                                                          string
	Prefix_code                                                   string
	Code_random_total_digit                                       int
	Condition_to_obtain_trx_min_total_spend                       float64
	Condition_to_obtain_max_count_same_voucher                    int32
	Condition_to_redeem_trx_has_items_type_ids                    []int64
	Condition_to_redeem_trx_min_qty_with_same_items_type_ids      int32
	Condition_to_redeem_trx_min_amount_with_same_items_type_ids   float64
	Voucher_deduction_is_on_total_trx                             bool
	Voucher_deduction_is_on_total_specific_items_type_ids         bool
	Voucher_deduction_is_on_one_item_with_specific_items_type_ids bool
	Voucher_deduction_is_on_percentage                            bool
	Voucher_deduction_is_constant_amount                          bool
	Voucher_deduction_percentage                                  int16
	Voucher_deduction_amount                                      float64
	Expired_add_day_from_generated                                time.Duration
}

var Vouchers = []Voucher{
	Voucher{
		Id:                                      202212001,
		Name:                                    "BASIC",
		Prefix_code:                             "BASIC-",
		Code_random_total_digit:                 16,
		Condition_to_obtain_trx_min_total_spend: 6000000.00,
		Condition_to_obtain_max_count_same_voucher:                    0,
		Condition_to_redeem_trx_has_items_type_ids:                    []int64{4},
		Condition_to_redeem_trx_min_qty_with_same_items_type_ids:      2,
		Condition_to_redeem_trx_min_amount_with_same_items_type_ids:   0.0,
		Voucher_deduction_is_on_total_trx:                             false,
		Voucher_deduction_is_on_total_specific_items_type_ids:         true,
		Voucher_deduction_is_on_one_item_with_specific_items_type_ids: false,
		Voucher_deduction_is_on_percentage:                            true,
		Voucher_deduction_is_constant_amount:                          false,
		Voucher_deduction_percentage:                                  5,
		Voucher_deduction_amount:                                      0.0,
		Expired_add_day_from_generated:                                (time.Hour * 24 * 30),
	},
	Voucher{
		Id:                                      202212002,
		Name:                                    "PREMI",
		Prefix_code:                             "PREMI-",
		Code_random_total_digit:                 16,
		Condition_to_obtain_trx_min_total_spend: 13000000.00,
		Condition_to_obtain_max_count_same_voucher:                    0,
		Condition_to_redeem_trx_has_items_type_ids:                    []int64{1},
		Condition_to_redeem_trx_min_qty_with_same_items_type_ids:      1,
		Condition_to_redeem_trx_min_amount_with_same_items_type_ids:   0.0,
		Voucher_deduction_is_on_total_trx:                             false,
		Voucher_deduction_is_on_total_specific_items_type_ids:         true,
		Voucher_deduction_is_on_one_item_with_specific_items_type_ids: false,
		Voucher_deduction_is_on_percentage:                            true,
		Voucher_deduction_is_constant_amount:                          false,
		Voucher_deduction_percentage:                                  15,
		Voucher_deduction_amount:                                      0.0,
		Expired_add_day_from_generated:                                (time.Hour * 24 * 30),
	},
	Voucher{
		Id:                                      202212003,
		Name:                                    "ULTI",
		Prefix_code:                             "ULTI-",
		Code_random_total_digit:                 16,
		Condition_to_obtain_trx_min_total_spend: 25000000.00,
		Condition_to_obtain_max_count_same_voucher:                    0,
		Condition_to_redeem_trx_has_items_type_ids:                    []int64{3, 4},
		Condition_to_redeem_trx_min_qty_with_same_items_type_ids:      1,
		Condition_to_redeem_trx_min_amount_with_same_items_type_ids:   0.0,
		Voucher_deduction_is_on_total_trx:                             false,
		Voucher_deduction_is_on_total_specific_items_type_ids:         true,
		Voucher_deduction_is_on_one_item_with_specific_items_type_ids: false,
		Voucher_deduction_is_on_percentage:                            true,
		Voucher_deduction_is_constant_amount:                          false,
		Voucher_deduction_percentage:                                  30,
		Voucher_deduction_amount:                                      0.0,
		Expired_add_day_from_generated:                                (time.Hour * 24 * 30),
	},
}
