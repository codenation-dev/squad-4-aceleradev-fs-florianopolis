package validations

import (
	"github.com/badoux/checkmail"
)

// IsEmpty returns true if the string is empty
func IsEmpty(s string) bool {
	if s == "" {
		return true
	}
	return false
}

// IsEmail return true if its a valid mail address
func IsEmail(email string) bool {
	err := checkmail.ValidateFormat(email)
	if err != nil {
		return false
	}
	return true
}
