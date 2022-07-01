package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "gorm.io/driver/mysql"
	"log"
	"os"
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
		DB_HOST = os.Getenv("DB_HOST")
		DB_PORT = os.Getenv("DB_PORT")
		DB_USER = os.Getenv("DB_USER")
		DB_PASS = os.Getenv("DB_PASS")
		DB_NAME = os.Getenv("DB_NAME")
	)

	db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME))
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func Close() {
	if err := db.Close(); err != nil {
		log.Fatal(err)
	}
}
