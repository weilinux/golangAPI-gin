package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)


var (
	DBConnect *gorm.DB
)

var err error

func Connect() {
	dsn := "gorm:gorm@tcp(localhost:3306)/simplerest?charset=utf8&parseTime=True&loc=Local"
	DBConnect, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
}