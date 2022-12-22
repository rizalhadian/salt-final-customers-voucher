package repository_gorm_mysql

import (
	"context"
	"database/sql"
	"errors"
	"salt-final-voucher/domain/entity"
	"salt-final-voucher/domain/interface_repo"
	mapper_gorm_mysql "salt-final-voucher/internal/repository/gorm/mapper"
	model_gorm_mysql "salt-final-voucher/internal/repository/gorm/models"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type RepoCustomersVoucher struct {
	db *gorm.DB
}

func NewRepoCustomersVoucher(db *gorm.DB) interface_repo.InterfaceRepoCustomersVoucher {
	return &RepoCustomersVoucher{
		db: db,
	}
}

func (rcv *RepoCustomersVoucher) GetRedeemableVoucherByCode(ctx context.Context, code string) (*entity.CustomersVoucher, error) {
	var customers_voucher_get_by_code model_gorm_mysql.ModelCustomersVoucher
	current_time := time.Now()
	result_get := rcv.db.Where("code = ?", code).Where("status = ?", "120").Where("expired_at > ?", current_time.Format(time.RFC3339)).First(&customers_voucher_get_by_code)
	if result_get.Error == gorm.ErrRecordNotFound {
		return nil, errors.New("404")
	}

	entity_customers_voucher := mapper_gorm_mysql.CustomersVoucherModelToEntity(&customers_voucher_get_by_code)
	return entity_customers_voucher, nil
}

func (rcv *RepoCustomersVoucher) GetByVoucherId(ctx context.Context, voucher_id int32) (*entity.CustomersVoucher, error) {
	var customers_voucher_get_by_voucher_id model_gorm_mysql.ModelCustomersVoucher
	result_get := rcv.db.First(&customers_voucher_get_by_voucher_id, "voucher_id = ?", strconv.Itoa(int(voucher_id)))
	if result_get.Error != nil {
		if result_get.Error == gorm.ErrRecordNotFound {
			return nil, errors.New("404")
		} else {
			return nil, errors.New("500")
		}
	}

	entity_customers_voucher := mapper_gorm_mysql.CustomersVoucherModelToEntity(&customers_voucher_get_by_voucher_id)
	return entity_customers_voucher, nil
}

func (rcv *RepoCustomersVoucher) Store(ctx context.Context, customers_voucher *entity.CustomersVoucher) error {

	current_time := time.Now()
	customers_voucher_store := &model_gorm_mysql.ModelCustomersVoucher{
		Customer_id:           customers_voucher.GetCustomerId(),
		Voucher_id:            customers_voucher.GetVoucherId(),
		Voucher_name:          customers_voucher.GetVoucherName(),
		Code:                  customers_voucher.GetCode(),
		Expired_at:            customers_voucher.GetExpiredAt(),
		Transaction_id:        customers_voucher.GetTransactionId(),
		Total_amount:          customers_voucher.GetTotalAmount(),
		Total_discount_amount: customers_voucher.GetTotalDiscountAmount(),
		Final_total_amount:    customers_voucher.GetFinalTotalAmount(),
		Status:                customers_voucher.GetStatus(),
		Created_at:            current_time,
	}

	result_create := rcv.db.Create(customers_voucher_store)
	if result_create.Error != nil {
		return errors.New(result_create.Error.Error())
	}

	return nil
}

func (rcv *RepoCustomersVoucher) Update(ctx context.Context, customers_voucher *entity.CustomersVoucher) error {
	var customers_voucher_get_by_code model_gorm_mysql.ModelCustomersVoucher
	result_get := rcv.db.First(&customers_voucher_get_by_code, "code = ?", customers_voucher.GetCode())
	if result_get.Error == gorm.ErrRecordNotFound {
		return errors.New("404")
	}

	current_time := time.Now()

	customers_voucher_update := &model_gorm_mysql.ModelCustomersVoucher{
		Customer_id:           customers_voucher.GetCustomerId(),
		Voucher_id:            customers_voucher.GetVoucherId(),
		Voucher_name:          customers_voucher.GetVoucherName(),
		Code:                  customers_voucher.GetCode(),
		Expired_at:            customers_voucher.GetExpiredAt(),
		Transaction_id:        customers_voucher.GetTransactionId(),
		Total_amount:          customers_voucher.GetTotalAmount(),
		Total_discount_amount: customers_voucher.GetTotalDiscountAmount(),
		Final_total_amount:    customers_voucher.GetFinalTotalAmount(),
		Status:                customers_voucher.GetStatus(),
		Updated_at:            sql.NullTime{Time: current_time, Valid: true},
	}

	result_update := rcv.db.Model(customers_voucher_get_by_code).Updates(customers_voucher_update)
	if result_update.Error != nil {
		return errors.New(result_update.Error.Error())
	}

	return nil
}

func (rcv *RepoCustomersVoucher) Delete(ctx context.Context, customers_voucher_code string) error {
	var customers_voucher_get_by_code model_gorm_mysql.ModelCustomersVoucher
	result_get := rcv.db.First(&customers_voucher_get_by_code, "code = ?", customers_voucher_code)
	if result_get.Error == gorm.ErrRecordNotFound {
		return errors.New("404")
	}

	result_delete := rcv.db.Delete(&customers_voucher_get_by_code, customers_voucher_get_by_code.Id)
	if result_delete.Error != nil {
		return errors.New(result_delete.Error.Error())
	}

	return nil
}
