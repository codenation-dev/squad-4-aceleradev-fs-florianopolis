package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// Bcrypt encripts the password
func Bcrypt(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword(
		[]byte(password), bcrypt.DefaultCost,
	)
}

// IsPassword compares the stored hash with the given password
func IsPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword),
		[]byte(password),
	)
}
