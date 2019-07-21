package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/backend/pkg/model"
	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("Don't_Panic")

// Claims makes the struct to deal with the JWT auth
type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

// SignIn handles the login control to the API
func (s *Serv) SignIn(w http.ResponseWriter, r *http.Request) {
	var user model.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		model.ErrorResponse(w, fmt.Errorf("dados inválidos: %v", err), http.StatusBadRequest)
		return
	}

	receivedPassword := user.Pass

	user, err = s.read.GetUserByEmail(user.Email)
	if err != nil {
		model.ErrorResponse(w, fmt.Errorf("dados inválidos: %v", err), http.StatusUnauthorized)
		return
	}

	err = model.IsPassword(user.Pass, receivedPassword)
	if err != nil {
		model.ErrorResponse(w, fmt.Errorf("dados inválidos: %v", err), http.StatusUnauthorized)
		return
	}
	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}

// Middleware handles the authorization to use the API
func (s *Serv) Middleware(handler func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				model.ErrorResponse(w, fmt.Errorf("acesso não autorizado"), http.StatusUnauthorized)
				return
			}
			model.ErrorResponse(w, fmt.Errorf("dados inválidos: %v", err), http.StatusBadRequest)
			return
		}

		// Get the JWT string from the cookie
		tknStr := c.Value

		claims := &Claims{}

		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if !tkn.Valid {
			model.ErrorResponse(w, fmt.Errorf("acesso não autorizado"), http.StatusUnauthorized)
			return
		}
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				model.ErrorResponse(w, fmt.Errorf("acesso não autorizado"), http.StatusUnauthorized)
				return
			}
			model.ErrorResponse(w, fmt.Errorf("dados inválidos: %v", err), http.StatusBadRequest)
			return
		}

		// if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		// 	w.WriteHeader(http.StatusBadRequest)
		// 	return
		// }

		// Now, create a new token for the current use, with a renewed expiration time
		expirationTime := time.Now().Add(5 * time.Minute)
		claims.ExpiresAt = expirationTime.Unix()
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			model.ErrorResponse(w, fmt.Errorf("erro interno: %v", err), http.StatusInternalServerError)
			return
		}

		// Set the new token as the users `token` cookie
		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
		})

		handler(w, r)
	})
}
