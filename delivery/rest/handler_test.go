package rest

import (
	"codenation/squad-4-aceleradev-fs-florianopolis/adding"
	"codenation/squad-4-aceleradev-fs-florianopolis/deleting"
	"codenation/squad-4-aceleradev-fs-florianopolis/reading"
	"codenation/squad-4-aceleradev-fs-florianopolis/storage/postgres"
	"codenation/squad-4-aceleradev-fs-florianopolis/updating"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func makeFakeServices() http.Handler { // adding.Service, reading.Service,
	// deleting.Service, updating.Service, http.Handler,

	// set services
	var adder adding.Service
	var reader reading.Service
	var deleter deleting.Service
	var updater updating.Service

	// If have more than one storage types, make the case/switch here
	db, _, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	s, err := postgres.NewStorage(db)
	if err != nil {
		log.Fatalf("could not set new storage: %v", err)
	}

	adder = adding.NewService(s)
	reader = reading.NewService(s)
	deleter = deleting.NewService(s)
	updater = updating.NewService(s)

	// set uo HTTP server
	router := Handler(
		adder,
		reader,
		deleter,
		updater,
	)
	return router
}

func TestHandler(t *testing.T) {
	var tt = []struct {
		name           string
		path           string
		method         string
		expectedStatus int
	}{
		{"TestGetHome", "/", "GET", http.StatusOK},
		{"TestDeleteUserbyID", "/user?id=1", "DELETE", http.StatusBadRequest},
		{"TestDeleteCustomerByID", "/customer?id=1", "DELETE", http.StatusBadRequest},
		{"TestGetAllCustomers", "/customer/all", "GET", http.StatusBadRequest},
		{"TestUpdateCustomer", "/customer?id=1", "PUT", http.StatusBadRequest},
		{"TestDeleteCustomerByID", "/customer?id=1", "DELETE", http.StatusBadRequest},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			router := makeFakeServices()
			srv := httptest.NewServer(router)
			defer srv.Close()

			url := fmt.Sprintf("%s%s", srv.URL, tc.path)
			req, err := http.NewRequest(tc.method, url, nil)
			assert.NoError(t, err, "error on http.NewRequest")
			res, err := http.DefaultClient.Do(req)
			assert.NoError(t, err, "error on http.DefaultClient")
			assert.Equal(t, tc.expectedStatus, res.StatusCode)

		})
	}
}

// func TestGetHome(t *testing.T) {
// 	router := makeFakeServices()
// 	srv := httptest.NewServer(router)
// 	defer srv.Close()
// 	url := fmt.Sprintf("%s%s", srv.URL, "/")

// 	res, err := http.Get(url)
// 	assert.Equal(t, http.StatusOK, res.StatusCode)
// 	assert.NoError(t, err, "error in http.Get")
// }

// func TestGetAllCustomers(t *testing.T) {
// 	router := makeFakeServices()
// 	srv := httptest.NewServer(router)
// 	defer srv.Close()
// 	url := fmt.Sprintf("%s%s", srv.URL, "/customer/all")

// 	res, err := http.Get(url)
// 	// BadRequest is the right one because it reads the handler but
// 	// breaks in the call to the fake BD witch is the expected
// 	//
// 	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
// 	assert.NoError(t, err, "error in http.Get")
// }

// func TestUpdateCustomer(t *testing.T) {
// 	router := makeFakeServices()
// 	srv := httptest.NewServer(router)
// 	defer srv.Close()
// 	url := fmt.Sprintf("%s%s", srv.URL, "/customer?id=1")

// 	req, err := http.NewRequest("PUT", url, nil)
// 	assert.NoError(t, err, "error on http.NewRequest")
// 	res, err := http.DefaultClient.Do(req)
// 	assert.NoError(t, err, "error on http.DefaultClient")
// 	// BadRequest is the right choice because it reads the handler but
// 	// breaks in the call to the fake BD witch is the expected
// 	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
// }

// func TestDeleteCustomerByID(t *testing.T) {
// 	t.Parallel()
// 	router := makeFakeServices()
// 	srv := httptest.NewServer(router)
// 	defer srv.Close()
// 	url := fmt.Sprintf("%s%s", srv.URL, "/customer?id=1")
// 	req, err := http.NewRequest("DELETE", url, nil)

// 	assert.NoError(t, err, "error on http.NewRequest")
// 	res, err := http.DefaultClient.Do(req)
// 	assert.NoError(t, err, "error on http.DefaultClient")
// 	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
// }

// func TestDeleteUserByID(t *testing.T) {
// 	t.Parallel()

// 	router := makeFakeServices()
// 	srv := httptest.NewServer(router)
// 	defer srv.Close()

// 	url := fmt.Sprintf("%s%s", srv.URL, "/user?id=1")
// 	req, err := http.NewRequest("DELETE", url, nil)
// 	assert.NoError(t, err, "error on http.NewRequest")
// 	res, err := http.DefaultClient.Do(req)
// 	assert.NoError(t, err, "error on http.DefaultClient")
// 	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
// }
