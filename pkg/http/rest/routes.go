package rest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/pkg/emailserv"
	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/pkg/entity"

	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/pkg/service/adding"
	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/pkg/service/deleting"
	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/pkg/service/reading"
	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/pkg/service/updating"

	"github.com/gorilla/mux"
)

// NewRouter implements handlers to all routes
func NewRouter(adder adding.Service, reader reading.Service, updater updating.Service, deleter deleting.Service) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.Handle("/login", login(reader)).Methods(http.MethodPost)
	router.Handle("/", getIndex()).Methods(http.MethodGet)

	router.Handle("/user/{email}", getUser(reader)).Methods(http.MethodGet)
	router.Handle("/user/{email}", deleteUser(deleter)).Methods(http.MethodDelete)
	router.Handle("/user", addUser(adder)).Methods(http.MethodPost)
	router.Handle("/user", updateUser(updater)).Methods(http.MethodPut)

	router.Handle("/public_func", getPublicFunc(reader)).Methods(http.MethodGet)
	router.Handle("/public_func/import", importPublicFunc(adder)).Methods(http.MethodGet)

	router.Handle("/customer", getCustomer(reader)).Methods(http.MethodGet)
	router.Handle("/customer/import", importCustomer(adder)).Methods(http.MethodGet)

	router.Handle("/email_to", sendEmail(reader)).Methods(http.MethodPost)
	// router.Handle("/fetch/data/compare_customer_x_public_func/{company}/{uf}/{year}/{month}", compareCustomerPublicFunc(reader)).Methods(http.MethodGet)
	// router.Handle("/fetch/data/public_func_above_wage/{uf}/{year}/{month}/{wage}", getPublicFincByWage(reader)).Methods(http.MethodGet)

	router.Use(authorize)
	return router
}

func assertError(w http.ResponseWriter, status int, err error) {
	if err != nil {
		respondWithError(w, status, err)
		return
	}
}

func sendEmail(reader reading.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		b, err := ioutil.ReadAll(r.Body)
		assertError(w, http.StatusBadRequest, err)

		email := entity.Email{}
		err = json.Unmarshal(b, &email)
		assertError(w, http.StatusInternalServerError, err)

		err = emailserv.Send(email)
		assertError(w, http.StatusInternalServerError, err)

		respondWithJSON(w, http.StatusOK, "email enviado com sucesso")
	}
}

func getIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}

func respondWithError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	fmt.Fprint(w, err.Error())
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
