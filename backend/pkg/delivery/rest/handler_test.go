package rest

import (
	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/backend/pkg/adding"
	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/backend/pkg/deleting"
	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/backend/pkg/reading"
	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/backend/pkg/storage/postgres"
	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/backend/pkg/updating"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/adams-sarah/test2doc/test"
	"github.com/adams-sarah/test2doc/vars"
	"github.com/stretchr/testify/assert"
)

func MakeFakeServices() *mux.Router {

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

	// set up HTTP server
	router := Handler(
		adder,
		reader,
		deleter,
		updater,
	)
	return router
}

// import (
// 	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/backend/pkg/adding"
// 	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/backend/pkg/deleting"
// 	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/backend/pkg/reading"
// 	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/backend/pkg/storage/postgres"
// 	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/backend/pkg/updating"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"os"
// 	"testing"

// 	"github.com/gorilla/mux"
// 	"github.com/stretchr/testify/assert"

// 	"github.com/DATA-DOG/go-sqlmock"
// 	"github.com/adams-sarah/test2doc/test"
// 	"github.com/adams-sarah/test2doc/vars"
// )

// var srv *test.Server

// func makeFakeServices() *mux.Router {

// 	// set services
// 	var adder adding.Service
// 	var reader reading.Service
// 	var deleter deleting.Service
// 	var updater updating.Service

// 	// If have more than one storage types, make the case/switch here
// 	db, _, err := sqlmock.New()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	s, err := postgres.NewStorage(db)
// 	if err != nil {
// 		log.Fatalf("could not set new storage: %v", err)
// 	}

// 	adder = adding.NewService(s)
// 	reader = reading.NewService(s)
// 	deleter = deleting.NewService(s)
// 	updater = updating.NewService(s)

// 	// set up HTTP server
// 	router := Handler(
// 		adder,
// 		reader,
// 		deleter,
// 		updater,
// 	)
// 	return router
// }

func TestMain(m *testing.M) {
	fakeRouter := MakeFakeServices()

	extractor := vars.MakeGorillaMuxExtractor(fakeRouter)
	test.RegisterURLVarExtractor(extractor)

	srv, err := test.NewServer(fakeRouter)
	if err != nil {
		panic(err)
	}

	code := m.Run()
	srv.Finish()
	os.Exit(code)

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
		{"TestDeleteWarningByID", "/warning?id=1", "DELETE", http.StatusBadRequest},
		{"TestDeletePublicByID", "/public_func?id=1", "DELETE", http.StatusBadRequest},

		{"TestGetCustomerByID", "/customer?id=1", "GET", http.StatusBadRequest},
		{"TestGetUserByID", "/user?id=1", "GET", http.StatusBadRequest},
		{"TestGetWarningByID", "/warning?id=1", "GET", http.StatusBadRequest},

		{"TestGetAllCustomers", "/customer/all", "GET", http.StatusBadRequest},
		{"TestGetAllUsers", "/user/all", "GET", http.StatusBadRequest},
		{"TestGetAllWarnings", "/warning/all", "GET", http.StatusBadRequest},

		{"TestGetCustomerByName", "/customer?name=Teste", "GET", http.StatusBadRequest},
		{"TestGetUserByEmail", "/user?email=teste@email", "GET", http.StatusBadRequest},
		// TODO: teste não funciona se uso espaço na query
		{"TestGetWarningByCustomer", "/warning?customer=teste_customer", "GET", http.StatusBadRequest},
		{"TestGetWarningByUser", "/warning?user=teste_user", "GET", http.StatusBadRequest},
		{"TestGetPublicByWage", "/public_func?wage=1234.56", "GET", http.StatusBadRequest},

		{"TestUpdateCustomer", "/customer?id=1", "PUT", http.StatusBadRequest},
		{"TestUpdateUser", "/user?id=1", "PUT", http.StatusBadRequest},
		{"TestUpdateWarning", "/warning?id=1", "PUT", http.StatusBadRequest},

		{"TestAddCustomer", "/customer", "POST", http.StatusBadRequest},
		{"TestAddUser", "/user", "POST", http.StatusBadRequest},
		{"TestAddWarning", "/warning", "POST", http.StatusBadRequest},
		{"TestAddPublicFunc", "/public_func", "POST", http.StatusBadRequest},
	}

	//TODO: Trocar estes StatusBadRequest por uma resposta mais informativa

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			router := MakeFakeServices()
			srv := httptest.NewServer(router)
			// srv, err := test.NewServer(router)
			// if err != nil {
			// 	panic(err)
			// }
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