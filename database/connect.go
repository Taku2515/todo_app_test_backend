package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func init() {
	dsn := "tester:password@tcp(db:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(dsn + "database can't connect")
	}
}