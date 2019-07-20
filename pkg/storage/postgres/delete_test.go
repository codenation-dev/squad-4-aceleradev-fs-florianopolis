package postgres

import (
	"testing"

	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/pkg/entity"

	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/pkg/storage/memory"
)

var s = memory.NewStorage()

// var pipa = entity.User{"pipa@email.com", "42"}
var notFoundUser = entity.User{"not_found@email.com", "123"}

// TODO: isso testa a regra de negócio, não a funcão delete do BD
// Pra testar a função delete, acho que precisaria mockar o 'db.Exec'

func TestDeleteUser(t *testing.T) {
	s.Users = append(s.Users, pipa)
	tt := []struct {
		name     string
		email    string
		expected error
	}{
		{"success", pipa.Email, nil},
		{"not found", notFoundUser.Email, entity.ErrUserNotFound},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			err := s.DeleteUser(tc.email)
			if err != tc.expected {
				t.Errorf("got %v, want %v", err, tc.expected)
			}
		})
	}
}
