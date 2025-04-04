package middleware

import (
	"5-order-api-auth/configs"
	"5-order-api-auth/pkg/jwt"
	"context"
	"net/http"
	"strings"
)

type key string

const (
	ContextPhoneNumberKey key = "ContextPhoneNumberKey"
)

func IsAuth(next http.Handler, conf *configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		isValid, data := jwt.NewJwt(conf.Jwt.Secret).Parse(token)
		if !isValid {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
			return
		}

		ctx := context.WithValue(r.Context(), ContextPhoneNumberKey, data.PhoneNumber)
		req := r.WithContext(ctx)
		next.ServeHTTP(w, req)
	})
}
