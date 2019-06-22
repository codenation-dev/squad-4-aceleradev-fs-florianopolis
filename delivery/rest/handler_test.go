package rest

import (
	"codenation/squad-4-aceleradev-fs-florianopolis/reading"
	"codenation/squad-4-aceleradev-fs-florianopolis/storage/postgres"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"

	"github.com/stretchr/testify/assert"
)

func TestGetAllCustomers(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "", nil)
	db, _, _ := sqlmock.New()
	read := reading.NewService(postgres.Storage{db})
	s := serv{}
	s.getAllCustomers(w, req)
	assert.Equal(t, 200, w.Code)

}
