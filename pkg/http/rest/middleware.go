package rest

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
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

// func loginProcess(reader reading.Service, tpl *template.Template) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		receivedUser := entity.User{}
// 		receivedUser.Email = r.FormValue("email")
// 		receivedUser.Password = r.FormValue("pass")

// 		existingUser, err := reader.GetUser(receivedUser.Email)
// 		if err != nil {
// 			http.Error(w, entity.ErrUnauthorized.Error(), http.StatusUnauthorized)
// 			return
// 		}

// 		err = entity.IsPassword(existingUser.Password, receivedUser.Password)
// 		if err != nil {
// 			http.Error(w, entity.ErrUnauthorized.Error(), http.StatusUnauthorized)
// 			return
// 		}

// 		tpl.ExecuteTemplate(w, "options.html", nil)

// 	}
// }

// SignIn handles the login control to the API
func signin(reader reading.Service, tpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// var user entity.User

		receivedUser := entity.User{}
		receivedUser.Email = r.FormValue("email")
		receivedUser.Password = r.FormValue("pass")

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

		expirationTime := time.Now().Add(5 * time.Minute)

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

		// var resp = map[string]interface{}{"status": false, "message": "logged in"}
		// resp["token"] = tokenString //Store the token in the response
		// resp["user"] = receivedUser.Email

		// json.NewEncoder(w).Encode(resp)

		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
		})
		c, err := r.Cookie("token")
		fmt.Println(err, c)

		tpl.ExecuteTemplate(w, "options.html", nil)
	}
}

// func loggingMiddleware(next http.Handler) http.Handler {
//     return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//         // Do stuff here
//         log.Println(r.RequestURI)
//         // Call the next handler, which can be another middleware in the chain, or the final handler.
//         next.ServeHTTP(w, r)
//     })
// }

// Middleware handles the authorization to use the API
func authorize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var header = r.Header.Get("x-access-token") //Grab the token from the header

		header = strings.TrimSpace(header)
		fmt.Println("header:", header)

		listExcluded := []string{"/", "/login", "/login/process", "/user", "/user/process"}

		for _, path := range listExcluded {
			if r.RequestURI == path {
				next.ServeHTTP(w, r)
				return
			}
		}
		// for _, c := range r.Cookies() {
		// 	fmt.Println(c.Name)
		// }

		c, err := r.Cookie("token")
		// fmt.Println(err, c)
		if err != nil {
			if err == http.ErrNoCookie {
				http.Error(w, "merda de cookie", http.StatusUnauthorized)
				// http.Error(w, entity.ErrUnauthorized.Error(), http.StatusUnauthorized)
				return
			}
			http.Error(w, entity.ErrUnauthorized.Error(), http.StatusInternalServerError)
			return
		}

		// Get the JWT string from the cookie
		tknStr := c.Value

		claims := &Claims{}

		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if !tkn.Valid {
			http.Error(w, entity.ErrUnauthorized.Error(), http.StatusUnauthorized)
			return
		}
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				http.Error(w, entity.ErrUnauthorized.Error(), http.StatusUnauthorized)
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
		expirationTime := time.Now().Add(5 * time.Minute)
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
