package res

import (
	"encoding/json"
	"net/http"
)

func Json(w *http.ResponseWriter, response interface{}){
	(*w).Header().Set("Content-Type", "application/json")
	(*w).WriteHeader(http.StatusOK)
	json.NewEncoder(*w).Encode(response)
}