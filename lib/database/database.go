package database

import (
	"log"

	"github.com/donairl/gofiber-dontemplate/models"
	"gorm.io/gorm"
)

var Connection *gorm.DB

func ConnectDb() {
	// Connect to the database
	db, err := gorm.Open("mysql", "user:password@tcp(localhost:3306)/your_database_name?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Migrate the User model
	db.AutoMigrate(&models.User{})

	Connection = db

}
