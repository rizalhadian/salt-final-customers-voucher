package usecase_test

import (
	"context"
	"net/http"
	infrastructure_transaction "salt-final-voucher/internal/infrastructure/transaction"
	repository_gorm "salt-final-voucher/internal/repository/gorm"
	"salt-final-voucher/internal/usecase"
	pkg_database_gorm_mysql "salt-final-voucher/pkg/database/gorm_mysql"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CustomersVoucher_Generate_Positive(t *testing.T) {
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
	resp, err := usecaseCustomersVoucher.Generate(ctx, 0)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
}
