package db

import (
	"4-order-api/configs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db struct{
	*gorm.DB
}

func NewDB(conf configs.Config) *Db {
	db, err := gorm.Open(postgres.Open(conf.Db.Db), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	return &Db{db}
}