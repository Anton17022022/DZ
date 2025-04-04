package user

import (
	"6-order-api-cart/pkg/jwt"
	"6-order-api-cart/pkg/req"
	"6-order-api-cart/pkg/res"
	"net/http"
)

type UserHanlder struct {
	*UserRepository
	*jwt.Jwt
}

type UserHandlerDeps struct {
	*UserRepository
	*jwt.Jwt
}

func NewUserHandler(router *http.ServeMux, deps *UserHandlerDeps) {
	handler := &UserHanlder{
		UserRepository: deps.UserRepository,
		Jwt:            deps.Jwt,
	}
	router.HandleFunc("POST /user/auth", handler.Auth())
}

func (handler *UserHanlder) Auth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userRequest, err := req.HandleBody[UserRequest](r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		user := NewUser(userRequest.PhoneNumber)
		user, err = handler.UserRepository.Auth(user, handler.Jwt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		res.Json(&w, UserResponse{Token: user.Token})
	}
}
