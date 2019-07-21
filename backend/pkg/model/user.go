// Package model gives the models to the data, taking care of them validation
package model

import (
	"errors"

	"github.com/badoux/checkmail"
	"golang.org/x/crypto/bcrypt"
)

// User models the user component
type User struct {
	ID int `json:"id"`
	// Login string `json:"login"` // TODO: retirar login, já implementei tudo com o email
	Email string `json:"email"`
	Pass  string `json:"pass"`
	// TODO: implementar uma opção de quais avisos receber (ex: salarios acima de 100k, estado de SP...)
}

var (
	// ErrEmptyFields is used when there is an empty field
	ErrEmptyFields = errors.New("Um ou mais campos vazios")
	// ErrInvalidEmail is the error used if the mail is invalid
	ErrInvalidEmail = errors.New("Email inválido")
)

// IsEmpty returns true if the field is empty
func IsEmpty(param string) bool {
	if param == "" {
		return true
	}
	return false
}

// IsEmail verify if its a validmail
func IsEmail(email string) bool {
	err := checkmail.ValidateFormat(email)
	if err != nil {
		return false
	}
	return true
}

// ValidateNewUser verify if its a valid user
func ValidateNewUser(user User) (User, error) {
	if IsEmpty(user.Email) || IsEmpty(user.Pass) {
		return User{}, ErrEmptyFields
	}
	if !IsEmail(user.Email) {
		return User{}, ErrInvalidEmail
	}
	return user, nil
}

// Bcrypt returns the hash generated from the given password
func Bcrypt(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// IsPassword verifies if the password is correct
func IsPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
