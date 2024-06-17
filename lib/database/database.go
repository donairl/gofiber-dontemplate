package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Connection *gorm.DB

func ConnectDb() {

	// Connect to the database

	dsn := "admin:asdf1234%@tcp(localhost:3306)/donwebgolang?charset=utf8&parseTime=True&loc=Local"
	db, err := xgorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	//defer db.Close()

	Connection = db

}

func CreateDatabase(name string) error {
	dsn := "admin:asdf1234%@tcp(localhost:3306)/?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to create database: %w", err)
	}

	// Create the database
	err = db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s`;", name)).Error
	if err != nil {
		return fmt.Errorf("failed to create database: %w", err)
	}

	log.Printf("Database '%s' created successfully.", name)
	return nil
}
