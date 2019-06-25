package rest

import (
	"codenation/squad-4-aceleradev-fs-florianopolis/adding"
	"codenation/squad-4-aceleradev-fs-florianopolis/deleting"
	"codenation/squad-4-aceleradev-fs-florianopolis/reading"
	"codenation/squad-4-aceleradev-fs-florianopolis/updating"
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
) http.Handler {
	s := serv{add: add, read: read, del: del, update: update}

	router := mux.NewRouter()
	router.HandleFunc("/", getHome).Methods("GET")

	// Get All
	router.HandleFunc("/customer/all", s.getAllCustomers).Methods("GET")
	router.HandleFunc("/user/all", s.getAllUsers).Methods("GET")
	router.HandleFunc("/warning/all", s.getAllWarnings).Methods("GET")

	// Get ByID
	router.HandleFunc("/customer", s.getCustomerByID).Methods("GET").Queries("id", "{id}")
	router.HandleFunc("/user", s.getUserByID).Methods("GET").Queries("id", "{id}")
	router.HandleFunc("/warning", s.getWarningByID).Methods("GET").Queries("id", "{id}")

	// Get ByName
	router.HandleFunc("/customer", s.getCustomerByName).Methods("GET").Queries("name", "{pattern}")
	router.HandleFunc("/user", s.getUserByEmail).Methods("GET").Queries("email", "{pattern}")
	router.HandleFunc("/warning", s.getWarningByCustomer).Methods("GET").Queries("customer", "{pattern}")
	router.HandleFunc("/warning", s.getWarningByUser).Methods("GET").Queries("user", "{pattern}")

	// Post
	router.HandleFunc("/customer", s.addCustomer).Methods("POST")
	router.HandleFunc("/user", s.addUser).Methods("POST")
	router.HandleFunc("/warning", s.addWarning).Methods("POST")

	// Delete
	router.HandleFunc("/customer", s.deleteCustomerByID).Methods("DELETE").Queries("id", "{id}")
	router.HandleFunc("/user", s.deleteUserByID).Methods("DELETE").Queries("id", "{id}")
	router.HandleFunc("/warning", s.deleteWarningByID).Methods("DELETE").Queries("id", "{id}")

	// Put
	router.HandleFunc("/customer", s.updateCustomer).Methods("PUT").Queries("id", "{id}")
	router.HandleFunc("/user", s.updateUser).Methods("PUT").Queries("id", "{id}")
	router.HandleFunc("/warning", s.updateWarning).Methods("PUT").Queries("id", "{id}")

	return router
}
