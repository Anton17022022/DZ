package main

import (
	"5-order-api-auth/configs"
	"5-order-api-auth/internal/mobileAuth"
	"5-order-api-auth/pkg/db"
	"5-order-api-auth/pkg/jwt"
	"log"
	"net/http"
)

func main() {
	// Packages
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()
	jwt := jwt.NewJwt(conf.Jwt.Secret)

	// Repositories
	mobileAuthRepository := mobileauth.NewMobileAuthRepository(db)

	// Handlers
	mobileauth.NewMobileAuthHandler(router, mobileauth.MobileAuthHandlerDeps{
		MobileAuthRepository: mobileAuthRepository,
		Config: conf,
		Jwt: jwt,
	})

	server := http.Server{
		Addr:    ":8082",
		Handler: router,
	}
	
	log.Println("server listening on port 8082")
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
