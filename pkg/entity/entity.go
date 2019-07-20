package entity

import (
	"github.com/badoux/checkmail"
	"golang.org/x/crypto/bcrypt"
)

// type App struct {
// 	db      Storage
// 	service Service
// 	router  *mux.Router
// }

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

const (
	CacheFolder = "../cmd/data/downloaded"
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
	if IsEmpty(user.Email) || IsEmpty(user.Password) {
		return User{}, ErrEmptyFields
	}
	if !IsEmail(user.Email) {
		return User{}, ErrInvalidEmail
	}
	bPassword, err := Bcrypt(user.Password)
	if err != nil {
		return user, err
	}
	user.Password = string(bPassword)

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
