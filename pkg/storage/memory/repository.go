package memory

import (
	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/pkg/entity"
)

// Storage mock db
type Storage struct {
	Users []entity.User
}

// NewStorage implements a new mock storage
func NewStorage() *Storage {
	return new(Storage)
}

// ReadUser reads user from mock db
func (m *Storage) ReadUser(email string) (entity.User, error) {
	for _, u := range m.Users {
		if u.Email == email {
			return u, nil
		}
	}
	return entity.User{}, entity.ErrUserNotFound
}

// CreateUser inserts a new user on the fake DB
func (m *Storage) CreateUser(u entity.User) error {
	_, err := m.ReadUser(u.Email)
	if err != nil {
		if err != entity.ErrUserNotFound {
			return entity.ErrDuplicatedUser
		}
		return err
	}
	m.Users = append(m.Users, u)
	return nil
}

func (m *Storage) UpdateUser(u entity.User) error {
	for i, existingUser := range m.Users {
		if existingUser.Email == u.Email {
			m.Users[i].Password = u.Password
			return nil
		}
	}
	return entity.ErrUserNotFound
}

func (m *Storage) DeleteUser(email string) error {
	for i, existingUser := range m.Users {
		if existingUser.Email == email {
			m.Users[i] = entity.User{} // do not delete it to mantain the index as an internal id
			return nil
		}
	}
	return entity.ErrUserNotFound
}
