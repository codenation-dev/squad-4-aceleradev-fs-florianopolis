package routes

import (
	"codenation/squad-4-aceleradev-fs-florianopolis/controllers"

	"github.com/gorilla/mux"
)

// NewRouter makes a new mux.Router
func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", controllers.GetHome).Methods("GET")
	r.HandleFunc("/customers", controllers.GetCustomers).Methods("GET")
	r.HandleFunc("/customers/public", controllers.GetCustomersPublicFuncs).Methods("GET")
	r.HandleFunc("/customers/{wage}", controllers.GetVIPCustomers).Methods("GET")
	// TODO: Como implementar o name??? Esta confundindo com o wage em cima
	r.HandleFunc("/customers/{name}", controllers.GetCustomerByName).Methods("GET")
	r.HandleFunc("/customers", controllers.PostCustomer).Methods("POST")
	r.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	r.HandleFunc("/users", controllers.PostUser).Methods("POST")
	r.HandleFunc("/warnings", controllers.GetWarnings).Methods("GET")
	r.HandleFunc("/warnings", controllers.PostWarning).Methods("POST")
	r.HandleFunc("/warnings/{sent_to}", controllers.GetWarningsByUser).Methods("GET")
	return r
}
