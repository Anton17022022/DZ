package main

import (
	mobileauth "5-order-api-auth/internal/mobileAuth"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main(){
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Error %s. Using default value.", err.Error())
	}
	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")))
	if err != nil {
		//TODO: обработка ошибок и ретраи
		panic(err)
	}
	db.AutoMigrate(&mobileauth.MobileAuthUser{}, &mobileauth.MobileVerifyUser{})
}