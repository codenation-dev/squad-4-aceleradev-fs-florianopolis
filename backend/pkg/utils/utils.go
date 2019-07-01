package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// func SignIn(user entity.User)(entity.User error){
// 	password := user.Pass
// 	user, err :=
// 	!= nil {
// 		return user, err
// 	}
// 	err =
// }

func Bcrypt(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func IsPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
