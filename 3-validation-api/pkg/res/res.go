package res

import (
	"encoding/json"
	"fmt"
	"net/http"
	"log"
)

func Json(w *http.ResponseWriter, response interface{}, statusCode string) {
	var status int
	_, err := fmt.Sscanf(statusCode, "%d", &status)
	if err != nil {
		log.Panic(err.Error())
	}
	(*w).Header().Set("Content-Type", "application/json")
	(*w).WriteHeader(status)
	json.NewEncoder(*w).Encode(response)
}
