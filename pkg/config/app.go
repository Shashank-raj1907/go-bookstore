package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "shashank:sql.Shash2@tcp(127.0.0.1:3306)/simplerest?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic("Failed to connect to database error: " + err.Error())
	}
	db = d
}

func GetDB() *gorm.DB {
	if db == nil {
		Connect()
	}
	return db
}
