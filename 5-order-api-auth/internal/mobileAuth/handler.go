package mobileauth

import (
	"5-order-api-auth/configs"
	"5-order-api-auth/pkg/jwt"
	"5-order-api-auth/pkg/middleware"
	"5-order-api-auth/pkg/req"
	"5-order-api-auth/pkg/res"
	"net/http"
)

type MobileAuthHandler struct {
	*MobileAuthRepository
	*jwt.Jwt
}

type MobileAuthHandlerDeps struct {
	*MobileAuthRepository
	*configs.Config
	*jwt.Jwt
}

func NewMobileAuthHandler(router *http.ServeMux, deps MobileAuthHandlerDeps) {
	handler := &MobileAuthHandler{
		MobileAuthRepository: deps.MobileAuthRepository,
		Jwt:                  deps.Jwt,
	}
	router.HandleFunc("GET /auth/mobile/registery/{phoneNumber}", handler.Registery())
	router.HandleFunc("POST /auth/mobile/verify", handler.Verify())
	router.Handle("GET /someusefull", middleware.IsAuth(handler.SomeUsefull(), deps.Config))
}

func (handler *MobileAuthHandler) Registery() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		phoneNumber := r.PathValue("phoneNumber")
		user := NewMobileAuthUser(phoneNumber)

		err := req.IsValid(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = handler.MobileAuthRepository.WriteSessionId(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		res.Json(&w, SessionIdResponse{user.SessionId}, http.StatusOK)
	}
}

func (handler *MobileAuthHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[MobileVerifyRequest](&w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		userAuth, err := handler.MobileAuthRepository.FindBySessionId(body.SessionId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		token, err := handler.Jwt.Create(userAuth.SessionId, userAuth.Phone)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		userVerify, err := handler.MobileAuthRepository.WriteToken(&MobileVerifyUser{
			SessionId:  body.SessionId,
			VerifyCode: body.Code,
			Token:      token,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		res.Json(&w, MobileVerifyResponse{userVerify.Token}, http.StatusOK)
	}
}

func (handler *MobileAuthHandler) SomeUsefull() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		phoneNumber := r.Context().Value(middleware.ContextPhoneNumberKey).(string)
		res.Json(&w, SomeUsefullResponse{PhoneNumber: phoneNumber}, http.StatusOK)
	}
}
