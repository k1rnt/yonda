package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "gorm.io/driver/mysql"
	"log"
)

const (
	DB_HOST = "localhost"
	DB_PORT = "3306"
	DB_USER = "k1rnt"
	DB_PASS = "k1rnt_pass"
	DB_NAME = "yonda"
)

var (
	db  *gorm.DB
	err error
)

func Connect() *gorm.DB {
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
