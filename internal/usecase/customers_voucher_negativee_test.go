package usecase_test

import (
	"context"
	"net/http"
	"salt-final-voucher/domain/entity"
	infrastructure_transaction "salt-final-voucher/internal/infrastructure/transaction"
	repository_gorm "salt-final-voucher/internal/repository/gorm"
	"salt-final-voucher/internal/usecase"
	pkg_database_gorm_mysql "salt-final-voucher/pkg/database/gorm_mysql"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CustomersVoucher_Generate_Negative(t *testing.T) {
	var (
		ctx = context.Background()
		// ============ Connection to Storage / Cache
		http_client         = http.Client{}
		connectionGormMysql = pkg_database_gorm_mysql.InitDBGormMysql()
		// ============ Infrastructue
		infrastructureTransaction = infrastructure_transaction.NewInfrastructureTransaction(http_client, "http://localhost:8000/api/customer/{customer_id}/transaction-count")
		// ============ Repos
		repoCustomersVoucher = repository_gorm.NewRepoCustomersVoucher(connectionGormMysql)
		// ============ Usecasese
		usecaseCustomersVoucher = usecase.NewUsecaseCustomerVoucher(infrastructureTransaction, repoCustomersVoucher)
	)
	resp, err := usecaseCustomersVoucher.Generate(ctx, 999)
	assert.Nil(t, resp)
	assert.NotNil(t, err)
}

func Test_CustomersVoucher_Redeem_Negative(t *testing.T) {
	var (
		ctx = context.Background()
		// ============ Connection to Storage / Cache
		http_client         = http.Client{}
		connectionGormMysql = pkg_database_gorm_mysql.InitDBGormMysql()
		// ============ Infrastructue
		infrastructureTransaction = infrastructure_transaction.NewInfrastructureTransaction(http_client, "http://localhost:8000/api/customer/{customer_id}/transaction-count")
		// ============ Repos
		repoCustomersVoucher = repository_gorm.NewRepoCustomersVoucher(connectionGormMysql)
		// ============ Usecasese
		usecaseCustomersVoucher = usecase.NewUsecaseCustomerVoucher(infrastructureTransaction, repoCustomersVoucher)
	)

	dto_transaction := &entity.DTOTransaction{
		Id:           3,
		Customer_id:  0,
		Note:         "",
		Total_amount: 200000.00,
	}

	dto_items := []*entity.DTOTransactionsItem{
		&entity.DTOTransactionsItem{
			Item_id:       10,
			Items_type_id: 4,
			Price:         100000.00,
			Qty:           1,
			Total_price:   100000.00,
			Note:          "Note",
		},
		&entity.DTOTransactionsItem{
			Item_id:       11,
			Items_type_id: 4,
			Price:         100000.00,
			Qty:           1,
			Total_price:   100000.00,
			Note:          "Note",
		},
		&entity.DTOTransactionsItem{
			Item_id:       7,
			Items_type_id: 3,
			Price:         100000.00,
			Qty:           1,
			Total_price:   100000.00,
			Note:          "Note",
		},

		&entity.DTOTransactionsItem{
			Item_id:       3,
			Items_type_id: 1,
			Price:         100000.00,
			Qty:           1,
			Total_price:   100000.00,
			Note:          "Note",
		},
		&entity.DTOTransactionsItem{
			Item_id:       14,
			Items_type_id: 1,
			Price:         100000.00,
			Qty:           1,
			Total_price:   100000.00,
			Note:          "Note",
		},
		&entity.DTOTransactionsItem{
			Item_id:       15,
			Items_type_id: 1,
			Price:         100000.00,
			Qty:           1,
			Total_price:   100000.00,
			Note:          "Note",
		},
		&entity.DTOTransactionsItem{
			Item_id:       16,
			Items_type_id: 1,
			Price:         100000.00,
			Qty:           1,
			Total_price:   100000.00,
			Note:          "Note",
		},
		&entity.DTOTransactionsItem{
			Item_id:       17,
			Items_type_id: 1,
			Price:         100000.00,
			Qty:           1,
			Total_price:   100000.00,
			Note:          "Note",
		},
	}

	dto_vouchers := []*entity.DTOCustomersVoucher{
		&entity.DTOCustomersVoucher{
			Code: "XXXBASIC-7006791947780029",
		},
		&entity.DTOCustomersVoucher{
			Code: "XXXULTI-1484611666146502",
		},
		&entity.DTOCustomersVoucher{
			Code: "XXXPREMI-8665223082154514",
		},
	}

	resp, resp_err := usecaseCustomersVoucher.Redeem(ctx, dto_transaction, dto_items, dto_vouchers)
	assert.NotNil(t, resp_err)
	assert.Nil(t, resp)
}
