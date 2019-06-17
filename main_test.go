package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var a = App{}

func TestMain(m *testing.M) {
	a.NewRouter()

	a.connectDB()

	code := m.Run()
	os.Exit(code)
}

func TestGet(t *testing.T) {
	tt := []struct {
		name           string
		path           string
		expectedStatus int
		expectedBody   interface{}
	}{
		{"get root", "/", http.StatusOK, "\"API Banco Uati\"\n"},
		{"get customers", "/customers", http.StatusOK, "[]"},
	}

	for _, tc := range tt {
		t.Run(tc.path, func(t *testing.T) {
			srv := httptest.NewServer(a.router) // mock the URL
			path := fmt.Sprintf("%s%s", srv.URL, tc.path)
			resp, err := http.Get(path)
			assert.Nil(t, err, "error in http.Get function")
			assert.Equal(t, tc.expectedStatus, resp.StatusCode)

			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}
			body := string(b)
			assert.Equal(t, tc.expectedBody, body)
		})

	}

}

func TestPost(t *testing.T) {
	t.Errorf("not implemented")

	// db, mock, err := sqlmock.New()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()
	// a.db = db

	// payload := Customer{Name: "Gui"}
	// b, err := json.Marshal(payload)
	// if err != nil {
	// 	t.Error(err)
	// }

	// c := Customer{}
	// err = json.Unmarshal(b, &c)
	// if err != nil {
	// 	t.Error(err)
	// }

	// mock.ExpectBegin()
	// result := sqlmock.NewResult(1, 1)
	// mock.ExpectExec("^INSERT INTO customers (.+)").WillReturnResult(nil)
	// mock.ExpectCommit()

	// srv := httptest.NewServer(a.router)
	// path := fmt.Sprintf("%s%s", srv.URL, "/customers")

	// buf := bytes.NewBuffer(b)
	// http.Post(path, "application/json", buf)

	// err = mock.ExpectationsWereMet()
	// if err != nil {
	// 	fmt.Printf("expectation failed: %v", err)
	// }

}

func TestPut(t *testing.T) {
	t.Errorf("not implemented")
}

func TestDelete(t *testing.T) {
	t.Errorf("not implemented")
}
