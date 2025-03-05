package req

import (
	"3-validation-api/configs"
	"3-validation-api/pkg/res"
	"net/http"
)

func HandelBody[T any](w *http.ResponseWriter, r *http.Request, conf *configs.StatusCodes) (*T, error) {
	body, err := Decode[T](r.Body)
	if err != nil {
		res.Json(w, err.Error(), conf.StatusCodeBadRequest)
		return nil, err
	}
	err = IsValid(body)
	if err != nil {
		res.Json(w, err.Error(), conf.StatusCodeBadRequest)
		return nil, err
	}
	return &body, nil
}
