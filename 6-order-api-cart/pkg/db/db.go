package db

import (
	"6-order-api-cart/configs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db struct {
	*gorm.DB
}

func NewDb(conf *configs.Config) *Db {
	db, err := gorm.Open(postgres.Open(conf.DB.DSN))
	if err != nil {
		panic(err)
	}
	return &Db{db}
}
