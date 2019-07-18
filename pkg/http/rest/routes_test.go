package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gbletsch/squad-4-aceleradev-fs-florianopolis/pkg/service/deleting"
	"github.com/gbletsch/squad-4-aceleradev-fs-florianopolis/pkg/service/updating"

	"github.com/gbletsch/squad-4-aceleradev-fs-florianopolis/pkg/entity"
	"github.com/gbletsch/squad-4-aceleradev-fs-florianopolis/pkg/service/adding"
	"github.com/gbletsch/squad-4-aceleradev-fs-florianopolis/pkg/service/reading"
	"github.com/gbletsch/squad-4-aceleradev-fs-florianopolis/pkg/storage/memory"
)

var db = memory.Storage{
	Users: []entity.User{
		{"pame@email.com", "123"},
	},
}

var pipa = entity.User{"pipa@email.com", "42"}
var notFoundUser = entity.User{"not_found@email.com", "not_found"}

var (
	adder   = adding.NewService(&db)
	reader  = reading.NewService(&db)
	updater = updating.NewService(&db)
	deleter = deleting.NewService(&db)

	router = NewRouter(adder, reader, updater, deleter)
)

func TestHandler(t *testing.T) {
	t.Run("test index", func(t *testing.T) {
		r, _ := http.NewRequest(http.MethodGet, "/index", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, r)
		assertResponse(t, w.Code, http.StatusOK)
		assertResponse(t, w.Body.String(), "TODO: inserir dos da API")
	})
}

func TestAddUser(t *testing.T) {
	tt := []struct {
		name           string
		payload        entity.User
		expectedStatus int
		expectedBody   string
	}{
		{"add success", pipa, http.StatusOK, "usuário adicionado com sucesso"},
		{"repeated user", db.Users[0], http.StatusBadRequest, entity.ErrDuplicatedUser.Error()},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			b, _ := json.Marshal(tc.payload)
			r, _ := http.NewRequest(http.MethodPost, "/user", bytes.NewBuffer(b))
			w := httptest.NewRecorder()
			router.ServeHTTP(w, r)
			assertResponse(t, w.Code, tc.expectedStatus)
			assertResponse(t, w.Body.String(), tc.expectedBody)
		})
	}
}

func TestGetUser(t *testing.T) {
	tt := []struct {
		name           string
		path           string
		expectedStatus int
		expectedBody   interface{}
	}{
		{"pipa", "/user/pipa@email.com", http.StatusOK, `{"email":"pipa@email.com","password":"42"}`},
		{"pame", "/user/pame@email.com", http.StatusOK, `{"email":"pame@email.com","password":"123"}`},
		{"not found", "/user/laika@email.com", http.StatusBadRequest, entity.ErrUserNotFound.Error()},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			r, _ := http.NewRequest(http.MethodGet, tc.path, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, r)
			assertResponse(t, w.Code, tc.expectedStatus)
			assertResponse(t, w.Body.String(), tc.expectedBody)
		})
	}
}

func TestUpdateUser(t *testing.T) {
	tt := []struct {
		name           string
		u              entity.User
		expectedStatus int
		expectedBody   string
	}{
		{"success", db.Users[0], http.StatusOK, "usuário modificado com sucesso"},
		{"fail", notFoundUser, http.StatusBadRequest, entity.ErrUserNotFound.Error()},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			b, err := json.Marshal(tc.u)
			assertResponse(t, err, nil)
			r, _ := http.NewRequest(http.MethodPut, "/user", bytes.NewBuffer(b))
			w := httptest.NewRecorder()
			router.ServeHTTP(w, r)
			assertResponse(t, w.Code, tc.expectedStatus)
			assertResponse(t, w.Body.String(), tc.expectedBody)
		})
	}
}

func TestDeleteUser(t *testing.T) {
	tt := []struct {
		name           string
		email          string
		expectedStatus int
		expectedBody   string
	}{
		{"success", "pame@email.com", http.StatusOK, "usuário deletado com sucesso"},
		{"fail", "not_found@email.com", http.StatusBadRequest, entity.ErrUserNotFound.Error()},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			r, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("/user/%s", tc.email), nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, r)
			assertResponse(t, w.Code, tc.expectedStatus)
			assertResponse(t, w.Body.String(), tc.expectedBody)
		})
	}
}

func assertResponse(t *testing.T, got, want interface{}) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
