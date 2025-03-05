package main

import (
	"2-random-api/handler"
	"fmt"
	"net/http"
)

func main() {
	router := &http.ServeMux{}
	handler.NewHanlder(router)

	server := http.Server{
		Addr: ":8080",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err.Error())
	}
}
