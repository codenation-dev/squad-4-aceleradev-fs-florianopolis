package validations

import (
	"codenation/squad-4-aceleradev-fs-florianopolis/models"
	"errors"
)

var (
	// ErrEmptyFields is a default error to empty data
	ErrEmptyFields = errors.New("Campos vazios")
	// ErrInvalidEmail is a default error to empty data
	ErrInvalidEmail = errors.New("Email inválido")
)

// ValidateNewUser validates user before saving on the db
func ValidateNewUser(u models.User) (models.User, error) {
	if IsEmpty(u.Email) || IsEmpty(u.Pass) || IsEmpty(u.Login) {
		return models.User{}, ErrEmptyFields
	}
	if !IsEmail(u.Email) {
		return models.User{}, ErrInvalidEmail
	}
	return u, nil
}
