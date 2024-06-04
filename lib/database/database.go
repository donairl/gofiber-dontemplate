package database

import (
	"log"

	"ariga.io/atlas/sql/mysql"
	"gorm.io/gorm"
)

var Connection *gorm.DB

func ConnectDb() {
	// Connect to the database
	dsn := "user:password@tcp(localhost:3306)/your_database_name?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	//defer db.Close()

	Connection = db

}
