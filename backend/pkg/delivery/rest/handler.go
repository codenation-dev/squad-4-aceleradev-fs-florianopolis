// @SubApi [/users]
// @SubApi Allows you access to different features of the users , login , get status etc [/users]

package rest

import (
	"codenation/squad-4-aceleradev-fs-florianopolis/backend/pkg/adding"
	"codenation/squad-4-aceleradev-fs-florianopolis/backend/pkg/deleting"
	"codenation/squad-4-aceleradev-fs-florianopolis/backend/pkg/reading"
	"codenation/squad-4-aceleradev-fs-florianopolis/backend/pkg/updating"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type serv struct {
	add    adding.Service
	read   reading.Service
	del    deleting.Service
	update updating.Service
}

// Handler handle the API routes
func Handler(
	add adding.Service,
	read reading.Service,
	del deleting.Service,
	update updating.Service,
) *mux.Router {
	s := serv{add: add, read: read, del: del, update: update}

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", getHome).Methods("GET")

	// Get All
	router.HandleFunc("/customer/all", s.getAllCustomers).Methods("GET")
	router.HandleFunc("/user/all", s.getAllUsers).Methods("GET")
	router.HandleFunc("/warning/all", s.getAllWarnings).Methods("GET")

	// Get ByID
	router.HandleFunc("/customer", s.getCustomerByID).Methods("GET").Queries("id", "{id}")
	router.HandleFunc("/user", s.getUserByID).Methods("GET").Queries("id", "{id}")
	router.HandleFunc("/warning", s.getWarningByID).Methods("GET").Queries("id", "{id}")
	router.HandleFunc("/public_func", s.getPublicByID).Methods("GET").Queries("id", "{id}")

	// Get ByName - by pattern
	router.HandleFunc("/customer", s.getCustomerByName).Methods("GET").Queries("name", "{pattern}")
	router.HandleFunc("/user", s.getUserByEmail).Methods("GET").Queries("email", "{pattern}")
	router.HandleFunc("/warning", s.getWarningByCustomer).Methods("GET").Queries("customer", "{pattern}")
	router.HandleFunc("/warning", s.getWarningByUser).Methods("GET").Queries("user", "{pattern}")
	router.HandleFunc("/public_func", s.getPublicByWage).Methods("GET").Queries("wage", "{pattern}")

	// Import
	router.HandleFunc("/customer/import", s.importCustomerFile).Methods("POST")      // TODO: posso fazer opção para escolher o arquivo
	router.HandleFunc("/public_func/import", s.importPublicFuncFile).Methods("POST") // TODO: posso fazer opção para escolher o mês

	// Post
	router.HandleFunc("/customer", s.addCustomer).Methods("POST")
	router.HandleFunc("/user", s.addUser).Methods("POST")
	router.HandleFunc("/warning", s.addWarning).Methods("POST")
	router.HandleFunc("/public_func", s.addPublicFunc).Methods("POST")

	// Delete
	router.HandleFunc("/customer", s.deleteCustomerByID).Methods("DELETE").Queries("id", "{id}")
	router.HandleFunc("/user", s.deleteUserByID).Methods("DELETE").Queries("id", "{id}")
	router.HandleFunc("/warning", s.deleteWarningByID).Methods("DELETE").Queries("id", "{id}")
	router.HandleFunc("/public_func", s.deletePublicByID).Methods("DELETE").Queries("id", "{id}")

	// Put
	router.HandleFunc("/customer", s.updateCustomer).Methods("PUT").Queries("id", "{id}")
	router.HandleFunc("/user", s.updateUser).Methods("PUT").Queries("id", "{id}")
	router.HandleFunc("/warning", s.updateWarning).Methods("PUT").Queries("id", "{id}")

	router.HandleFunc("/swagger.json", swagger).Methods("GET")

	return router
}

func swagger(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	http.ServeFile(w, r, "swagger.json")
}
