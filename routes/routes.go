package routes

import (
	"codenation/squad-4-aceleradev-fs-florianopolis/controllers"

	"github.com/gorilla/mux"
)

// NewRouter makes a new mux.Router
func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", controllers.GetHome).Methods("GET")
	return r
}
