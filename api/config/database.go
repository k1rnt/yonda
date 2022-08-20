package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "gorm.io/driver/mysql"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	LoadEnvFile()
}

func Connect() *gorm.DB {
	var (
		DbHost = os.Getenv("DB_HOST")
		DbPort = os.Getenv("DB_PORT")
		DbUser = os.Getenv("DB_USER")
		DbPass = os.Getenv("DB_PASS")
		DbName = os.Getenv("DB_NAME")
	)

	db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPass, DbHost, DbPort, DbName))
	if err != nil {
		log.Fatal(err)
	}
	return db
}
