// Package rest implementa os endpoints da API Rest.
package rest

import (
	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/backend/pkg/adding"
	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/backend/pkg/deleting"
	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/backend/pkg/reading"
	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/backend/pkg/updating"

	"github.com/gorilla/mux"
)

// Serv implements the struct to inject the dependencies
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
	router.HandleFunc("/customer/all", s.Middleware(s.getAllCustomers)).Methods("GET")
	router.HandleFunc("/user/all", s.Middleware(s.getAllUsers)).Methods("GET")
	router.HandleFunc("/warning/all", s.Middleware(s.getAllWarnings)).Methods("GET")

	// Get ByID
	router.HandleFunc("/customer", s.Middleware(s.getCustomerByID)).Methods("GET").Queries("id", "{id}")
	router.HandleFunc("/user", s.Middleware(s.getUserByID)).Methods("GET").Queries("id", "{id}")
	router.HandleFunc("/warning", s.Middleware(s.getWarningByID)).Methods("GET").Queries("id", "{id}")
	router.HandleFunc("/public_func", s.Middleware(s.getPublicByID)).Methods("GET").Queries("id", "{id}")

	// Get ByName - by pattern
	router.HandleFunc("/customer", s.Middleware(s.getCustomerByName)).Methods("GET").Queries("name", "{pattern}")
	router.HandleFunc("/user", s.Middleware(s.getUserByEmail)).Methods("GET").Queries("email", "{pattern}")
	router.HandleFunc("/warning", s.Middleware(s.getWarningByCustomer)).Methods("GET").Queries("customer", "{pattern}")
	router.HandleFunc("/warning", s.Middleware(s.getWarningByUser)).Methods("GET").Queries("user", "{pattern}")
	router.HandleFunc("/public_func", s.Middleware(s.getPublicByWage)).Methods("GET").Queries("wage", "{pattern}")

	// Import
	router.HandleFunc("/customer/import", s.Middleware(s.ImportCustomerFile)).Methods("POST")      // TODO: posso fazer opção para escolher o arquivo
	router.HandleFunc("/public_func/import", s.Middleware(s.ImportPublicFuncFile)).Methods("POST") // TODO: posso fazer opção para escolher o mês

	// Post
	router.HandleFunc("/customer", s.Middleware(s.AddCustomer)).Methods("POST")
	router.HandleFunc("/user", s.addUser).Methods("POST") // without middleware
	router.HandleFunc("/warning", s.Middleware(s.AddWarning)).Methods("POST")
	router.HandleFunc("/public_func", s.Middleware(s.addPublicFunc)).Methods("POST")
	router.HandleFunc("/email", s.Middleware(s.sendEmail)).Methods("POST").Queries("user", "{userEmail}")

	// Delete
	router.HandleFunc("/customer", s.Middleware(s.deleteCustomerByID)).Methods("DELETE").Queries("id", "{id}")
	router.HandleFunc("/user", s.Middleware(s.deleteUserByID)).Methods("DELETE").Queries("id", "{id}")
	router.HandleFunc("/warning", s.Middleware(s.deleteWarningByID)).Methods("DELETE").Queries("id", "{id}")
	router.HandleFunc("/public_func", s.Middleware(s.deletePublicByID)).Methods("DELETE").Queries("id", "{id}")

	// Put
	router.HandleFunc("/customer", s.Middleware(s.updateCustomer)).Methods("PUT").Queries("id", "{id}")
	router.HandleFunc("/user", s.Middleware(s.updateUser)).Methods("PUT").Queries("id", "{id}")
	router.HandleFunc("/warning", s.Middleware(s.updateWarning)).Methods("PUT").Queries("id", "{id}")

	router.HandleFunc("/login", s.SignIn).Methods("POST")
	return router
}
