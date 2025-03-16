package main

import (
	"4-order-api/internal/product"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}
	db, err := gorm.Open(postgres.Open(os.Getenv("DSN_PG_Orders")), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	db.AutoMigrate(&product.Product{})
}
