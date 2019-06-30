// Package rest implementa os endpoints da API Rest.
package rest

import (
	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/backend/pkg/adding"
	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/backend/pkg/deleting"
	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/backend/pkg/reading"
	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/backend/pkg/updating"

	"github.com/gorilla/mux"
)

type Serv struct {
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
	s := Serv{add: add, read: read, del: del, update: update}

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
	router.HandleFunc("/customer/import", s.ImportCustomerFile).Methods("POST")      // TODO: posso fazer opção para escolher o arquivo
	router.HandleFunc("/public_func/import", s.ImportPublicFuncFile).Methods("POST") // TODO: posso fazer opção para escolher o mês

	// Post
	router.HandleFunc("/customer", s.AddCustomer).Methods("POST")
	router.HandleFunc("/user", s.addUser).Methods("POST")
	router.HandleFunc("/warning", s.AddWarning).Methods("POST")
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

	return router
}
