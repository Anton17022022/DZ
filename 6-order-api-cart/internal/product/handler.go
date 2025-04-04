package product

import (
	"6-order-api-cart/configs"
	"6-order-api-cart/middleware"
	"6-order-api-cart/pkg/req"
	"6-order-api-cart/pkg/res"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

type ProductHandler struct {
	*ProductRepository
}

type ProductHandlerDeps struct {
	*ProductRepository
	*configs.Config
}

func NewProductHandler(router *http.ServeMux, deps *ProductHandlerDeps) {
	handler := &ProductHandler{
		ProductRepository: deps.ProductRepository,
	}

	router.Handle("POST /product/create", middleware.IsAuth(handler.Create(), *deps.Config))
	router.Handle("PATCH /product/update", middleware.IsAuth(handler.Update(), *deps.Config))
	router.Handle("DELETE /product/delete/{name}", middleware.IsAuth(handler.Delete(), *deps.Config))
	router.Handle("GET /products", middleware.IsAuth(handler.GetAll(), *deps.Config))
}

func (handler *ProductHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO добавление других параметров. Например: цена, описание и т.д.
		body, err := req.HandleBody[ProductCreateRequest](r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		product, err := handler.ProductRepository.Create(body.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		res.Json(&w, ProductCreateResponse{
			Name: product.Name,
			Id:   strconv.FormatUint(uint64(product.ID), 10),
		})
	}
}

func (handler *ProductHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[ProductUpdateRequest](r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		id, err := strconv.ParseUint(body.Id, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		product, err := handler.ProductRepository.Update(&Product{
			Model: gorm.Model{ID: uint(id)},
			Name:  body.Name,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		res.Json(&w, ProductUpdateResponse{
			ProductCreateResponse: ProductCreateResponse{
				Name: product.Name,
				Id:   strconv.FormatUint(uint64(product.ID), 10),
			},
		})
	}
}

func (handler *ProductHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.PathValue("name")

		err := handler.ProductRepository.Delete(name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		res.Json(&w, ProductDeleteResponse{
			Message: "product deleted",
		})
	}
}

func (handler *ProductHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		products, err := handler.ProductRepository.GetAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		res.Json(&w, ProductGetAllResponse{Products: products})
	}
}
