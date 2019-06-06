package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

// DefaultResponse implements a struct to use in response data
type DefaultResponse struct {
	Data   interface{} `json:"data"`
	Status int         `Json:"status"`
}

// ErrorResponse send the error message to http.ResponseWriter
func ErrorResponse(w http.ResponseWriter, err error, status int) {
	w.WriteHeader(status)
	ToJSON(w, struct {
		Message string `json:"message"`
	}{
		Message: err.Error(),
	})
}

// ToJSON encodes a struct in a json format and
// saves it in an w object
func ToJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Fatal(err)
	}
}
