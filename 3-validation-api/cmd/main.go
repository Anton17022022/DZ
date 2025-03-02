package main

import (
	"3-validation-api/configs"
	"3-validation-api/internal/temporarydb"
	verify "3-validation-api/internal/verify"
	"log"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	db := *temporarydb.NewTemporaryDb()
	router := http.NewServeMux()
	verify.NewVerifyHandler(
		router,
		verify.VerifyHandlerDeps{
			Config: conf,
		},
		db,
	)

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Panic(err.Error())
	}
}
