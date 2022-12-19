package repository_gorm_mysql_test

import (
	"context"
	"fmt"
	entity "salt-final-voucher/domain/entity"
	repo_gorm_mysql "salt-final-voucher/internal/repository/gorm"
	pkg_database_gorm_mysql "salt-final-voucher/pkg/database/gorm_mysql"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_CustomersVoucher_Store_Positive(t *testing.T) {
	var (
		ctx                  = context.Background()
		connectionGormMysql  = pkg_database_gorm_mysql.InitDBGormMysql()
		repoCustomersVoucher = repo_gorm_mysql.NewRepoCustomersVoucher(connectionGormMysql)
	)

	dto_customers_voucher := entity.DTOCustomersVoucher{
		Customer_id:           0,
		Voucher_id:            11,
		Voucher_name:          "BASIC",
		Code:                  "BASIC-123123123123",
		Expired_at:            time.Now(),
		Transaction_id:        1,
		Total_amount:          100000.00,
		Total_discount_amount: 50000.00,
		Final_total_amount:    50000.00,
		Status:                121,
	}

	entity_customers_voucher, entity_customers_voucher_err := entity.NewCustomersVoucher(&dto_customers_voucher)
	if entity_customers_voucher_err != nil {
		panic(entity_customers_voucher_err)
	}

	store_err := repoCustomersVoucher.Store(ctx, entity_customers_voucher)

	assert.Nil(t, store_err)
}

func Test_CustomersVoucher_GetByCode_Positive(t *testing.T) {
	var (
		ctx                  = context.Background()
		connectionGormMysql  = pkg_database_gorm_mysql.InitDBGormMysql()
		repoCustomersVoucher = repo_gorm_mysql.NewRepoCustomersVoucher(connectionGormMysql)
	)

	data, get_err := repoCustomersVoucher.GetByCode(ctx, "BASIC-123123123123")
	fmt.Println(data)
	assert.Nil(t, get_err)
	assert.NotNil(t, data)
}

func Test_CustomersVoucher_Update_Positive(t *testing.T) {
	var (
		ctx                  = context.Background()
		connectionGormMysql  = pkg_database_gorm_mysql.InitDBGormMysql()
		repoCustomersVoucher = repo_gorm_mysql.NewRepoCustomersVoucher(connectionGormMysql)
	)

	dto_customers_voucher := entity.DTOCustomersVoucher{
		Id:                    1,
		Customer_id:           0,
		Voucher_id:            11,
		Voucher_name:          "BASIC",
		Code:                  "BASIC-123123123123",
		Expired_at:            time.Now(),
		Transaction_id:        1,
		Total_amount:          100000.00,
		Total_discount_amount: 30000.00,
		Final_total_amount:    70000.00,
		Status:                121,
	}

	entity_customers_voucher, entity_customers_voucher_err := entity.NewCustomersVoucher(&dto_customers_voucher)
	if entity_customers_voucher_err != nil {
		panic(entity_customers_voucher_err)
	}

	update_err := repoCustomersVoucher.Update(ctx, entity_customers_voucher)

	assert.Nil(t, update_err)
}

func Test_CustomersVoucher_Delete_Positive(t *testing.T) {
	var (
		ctx                  = context.Background()
		connectionGormMysql  = pkg_database_gorm_mysql.InitDBGormMysql()
		repoCustomersVoucher = repo_gorm_mysql.NewRepoCustomersVoucher(connectionGormMysql)
	)

	delete_err := repoCustomersVoucher.Delete(ctx, "BASIC-123123123123")
	assert.Nil(t, delete_err)
}
