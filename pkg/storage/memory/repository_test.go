package memory

import (
	"testing"

	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/pkg/entity"
)

var s = Storage{
	Users: []entity.User{{"pipa@email.com", "42"}},
}

var pame = entity.User{
	Email:    "pame@email.com",
	Password: "42",
}
var pipa = entity.User{
	Email:    "pipa@email.com",
	Password: "43",
}

var invalidUser = entity.User{
	Email:    "invalid",
	Password: "invalid",
}

func TestReadtUser(t *testing.T) {
	tt := []struct {
		name    string
		email   string
		want    string
		errWant error
	}{
		{"pipa", "pipa@email.com", "42", nil},
		{"not found", "not_found@email.com", "", entity.ErrUserNotFound},
	}

	for _, tc := range tt {

		t.Run(tc.name, func(t *testing.T) {
			u, err := s.ReadUser(tc.email)
			assertResponse(t, err, tc.errWant)
			assertResponse(t, u.Password, tc.want)

		})
	}
}

func TestCreateUser(t *testing.T) {
	tt := []struct {
		name        string
		user        entity.User
		errExpected error
	}{
		{"insert", pame, nil},
		{"duplicated", pipa, entity.ErrDuplicatedUser},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			err := s.CreateUser(tc.user)
			assertResponse(t, err, tc.errExpected)
		})
	}
}

func TestUpdateUser(t *testing.T) {
	tt := []struct {
		name          string
		payload       entity.User
		expectedError error
	}{
		{"success", pipa, nil},
		{"not found", invalidUser, entity.ErrUserNotFound},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			err := s.UpdateUser(tc.payload)
			assertResponse(t, err, tc.expectedError)
		})
	}
}

func TestDeleteUser(t *testing.T) {
	tt := []struct {
		name          string
		payload       string
		expectedError error
	}{
		{"success", pipa.Email, nil},
		{"not found", invalidUser.Email, entity.ErrUserNotFound},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			err := s.DeleteUser(tc.payload)
			assertResponse(t, err, tc.expectedError)
		})
	}
}

func assertResponse(t *testing.T, got, want interface{}) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
