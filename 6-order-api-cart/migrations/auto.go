package main

import (
	"6-order-api-cart/configs"
	"6-order-api-cart/internal/order"
	"6-order-api-cart/internal/product"
	"6-order-api-cart/internal/user"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	conf := configs.NewConfig()
	db, err := gorm.Open(postgres.Open(conf.DB.DSN))
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&user.User{}, &user.UserAuth{}, &product.Product{}, &order.Order{})
	if err != nil {
		panic(err)
	}
}
