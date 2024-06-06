package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Connection *gorm.DB

func ConnectDb() {

	// Connect to the database
	dsn := "admin:asdf1234%@tcp(localhost:3306)/donwebgolang?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	//defer db.Close()

	Connection = db

}
