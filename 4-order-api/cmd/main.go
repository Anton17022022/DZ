package main

import (
	"4-order-api/configs"
	"4-order-api/internal/product"
	"4-order-api/pkg/db"
	"4-order-api/pkg/middleware"
	"log"
	"net/http"

	"github.com/sirupsen/logrus"
)

func main() {
	conf := configs.LoadConfig()
	db := db.NewDB(*conf)

	router := http.NewServeMux()

	// handlres
	product.NewProductHandler(router, &product.ProductHandlerDeps{
		Config: conf,
		ProductRepository: &product.ProductRepository{
			Database: db,
		},
	})

	// Logging
	logJson := logrus.New()
	logJson.SetFormatter(&logrus.JSONFormatter{})
	middlewareLoging := middleware.NewMiddlewareLoging(logJson)

	// Middleware
	stack := middleware.Chain(
		middlewareLoging.Logging,
	)

	server := http.Server{
		Addr:    ":8080",
		Handler: stack(router),
	}

	log.Println("Server listening on port 8080")
	err := server.ListenAndServe()
	if err != nil {
		log.Println(err.Error())
	}
}
