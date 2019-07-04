package model

import (
	"encoding/json"
	"log"
	"net/http"
)

// DefaultResponse models the response data
type DefaultResponse struct {
	Data   interface{} `json:"data"`
	Status int         `json:"status"`
}

// ErrorResponse implements the standard error response
func ErrorResponse(w http.ResponseWriter, err error, status int) {
	w.WriteHeader(status)
	ToJSON(w, struct {
		Message string `json:"message"`
	}{
		Message: err.Error(),
	})
}

// ToJSON encodes the answer to json format
func ToJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Fatal(err)
	}
}

// func Message(status bool, message string) map[string]interface{} {
// 	return map[string]interface{}{"status": status, "message": message}
// }

// func Respond(w http.ResponseWriter, data map[string]interface{}) {
// 	w.Header().Add("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(data)
// }
