package auth

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

var (
	// ErrInvalidCredentials is the standard error message to invalid authentication
	ErrInvalidCredentials = errors.New("Senha ou login inválidos")
	JwtKey                = []byte("secret_key")
)

type Credentials struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

// func SignIn(receivedPassword string, user model.User)(model.User error){
// 	user, err :=
// 	!= nil {
// 		return user, err
// 	}
// 	err =
// }

// // SignIn handles the route to authenticate the user
// func SignIn(w http.ResponseWriter, r *http.Request) {
// 	var creds Credentials

// 	err := json.NewDecoder(r.Body).Decode(&creds)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// }

// func SignIn(user model.User) (model.User, error) {
// 	password := user.Pass

// }
