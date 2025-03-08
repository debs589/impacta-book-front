package response

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorAPI struct {
	Error string `json:"error"`
}

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if statusCode != http.StatusNoContent {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Fatal(err)
		}
	}
}

func TreatStatusCodeError(w http.ResponseWriter, r *http.Response) {
	var errApi ErrorAPI
	json.NewDecoder(r.Body).Decode(&errApi)
	JSON(w, r.StatusCode, errApi)
}
