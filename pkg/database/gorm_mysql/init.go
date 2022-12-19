package pkg_database_gorm_mysql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDBGormMysql() *gorm.DB {
	dbHost := "127.0.0.1"
	dbPort := "3306"
	dbUser := "root"
	dbPass := ""
	dbName := "2022_salt_final"

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(connection), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	return db
}
