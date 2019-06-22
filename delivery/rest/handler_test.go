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

func makeFakeServices() (
	serv, http.Handler,
	// adding.Service, reading.Service,
	// deleting.Service, updating.Service, http.Handler,
) {
	// set services
	var adder adding.Service
	var reader reading.Service
	var deleter deleting.Service
	var updater updating.Service

	// If have more than one storage types, make the case/switch here
	db, _, err := sqlmock.New()
	if err != nil {
		panic(err)
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

	serv := serv{add: adder, read: reader,
		del: deleter, update: updater}

	return serv, router
}

func TestGetAllCustomers(t *testing.T) {
	_, router := makeFakeServices()
	srv := httptest.NewServer(router)
	url := fmt.Sprintf("%s%s", srv.URL, "/customer/all")

	res, err := http.Get(url)
	// BadRequest is the right one because it reads the handler but
	// breaks in the call to the fake BD witch is the expected
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.NoError(t, err, "error in http.Get")

}
