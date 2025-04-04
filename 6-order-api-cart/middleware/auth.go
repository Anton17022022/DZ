package middleware

import (
	"6-order-api-cart/configs"
	"6-order-api-cart/internal/user"
	"6-order-api-cart/pkg/jwt"
	"context"
	"net/http"
	"strings"
)

type key string

const (
	BearerPrefix     = "Bearer "
	ContextUser  key = "UserAuthed"
)

func IsAuth(next http.Handler, conf configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authorizationHeader, BearerPrefix) {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
			return
		}

		token := strings.TrimPrefix(authorizationHeader, BearerPrefix)
		isValid, data := jwt.NewJwt(conf.Jwt.Secret).Parse(token)
		if !isValid {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
			return
		}
		// TODO проверка на протухание токена
		ctx := context.WithValue(r.Context(), ContextUser, &user.User{PhoneNumber: data.PhoneNumber})
		req := r.WithContext(ctx)
		next.ServeHTTP(w, req)
	})
}
