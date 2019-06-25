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

func makeFakeServices() http.Handler {

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

		{"TestGetCustomerByID", "/customer?id=1", "GET", http.StatusBadRequest},
		{"TestGetUserByID", "/user?id=1", "GET", http.StatusBadRequest},
		{"TestGetWarningByID", "/warning?id=1", "GET", http.StatusBadRequest},

		{"TestGetAllCustomers", "/customer/all", "GET", http.StatusBadRequest},
		{"TestGetAllUsers", "/user/all", "GET", http.StatusBadRequest},
		{"TestGetAllWarnings", "/warning/all", "GET", http.StatusBadRequest},

		{"TestGetCustomerByName", "/customer?name=Teste", "GET", http.StatusBadRequest},
		{"TestGetUserByEmail", "/user?email=teste@email", "GET", http.StatusBadRequest},

		{"TestUpdateCustomer", "/customer?id=1", "PUT", http.StatusBadRequest},
		{"TestUpdateUser", "/user?id=1", "PUT", http.StatusBadRequest},

		{"TestAddCustomer", "/customer", "POST", http.StatusBadRequest},
		{"TestAddUser", "/user", "POST", http.StatusBadRequest},
	}

	//TODO: Trocar estes StatusBadRequest por uma resposta mais informativa

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
