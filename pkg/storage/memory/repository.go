package memory

import (
	"github.com/gbletsch/squad-4-aceleradev-fs-florianopolis/pkg/entity"
)

type Storage struct {
	Users []entity.User
}

func NewStorage() *Storage {
	return new(Storage)
}

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
