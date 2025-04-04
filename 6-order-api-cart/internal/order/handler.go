package order

import (
	"6-order-api-cart/configs"
	"6-order-api-cart/pkg/middleware"
	"6-order-api-cart/pkg/req"
	"6-order-api-cart/pkg/res"
	"net/http"
	"strconv"
)

type OrderHandler struct {
	*OrderRepository
}

type OrderHandlerDeps struct {
	*OrderRepository
	*configs.Config
}

func NewOrderHandler(router *http.ServeMux, deps *OrderHandlerDeps) {
	handler := &OrderHandler{
		OrderRepository: deps.OrderRepository,
	}

	router.Handle("POST /order", middleware.IsAuth(handler.Create(), deps.Config))
	router.Handle("GET /order/{id}", middleware.IsAuth(handler.GetOrderById(), deps.Config))
	router.Handle("GET /my-orders", middleware.IsAuth(handler.GetAllOrdersByUser(), deps.Config))
}

func (handler *OrderHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[OrderCreateRequest](r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		UserPhoneNumber, ok := r.Context().Value(middleware.ContextPhoneNumberKey).(string)
		if !ok {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		order := NewOrder(body.ProductName, UserPhoneNumber, body.Amount)
		err = handler.OrderRepository.Create(order)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		res.Json(&w, OrderCreateResponse{
			OrderId: strconv.FormatUint(uint64(order.Model.ID), 10),
		})
	}
}

func (handler *OrderHandler) GetOrderById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		stringId := r.PathValue("id")
		id, err := strconv.ParseUint(stringId, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		order, err := handler.OrderRepository.GetOrderByID(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		contextUserPhoneNumber, ok := r.Context().Value(middleware.ContextPhoneNumberKey).(string)
		if !ok {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		if order.UserPhoneNumber != contextUserPhoneNumber {
			http.Error(w, "not user order", http.StatusBadRequest)
			return
		}

		res.Json(&w, OrderGetByIdResponse{
			ProductName: order.ProductName,
			Amount:      order.Amount,
		})
	}
}

func (handler *OrderHandler) GetAllOrdersByUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		UserPhoneNumber, ok := r.Context().Value(middleware.ContextPhoneNumberKey).(string)
		if !ok {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		orders, err := handler.OrderRepository.GetOrdersByPhoneNumber(UserPhoneNumber)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		res.Json(&w, OrderGetAllOrdersByUserResponse{
			Orders: orders,
		})
	}
}
