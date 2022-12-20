package main

import (
	"fmt"
	"log"
	"net/http"
	http_handler "salt-final-voucher/internal/delivery/http/handler"
	infrastructure_transaction "salt-final-voucher/internal/infrastructure/transaction"
	repository_gorm "salt-final-voucher/internal/repository/gorm"
	"salt-final-voucher/internal/usecase"
	pkg_database_gorm_mysql "salt-final-voucher/pkg/database/gorm_mysql"
	"time"

	"github.com/gorilla/mux"
)

var (

	// ============ Connection to Storage / Cache
	http_client         = http.Client{}
	connectionGormMysql = pkg_database_gorm_mysql.InitDBGormMysql()
	// ============ Infrastructure
	infrastructureTransaction = infrastructure_transaction.NewInfrastructureTransaction(http_client, "http://127.0.0.2:8000/api/customer/{customer_id}/transaction-count")
	// ============ Repos
	repoCustomersVoucher = repository_gorm.NewRepoCustomersVoucher(connectionGormMysql)
	// ============ Usecasese
	usecaseCustomersVoucher = usecase.NewUsecaseCustomerVoucher(infrastructureTransaction, repoCustomersVoucher)
)

func main() {
	router := mux.NewRouter()

	http_handler.NewHandlerCustomersVoucher(router, usecaseCustomersVoucher)

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.3:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("Run on 127.0.0.3:8000")

	log.Fatal(srv.ListenAndServe())
}
