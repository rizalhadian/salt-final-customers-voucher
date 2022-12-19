package model_gorm_mysql

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type ModelCustomersVoucher struct {
	gorm.Model
	Id                    int64 `gorm:"primaryKey"`
	Customer_id           int64
	Voucher_id            int64
	Voucher_name          string
	Code                  string
	Expired_at            time.Time
	Transaction_id        int64
	Total_amount          float64
	Total_discount_amount float64
	Final_total_amount    float64
	Status                int16
	Created_at            time.Time
	Updated_at            sql.NullTime
	Deleted_at            sql.NullTime
}

func (ModelCustomersVoucher) TableName() string {
	return "customers_voucher"
}

// func (mcv *ModelCustomersVoucher) GetValueFromDbqTag(tag string) (any, error) {
// 	switch {
// 	case tag == "id":
// 		return mcv.Id, nil
// 	case tag == "customer_id":
// 		return mcv.Customer_id, nil
// 	case tag == "voucher_id":
// 		return mcv.Voucher_id, nil
// 	case tag == "voucher_name":
// 		return mcv.Voucher_name, nil
// 	default:
// 		return nil, errors.New("Value Not Found")
// 	}
// }

// func (ModelCustomersVoucher) GetFieldsNeededToStoreProcess() []string {
// 	return []string{
// 		"customer_id",
// 		"voucher_id",
// 		"voucher_name",
// 		"code",
// 		"expired_at",
// 		"transaction_id",
// 		"total_amount",
// 		"total_discount_amount",
// 		"final_total_amount",
// 		"status",
// 		"created_at",
// 	}
// }

// func (ModelCustomersVoucher) GetFieldsNeededToGetProcess() []string {
// 	return []string{
// 		"id",
// 		"customer_id",
// 		"voucher_id",
// 		"voucher_name",
// 		"code",
// 		"expired_at",
// 		"transaction_id",
// 		"total_amount",
// 		"total_discount_amount",
// 		"final_total_amount",
// 		"status",
// 		"created_at",
// 		"updated_at",
// 	}
// }

// func (ModelCustomersVoucher) GetFieldsNeededToUpdateProcess() []string {
// 	return []string{
// 		"voucher_id",
// 		"voucher_name",
// 		"code",
// 		"expired_at",
// 		"transaction_id",
// 		"total_amount",
// 		"total_discount_amount",
// 		"final_total_amount",
// 		"status",
// 		"updated_at",
// 	}
// }

// func (ModelCustomersVoucher) GetFieldsNeededToSoftDeleteProcess() []string {
// 	return []string{
// 		"deleted_at",
// 	}
// }
