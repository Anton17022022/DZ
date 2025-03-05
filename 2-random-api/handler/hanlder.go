package handler

import (
	"fmt"
	"math/rand"
	"net/http"
)

type handler struct {
}

func NewHanlder(router *http.ServeMux) {
	handler := &handler{}
	router.HandleFunc("/", handler.randOneToSix())
}

func (handler *handler) randOneToSix() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(fmt.Sprintf("%d", rand.Intn(6)+1)))
	}
}
