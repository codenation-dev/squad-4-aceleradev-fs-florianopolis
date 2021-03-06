package rest

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/pkg/service/reading"

	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/pkg/entity"
	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("Don't_Panic")

// Claims makes the struct to deal with the JWT auth
type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

type Token struct {
	Token string `json:"token"`
}

// SignIn handles the login control to the API
func login(reader reading.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		w.Header().Set("Content-Type", "application/json")

		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Headers", "Authorization") // You can add more headers here if needed
			return
		}
		// var user entity.User
		receivedUser := entity.User{}
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, entity.ErrUnauthorized.Error(), http.StatusUnauthorized)
			return
		}
		err = json.Unmarshal(b, &receivedUser)
		if err != nil {
			http.Error(w, entity.ErrUnauthorized.Error(), http.StatusUnauthorized)
			return
		}

		existingUser, err := reader.GetUser(receivedUser.Email)
		if err != nil {
			http.Error(w, entity.ErrUnauthorized.Error(), http.StatusUnauthorized)
			return
		}

		err = entity.IsPassword(existingUser.Password, receivedUser.Password)
		if err != nil {
			http.Error(w, entity.ErrUnauthorized.Error(), http.StatusUnauthorized)
			return
		}

		expirationTime := time.Now().Add(5 * time.Hour)

		claims := &Claims{
			Email: receivedUser.Email,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			http.Error(w, "error making tokenString", http.StatusInternalServerError)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
		})

		accessToken := &Token{
			Token: tokenString,
		}

		respondWithJSON(w, http.StatusOK, accessToken)
	}
}

func getClaims(c *http.Cookie) (*jwt.Token, *Claims, error) {
	// Get the JWT string from the cookie
	tknStr := c.Value

	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return tkn, claims, err
}

// Middleware handles the authorization to use the API
func authorize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		//w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Content-Type", "application/json")
		if (*r).Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Headers", "Authorization") // You can add more headers here if needed
			return
		}

		if r.RequestURI == "/" || r.RequestURI == "/login" {
			next.ServeHTTP(w, r)
			return
		} else if r.RequestURI == "/user" && r.Method == http.MethodPost {
			next.ServeHTTP(w, r)
			return
		}

		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				respondWithError(w, http.StatusUnauthorized, entity.ErrUnauthorized)
				return
			}
			http.Error(w, entity.ErrUnauthorized.Error(), http.StatusInternalServerError)
			return
		}

		tkn, claims, err := getClaims(c)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, err)
			return
		}

		if !tkn.Valid {
			respondWithError(w, http.StatusUnauthorized, entity.ErrUnauthorized)
			return
		}
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				respondWithError(w, http.StatusUnauthorized, entity.ErrUnauthorized)
				return
			}
			http.Error(w, entity.ErrUnauthorized.Error(), http.StatusInternalServerError)
			return
		}

		// // TODO: não consegui implementar essa contagem de 30 segundos
		// if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		// 	w.WriteHeader(http.StatusBadRequest)
		// 	return
		// }

		// create a new token for the current use, with a renewed expiration time
		expirationTime := time.Now().Add(5 * time.Hour)
		claims.ExpiresAt = expirationTime.Unix()
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			http.Error(w, "error creating new token", http.StatusInternalServerError)
			return
		}

		// Set the new token as the users `token` cookie
		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
		})

		next.ServeHTTP(w, r)
	})
}
