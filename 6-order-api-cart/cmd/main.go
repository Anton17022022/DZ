package main

import (
	"6-order-api-cart/configs"
	"6-order-api-cart/internal/order"
	"6-order-api-cart/internal/product"
	"6-order-api-cart/internal/user"
	"6-order-api-cart/pkg/db"
	"6-order-api-cart/pkg/jwt"
	"log"
	"net/http"
)

func main() {
	// TODO логирование ошибок одельным потоком
	// TODO сбор/отправка потребной статистики отдельным потоком. Например по проведенному времени на старнице или стастистики заказов по промежуткам времени

	// Packagees
	conf := configs.NewConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()
	jwt := jwt.NewJwt(conf.Jwt.Secret)

	// Repositoryes
	userRepository := user.NewUseRepository(db)
	productRepository := product.NewProductRepository(db)
	orderRepository := order.NewOrderRepository(&order.OrderRepositoryDeps{
		Db:                db,
		ProductRepository: productRepository,
	})

	// Handlers
	user.NewUserHandler(router, &user.UserHandlerDeps{
		UserRepository: userRepository,
		Jwt:            jwt,
	})
	product.NewProductHandler(router, &product.ProductHandlerDeps{
		ProductRepository: productRepository,
		Config:            conf,
	})
	order.NewOrderHandler(router, &order.OrderHandlerDeps{
		OrderRepository: orderRepository,
		Config:          conf,
	})

	server := http.Server{
		Addr:    ":8083",
		Handler: router,
	}

	log.Println("server start on port 8083")
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
