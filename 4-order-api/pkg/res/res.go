package res

import (
	"encoding/json"
	"net/http"
)

func Json(w *http.ResponseWriter, response interface{}, statusCode int) {
	(*w).Header().Set("Content-Type", "application/json")
	(*w).WriteHeader(statusCode)
	json.NewEncoder(*w).Encode(response)
}
