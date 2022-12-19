package main

import (
	"net/http"
	repository_gorm "salt-final-voucher/internal/repository/gorm"
	pkg_database_gorm_mysql "salt-final-voucher/pkg/database/gorm_mysql"
)

var (

	// ============ Connection to Storage / Cache
	http_client         = http.Client{}
	connectionGormMysql = pkg_database_gorm_mysql.InitDBGormMysql()

	// ============ Repos
	repoCustomersVoucher = repository_gorm.NewRepoCustomersVoucher(connectionGormMysql)

	// ============ Usecasese
)

func main() {

}
