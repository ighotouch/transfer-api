package main

import (
	"fmt"
	"transfer-api/bank"
	"transfer-api/bank/database"
	"transfer-api/bank/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDatabase() {
	var err error
	dsn := "root:rootroot@tcp(127.0.0.1:3306)/transfer?charset=utf8mb4&parseTime=True&loc=Local"
	database.DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Database connected successfully")

	database.DBConn.AutoMigrate(&models.Transfer{})
	database.DBConn.AutoMigrate(&models.Account{})
	fmt.Println("Database Migrated")
}

func main() {
	// initDatabase()
	// defer database.DBConn.
	bank.Start()
}
