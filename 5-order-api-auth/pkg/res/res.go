package res

import (
	"encoding/json"
	"net/http"
)

func Json(w *http.ResponseWriter, response any, statusCode int) {
	(*w).Header().Set("Conntent-Type", "application/json")
	(*w).WriteHeader(statusCode)
	json.NewEncoder(*w).Encode(response)
}
