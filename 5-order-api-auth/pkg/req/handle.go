package req

import (
	"5-order-api-auth/pkg/res"
	"net/http"
)

const (
	statusCodeBadRequest = 400
)

func HandleBody[T any](w *http.ResponseWriter, r *http.Request) (*T, error) {
	body, err := Decode[T](r.Body)
	if err != nil {
		res.Json(w, err.Error(), statusCodeBadRequest)
		return nil, err
	}
	err = IsValid(body)
	if err != nil {
		res.Json(w, err.Error(), statusCodeBadRequest)
		return nil, err
	}
	return &body, nil
}
