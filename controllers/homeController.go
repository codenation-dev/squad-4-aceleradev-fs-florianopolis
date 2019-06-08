package controllers

import (
	"codenation/squad-4-aceleradev-fs-florianopolis/utils"
	"net/http"
)

// GetHome makes a defaul message to root
func GetHome(w http.ResponseWriter, r *http.Request) {
	utils.ToJSON(w, struct {
		Message string `json:"message"`
	}{
		Message: "Go RESTful API",
	})
}
