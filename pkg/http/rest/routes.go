package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/pkg/entity"
	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/pkg/service/adding"
	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/pkg/service/deleting"
	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/pkg/service/reading"
	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/pkg/service/updating"

	"github.com/gorilla/mux"
)

// NewRouter implements handlers to all routes
func NewRouter(adder adding.Service, reader reading.Service, updater updating.Service, deleter deleting.Service, tpl *template.Template) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.Handle("/login", login(tpl)).Methods(http.MethodGet)
	router.Handle("/login/process", signin(reader, tpl)).Methods(http.MethodPost)
	router.Handle("/", getIndex()).Methods(http.MethodGet)

	router.Handle("/user/{email}", getUser(reader)).Methods(http.MethodGet)
	router.Handle("/user/{email}", deleteUser(deleter)).Methods(http.MethodDelete)
	router.Handle("/user", addUser(tpl)).Methods(http.MethodGet)
	router.Handle("/user/process", addUserProcess(adder, tpl)).Methods(http.MethodPost)
	router.Handle("/user", updateUser(updater)).Methods(http.MethodPut)

	router.Handle("/public_func/all", readAllPublicFunc(tpl)).Methods(http.MethodGet)
	router.Handle("/public_func/all/process", processReadAllPublicFunc(tpl, reader)).Methods(http.MethodPost)

	router.Handle("/customer/all", authorize(getAllCustomer(tpl))).Methods(http.MethodGet)
	router.Handle("/customer/all/process", authorize(processGetAllCustomer(tpl, reader))).Methods(http.MethodPost)

	router.Handle("/fetch/data/compare_customer_x_public_func", compareCustomerPublicFunc(tpl)).Methods(http.MethodGet)
	router.Handle("/fetch/data/compare_customer_x_public_func/process", processCompareCustomerPublicFunc(tpl, reader)).Methods(http.MethodPost)

	// router.Use(authorize)
	return router
}

func getIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}

func login(tpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.ExecuteTemplate(w, "login.html", nil)
	}
}

func updateUser(updater updating.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := entity.User{}
		b, err := ioutil.ReadAll(r.Body)

		if err != nil {
			respondWithError(w, http.StatusInternalServerError, errors.New("erro no ioutil"))
			return
		}
		err = json.Unmarshal(b, &user)
		if err != nil {
			respondWithError(w, http.StatusUnprocessableEntity, entity.ErrUnmarshal)
			return
		}
		err = updater.ChangePassword(user)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, err)
			return
		}
		fmt.Fprint(w, "usuário modificado com sucesso")
	}
}

func deleteUser(deleter deleting.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		err := deleter.DeleteUser(params["email"])
		if err != nil {
			respondWithError(w, http.StatusBadRequest, err)
			return
		}
		fmt.Fprint(w, "usuário deletado com sucesso")
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
