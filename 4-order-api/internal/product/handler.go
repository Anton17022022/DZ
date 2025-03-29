package product

import (
	"4-order-api/configs"
	"4-order-api/pkg/req"
	"4-order-api/pkg/res"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

type ProductHandler struct {
	*configs.Config
	*ProductRepository
}

type ProductHandlerDeps struct {
	*configs.Config
	*ProductRepository
}

func NewProductHandler(router *http.ServeMux, deps *ProductHandlerDeps) {
	handler := &ProductHandler{
		Config:            deps.Config,
		ProductRepository: deps.ProductRepository,
	}
	router.HandleFunc("POST /product/create", handler.Create())
	router.HandleFunc("PATCH /product/update/{id}", handler.Update())
	router.HandleFunc("DELETE /product/delete/{id}", handler.Delete())
	router.HandleFunc("GET /product/get/{id}", handler.GetById())
}

func (handler *ProductHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandelBody[ProductCreateRequest](&w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		product := NewProduct(body.Name, body.Description, body.Images)
		err = handler.ProductRepository.GetByName(product.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = handler.ProductRepository.Create(product)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		res.Json(&w, product, http.StatusCreated)
	}
}

func (handler *ProductHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandelBody[ProductCreateRequest](&w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		stringId := r.PathValue("id")
		id, err := strconv.ParseUint(stringId, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		product, err := handler.ProductRepository.Update(&Product{
			Model:       gorm.Model{ID: uint(id)},
			Name:        body.Name,
			Description: body.Description,
			Images:      body.Images,
		})
		if err != nil && err.Error() == "not founded id" {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		} else if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		res.Json(&w, product, http.StatusOK)
	}
}

func (handler *ProductHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		stringId := r.PathValue("id")
		id, err := strconv.ParseUint(stringId, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = handler.ProductRepository.Delete(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		res.Json(&w, "Deleted", http.StatusOK)
	}
}

func (handler *ProductHandler) GetById() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		stringId := r.PathValue("id")
		id, err :=  strconv.ParseUint(stringId, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		product, err := handler.ProductRepository.GetById(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		res.Json(&w, product, http.StatusOK)
	}
}
